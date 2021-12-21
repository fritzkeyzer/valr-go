package valr

import (
	"encoding/json"
	"fmt"
	"time"
)

// Public
func (cl *Client) GetPublicOrderBook(pair string) (GetOrderBookResponse, error) {
	var res GetOrderBookResponse
	resJsonBytes, err := cl.SendRequest("GET", "/v1/public/"+pair+"/orderbook")
	if err != nil {
		return GetOrderBookResponse{}, err
	}
	json.Unmarshal(resJsonBytes, &res)

	return res, nil
}

// Public
func (cl *Client) GetPublicOrderBookFull(pair string) (GetOrderBookFullResponse, error) {
	var res GetOrderBookFullResponse
	resJsonBytes, err := cl.SendRequest("GET", "/v1/public/"+pair+"/orderbook/full")
	if err != nil {
		return GetOrderBookFullResponse{}, err
	}
	json.Unmarshal(resJsonBytes, &res)

	return res, nil
}

// Public
func (cl *Client) GetPublicCurrencies() ([]Currency, error) {
	var res []Currency
	resJsonBytes, err := cl.SendRequest("GET", "/v1/public/currencies")
	if err != nil {
		return []Currency{}, err
	}
	json.Unmarshal(resJsonBytes, &res)

	return res, nil
}

// Public
func (cl *Client) GetPublicPairs() ([]CurrencyPair, error) {
	var res []CurrencyPair
	resJsonBytes, err := cl.SendRequest("GET", "/v1/public/pairs")
	if err != nil {
		return []CurrencyPair{}, err
	}
	json.Unmarshal(resJsonBytes, &res)

	return res, nil
}

// Public
func (cl *Client) GetPublicOrderTypes() ([]OrderTypes, error) {
	var res []OrderTypes
	resJsonBytes, err := cl.SendRequest("GET", "/v1/public/ordertypes")
	if err != nil {
		return []OrderTypes{}, err
	}
	json.Unmarshal(resJsonBytes, &res)

	return res, nil
}

func (cl *Client) GetOrderBook(pair string) (GetOrderBookResponse, error) {
	var res GetOrderBookResponse
	resJsonBytes, err := cl.SendRequest("GET", "/v1/marketdata/"+pair+"/orderbook")
	if err != nil {
		return GetOrderBookResponse{}, err
	}
	json.Unmarshal(resJsonBytes, &res)

	return res, nil
}

func (cl *Client) GetBalances() ([]AccountBalance, error) {
	var res []AccountBalance
	resJsonBytes, err := cl.SendRequest("GET", "/v1/account/balances")
	if err != nil {
		return []AccountBalance{}, err
	}
	json.Unmarshal(resJsonBytes, &res)

	return res, nil
}

func (cl *Client) GetTradeHistory(pair string, skip int, limit int, startTime time.Time, stopTime time.Time) ([]TradeHistoryEntry, error) {
	var res []TradeHistoryEntry

	var startTimeString = startTime.UTC().Format("2006-01-02T15:04:05.999999Z")
	var stopTimeString = stopTime.UTC().Format("2006-01-02T15:04:05.999999Z")

	var path = fmt.Sprintf("/v1/marketdata/%s/tradehistory?skip=%v&limit=%v&startTime=%s&endTime=%s", pair, skip, limit, startTimeString, stopTimeString)
	resJsonBytes, err := cl.SendRequest("GET", path)
	if err != nil {
		return []TradeHistoryEntry{}, err
	}
	json.Unmarshal(resJsonBytes, &res)

	return res, nil
}
