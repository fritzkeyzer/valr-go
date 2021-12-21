package valr

import "github.com/shopspring/decimal"

type OrderBookEntry struct {
	Side       string          `json:"side"`
	Quantity   decimal.Decimal `json:"quantity"`
	Price      decimal.Decimal `json:"price"`
	Pair       string          `json:"currencyPair"`
	OrderCount int             `json:"orderCount"`
}
