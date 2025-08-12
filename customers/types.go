package customers

import (
	"time"

	"github.com/Nullbase-Technologies/paystack-go"
)

type CreateCustomerOptions struct {
	Email     string                 `json:"email"`
	FirstName string                 `json:"first_name"`
	LastName  string                 `json:"last_name"`
	Phone     string                 `json:"phone,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type CreateCustomerResponse struct {
	Email        string    `json:"email"`
	Integration  uint64    `json:"integration"`
	Domain       string    `json:"domain"`
	CustomerCode string    `json:"customer_code"`
	ID           uint64    `json:"id"`
	Identified   bool      `json:"identified"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type Customer struct {
	Email             string                   `json:"email"`
	FirstName         string                   `json:"first_name"`
	LastName          string                   `json:"last_name"`
	Phone             string                   `json:"phone,omitempty"`
	Metadata          map[string]interface{}   `json:"metadata,omitempty"`
	Integration       uint64                   `json:"integration"`
	Domain            string                   `json:"domain"`
	CustomerCode      string                   `json:"customer_code"`
	ID                uint64                   `json:"id"`
	Identified        bool                     `json:"identified"`
	CreatedAt         time.Time                `json:"created_at"`
	UpdatedAt         time.Time                `json:"updated_at"`
	Authorization     []paystack.Authorization `json:"authorizations"`
	RiskAction        string                   `json:"risk_action"`
	TotalTransactions uint16                   `json:"total_transactions"`
}

type ListCustomersParams struct {
	paystack.Paginator
	From time.Time `json:"from,omitempty"`
	To   time.Time `json:"to,omitempty"`
}

type UpdateCustomerOptions struct {
	FirstName string                 `json:"first_name"`
	LastName  string                 `json:"last_name"`
	Phone     string                 `json:"phone,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type ValidateCustomerOptions struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	// Only bank_account is supported at the moment
	// As at 12th August, 2025
	Type  string `json:"type"`
	Value string `json:"value"`
	// 2 letter country code of identification issuer
	Country       string `json:"country"`
	BVN           string `json:"bvn"`
	BankCode      string `json:"bank_code"`
	AccountNumber string `json:"account_number,omitempty"`
	MiddleName    string `json:"middle_name,omitempty"`
}

type WhitelistOrBlacklistCustomerOptions struct {
	// Customer's code, or email address
	Customer string `json:"customer"`
	// One of the possible risk actions [ default, allow, deny ].
	// allow to whitelist. deny to blacklist.
	// Customers start with a default risk action.
	RiskAction string `json:"risk_action,omitempty"`
}

type InitializeAuthorizationOptions struct {
	Email string `json:"email"`
	// direct-debit is the only supported option for now
	Channel     string `json:"channel"`
	CallbackURL string `json:"callback_url,omitempty"`
}

type InitializeAuthorizationResponse struct {
	AuthorizationURL string `json:"authorization_url"`
	AccessCode       string `json:"access_code"`
	Reference        string `json:"reference"`
}

type VerifyAuthorizationResponse struct {
	AuthorizationCode string `json:"authorization_code"`
	Channel           string `json:"channel"`
	Bank              string `json:"bank"`
	Active            bool   `json:"active"`
	Customer          struct {
		Code  string `json:"code"`
		Email string `json:"email"`
	} `json:"customer"`
}
