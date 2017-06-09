package api

import (
	"encoding/json"
	"time"
)

type SearchResponse struct {
	TotalResults 	int 	`json:"total_results"`
	Results 	[]Movie
}

type Movie struct {
	Id 		int
	Title 		string
	VoteAverage 	float64 `json:"vote_average"`
	ReleaseDate 	string  `json:"release_date"`
}

func Search(term string) (*SearchResponse, error) {
	body, err := GetBody("/search/movie", term)
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

func (resp *SearchResponse) IterateMovies(iter func(Movie)) {
	for _, v := range resp.Results {
		iter(v)
	}
}

func (movie Movie) ToTime() time.Time {
	date, _ := time.Parse("2006-01-02", movie.ReleaseDate)
	return date
}