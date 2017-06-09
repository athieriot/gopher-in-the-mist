package api

import (
	"encoding/json"
)

type SearchResponse struct {
	TotalResults 	int `json:"total_results"`
	Results 	[]Movie
}

type Movie struct {
	Id 		int
	Title 		string
	VoteAverage 	float64 `json:"vote_average"`
	ReleaseDate 	string 	`json:"release_date"`
}
func Search(term string) (*SearchResponse, error) {
	body, err := GetBody("/search/movie", term)
	if err != nil {
		return nil, err
	}

	res := SearchResponse{}
	err = json.Unmarshal(*body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (resp *SearchResponse) Iterate(iterator func (movie Movie)) {
	for _, movie := range resp.Results {
		iterator(movie)
	}
}