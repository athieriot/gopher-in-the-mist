package api

import (
	"encoding/json"
	"strconv"
)

type CastResponse struct {
	Id 		int
	Cast	 	[]Person
}

type Person struct {
	Id 		int
	Name 		string
	Character 	string
}

func Cast(movieId int) (*CastResponse, error) {
	body, err := DoGet("/movie/" + strconv.Itoa(movieId) + "/credits", "")
	if err != nil {
		return nil, err
	}

	res := CastResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (resp *CastResponse) Iterate(iterator func(value Person)) {
	for _, v := range resp.Cast {
		iterator(v)
	}
}