package api

import (
	"net/http"
	"net/url"
	"io/ioutil"
)

const apiKey = ""
const apiUrl = "https://api.themoviedb.org/3"

func DoGet(path string, query string) ([]byte, error) {
	client := &http.Client{}

	parameters := url.Values{}
	parameters.Add("api_key", apiKey)
	if query != "" {
		parameters.Add("query", query)
	}

	req, err := http.NewRequest("GET", apiUrl + path + "?" + parameters.Encode(), nil)
	req.Close = true
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