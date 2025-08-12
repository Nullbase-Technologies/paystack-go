package transactions

type InitializeTransactionOptions struct {
	Amount            int64                  `json:"amount"`
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
