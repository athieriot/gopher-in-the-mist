package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
	"encoding/json"
	"os"
)

const apiKey = ""
const apiUrl = "https://api.themoviedb.org/3/search/movie"

type SearchResponse struct {
	Page 		int
	TotalResults 	int `json:"total_results"`
	Results 	[]Movie
}

type Movie struct {
	Id 		int
	Title 		string
	Overview 	string
	VoteAverage	float64 `json:"vote_average"`
}

func Search(movie string) (*SearchResponse, error) {
	body, err := DoSearch(movie)
	if err != nil {
		return nil, err
	}

	res := SearchResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func DoSearch(movie string) ([]byte, error) {
	client := &http.Client{}

	parameters := url.Values{}
	parameters.Add("api_key", apiKey)
	parameters.Add("query", movie)

	req, err := http.NewRequest("GET", apiUrl + "?" + parameters.Encode(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func main() {
	body, err := Search("Wonder woman")
	if err != nil {
		fmt.Println("An error occured", err.Error())
		os.Exit(1)
	}

	fmt.Println("The result is:", body)
}