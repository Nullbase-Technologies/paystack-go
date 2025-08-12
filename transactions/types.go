package transactions

import (
	"time"

	"github.com/Nullbase-Technologies/paystack-go"
)

type InitializeTransactionOptions struct {
	Amount            uint64                 `json:"amount"`
	Email             string                 `json:"email"`
	Currency          string                 `json:"currency,omitempty"`
	Reference         string                 `json:"reference,omitempty"`
	CallbackURL       string                 `json:"callback_url,omitempty"`
	Plan              string                 `json:"plan,omitempty"`
	InvoiceLimit      int64                  `json:"invoice_limit,omitempty"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	Channels          []string               `json:"channels,omitempty"`
	SplitCode         string                 `json:"split_code,omitempty"`
	Subaccount        string                 `json:"subaccount,omitempty"`
	TransactionCharge int64                  `json:"transaction_charge,omitempty"`
	Bearer            string                 `json:"bearer,omitempty"`
}

type InitializeTransactionResponse struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type Transaction struct {
	ID              uint64     `json:"id"`
	Domain          string     `json:"domain"`
	Status          string     `json:"status"`
	Reference       string     `json:"reference"`
	ReceiptNumber   *string    `json:"receipt_number"`
	Amount          uint64     `json:"amount"`
	Message         string     `json:"message"`
	GatewayResponse string     `json:"gateway_response"`
	PaidAt          *time.Time `json:"paid_at"`
	Created         time.Time  `json:"created_at"`
	Channel         string     `json:"channel"`
	Currency        string     `json:"currency"`
	IPAddress       string     `json:"ip_address"`
	Log             struct {
		TransactionTimeline
	} `json:"log"`
	Fees uint64 `json:"fees"`
}

type TransactionTimeline struct {
	StartTime int64     `json:"start_time"`
	TimeSpent int64     `json:"time_spent"`
	Attempts  int16     `json:"attempts"`
	Errors    int16     `json:"errors"`
	Success   bool      `json:"success"`
	Mobile    bool      `json:"mobile"`
	Input     []string  `json:"input"`
	History   []History `json:"history"`
}

type History struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Time    int16  `json:"time"`
}

type ListTransactionsParams struct {
	paystack.Paginator
	CustomerID uint64    `url:"customer,omitempty"`
	TerminalID string    `url:"terminalid,omitempty"`
	Status     string    `url:"status,omitempty"`
	From       time.Time `url:"from,omitempty"`
	To         time.Time `url:"to,omitempty"`
	Amount     uint64    `url:"amount,omitempty"`
}

type ChargeAuthorizationOptions struct {
	Amount            uint64                 `json:"amount"`
	Email             string                 `json:"email"`
	AuthorizationCode string                 `json:"authorization_code"`
	Reference         string                 `json:"reference,omitempty"`
	Currency          string                 `json:"currency,omitempty"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
	Channels          []string               `json:"channels,omitempty"`
	Subaccount        string                 `json:"subaccount,omitempty"`
	TransactionCharge uint64                 `json:"transaction_charge,omitempty"`
	Bearer            string                 `json:"bearer,omitempty"`
	Queue             bool                   `json:"queue,omitempty"`
}
