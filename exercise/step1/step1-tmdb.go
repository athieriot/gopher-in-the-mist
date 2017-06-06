package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
)

const apiKey = "043941d9826350a407cd88a648f2d62c"
const apiUrl = "https://api.themoviedb.org/3/search/movie"

func Search(movie string) (string, error) {
	client := &http.Client{}

	parameters := url.Values{}
	parameters.Add("api_key", apiKey)
	parameters.Add("query", movie)

	req, err := http.NewRequest("GET", apiUrl + "?" + parameters.Encode(), nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func main() {
	body, err := Search("Wonder woman")
	if err != nil {
		fmt.Println("An error occured", err.Error())
	} else {
		fmt.Println("The result is:", body)
	}
}