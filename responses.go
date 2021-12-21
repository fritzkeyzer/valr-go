package valr

import "time"

type GetOrderBookResponse struct {
	Asks      []OrderBookEntry `json:"Asks"`
	Bids      []OrderBookEntry `json:"Bids"`
	Timestamp time.Time        `json:"LastChange"`
}

type GetOrderBookFullResponse struct {
	Asks      []OrderBookFullEntry `json:"Asks"`
	Bids      []OrderBookFullEntry `json:"Bids"`
	Timestamp time.Time            `json:"LastChange"`
}
