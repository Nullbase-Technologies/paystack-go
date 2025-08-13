package paystack

import "time"

type Metadata map[string]interface{}

type Paginator struct {
	PerPage   int64  `url:"perPage,omitempty"`
	Page      int64  `url:"page,omitempty"`
	Previous  string `url:"previous,omitempty"`
	UseCursor bool   `url:"use_cursor,omitempty"`
}

type Period struct {
	From time.Time `url:"from,omitempty"`
	To   time.Time `url:"from,omitempty"`
}

type Authorization struct {
	AuthorizationCode string  `json:"authorization_code"`
	Bin               string  `json:"bin"`
	LastFour          string  `json:"last4"`
	ExpMonth          string  `json:"exp_month"`
	ExpYear           string  `json:"exp_year"`
	Channel           string  `json:"channel"`
	Bank              string  `json:"bank"`
	CardType          string  `json:"card_type"`
	CountryCode       string  `json:"country_code"`
	Brand             string  `json:"brand"`
	Reusable          bool    `json:"reusable"`
	Signature         string  `json:"signature"`
	Account           *string `json:"account"`
}
