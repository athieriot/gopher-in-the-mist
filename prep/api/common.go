package api

import (
	"net/http"
	"net/url"
	"io/ioutil"
)

const API_KEY = ""
const API_URL = "https://api.themoviedb.org/3"

func GetBody(path string, query string) (*[]byte, error) {
	client := &http.Client{}

	parameters := url.Values{}
	parameters.Add("api_key", API_KEY)
	if query != "" {
		parameters.Add("query", query)
	}

	req, err := http.NewRequest("GET", API_URL + path + "?" + parameters.Encode(), nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return &body, nil
}
