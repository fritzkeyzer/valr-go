package valr

import (
	"github.com/shopspring/decimal"
	"time"
)

type OrderBookEntry struct {
	Side       string          `json:"side"`
	Quantity   decimal.Decimal `json:"quantity"`
	Price      decimal.Decimal `json:"price"`
	Pair       string          `json:"currencyPair"`
	OrderCount int             `json:"orderCount"`
}

type OrderBookFullEntry struct {
	Side            string          `json:"side"`
	Quantity        decimal.Decimal `json:"quantity"`
	Price           decimal.Decimal `json:"price"`
	Pair            string          `json:"currencyPair"`
	PositionAtPrice int             `json:"positionAtPrice"`
}

type Currency struct {
	Symbol                  string    `json:"symbol"`
	IsActive                bool      `json:"isActive"`
	ShortName               string    `json:"shortName"`
	LongName                string    `json:"longName"`
	DecimalPlaces           quotedInt `json:"decimalPlaces" type:"integer"`
	WithdrawalDecimalPlaces quotedInt `json:"withdrawalDecimalPlaces"`
}

type CurrencyPair struct {
	Symbol            string          `json:"symbol"`
	BaseCurrency      string          `json:"baseCurrency"`
	QuoteCurrency     string          `json:"quoteCurrency"`
	ShortName         string          `json:"shortName"`
	Active            bool            `json:"active"`
	MinBaseAmount     decimal.Decimal `json:"minBaseAmount"`
	MaxBaseAmount     decimal.Decimal `json:"maxBaseAmount"`
	MinQuoteAmount    decimal.Decimal `json:"minQuoteAmount"`
	MaxQuoteAmount    decimal.Decimal `json:"maxQuoteAmount"`
	TickSize          decimal.Decimal `json:"tickSize"`
	BaseDecimalPlaces quotedInt       `json:"baseDecimalPlaces"`
}

type OrderTypes struct {
	CurrencyPair string   `json:"currencyPair"`
	OrderTypes   []string `json:"orderTypes"`
}

type AccountBalance struct {
	Currency  string          `json:"currency"`
	Available decimal.Decimal `json:"available"`
	Reserved  decimal.Decimal `json:"reserved"`
	Total     decimal.Decimal `json:"total"`
	UpdatedAt time.Time       `json:"updatedAt"`
}

type TradeHistoryEntry struct {
	Price        decimal.Decimal `json:"price"`
	Quantity     decimal.Decimal `json:"quantity"`
	CurrencyPair string          `json:"currencyPair"`
	TradedAt     time.Time       `json:"tradedAt"`
	TakerSide    string          `json:"takerSide"`
	SequenceId   int             `json:"sequenceId"`
	Id           string          `json:"id"`
	QuoteVolume  decimal.Decimal `json:"quoteVolume"`
}
