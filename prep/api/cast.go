package api

import (
	"encoding/json"
	"strconv"
)

type CastResponse struct {
	Id 		int
	Cast	 	[]Cast
}

type Cast struct {
	Id 		int
	Name 		string
	Character 	string
}
func Credits(movieId int) (*CastResponse, error) {
	body, err := GetBody("/movie/" + strconv.Itoa(movieId) + "/credits", "")
	if err != nil {
		return nil, err
	}

	res := CastResponse{}
	err = json.Unmarshal(*body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (resp *CastResponse) Iterate(iterator func (person Cast)) {
	for _, person := range resp.Cast {
		iterator(person)
	}
}