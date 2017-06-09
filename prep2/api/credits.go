package api

import (
	"encoding/json"
	"strconv"
)

type CreditResponse struct {
	Id 	int
	Cast 	[]Cast
}

type Cast struct {
	Id 		int
	Name 		string
	Character	string
}

func Credits(movieId int) (*CreditResponse, error) {
	body, err := GetBody("/movie/" + strconv.Itoa(movieId) + "/credits", "")
	if err != nil {
		return nil, err
	}

	res := CreditResponse{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (resp *CreditResponse) IterateCast(iter func(Cast)) {
	for _, v := range resp.Cast {
		iter(v)
	}
}