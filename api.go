package valr

import (
	"encoding/json"
	"time"
)

type GetOrderBookRequest struct {
	// Currency pair of the Orders to retrieve
	//
	// required: true
	Pair string `json:"pair" url:"pair"`
}

type GetOrderBookResponse struct {
	Asks      []OrderBookEntry `json:"Asks"`
	Bids      []OrderBookEntry `json:"Bids"`
	Timestamp time.Time        `json:"LastChange"`
}

func (cl *Client) GetOrderBook(pair string) (GetOrderBookResponse, error) {
	res := GetOrderBookResponse{}
	resJsonBytes, err := cl.SendRequest("GET", "/v1/marketdata/"+pair+"/orderbook")
	if err != nil {
		return GetOrderBookResponse{}, err
	}
	json.Unmarshal(resJsonBytes, &res)

	return res, nil
}
