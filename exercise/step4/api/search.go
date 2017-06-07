package api

import (
	"encoding/json"
)

type SearchResponse struct {
	Page 		int
	TotalResults 	int 	`json:"total_results"`
	Results 	[]Movie
}

type Movie struct {
	Id 		int
	Title 		string
	Overview 	string
	VoteAverage	float64 `json:"vote_average"`
	ReleaseDate	string 	`json:"release_date"`
}

func Search(movie string) (*SearchResponse, error) {
	body, err := DoGet("/search/movie", movie)
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

func (resp *SearchResponse) Iterate(iterator func(value Movie)) {
	for _, v := range resp.Results {
		iterator(v)
	}
}