package customers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Nullbase-Technologies/paystack-go"
	"github.com/google/go-querystring/query"
)

type Customers struct {
	c *paystack.Client
}

func New(client *paystack.Client) *Customers {
	return &Customers{
		c: client,
	}
}

// Create a customer on your integration
func (c *Customers) Create(ctx context.Context, opts *CreateCustomerOptions) (*CreateCustomerResponse, error) {
	resp := new(CreateCustomerResponse)

	err := c.c.DoRequest(ctx, http.MethodPost, "customer", opts, resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (c *Customers) List(ctx context.Context, params *ListCustomersParams) ([]Customer, error) {
	url := "customer"

	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	if !paystack.IsStringEmpty(queryParams.Encode()) {
		url = fmt.Sprintf("%s?%s", url, queryParams.Encode())
	}

	var resp []Customer

	if err := c.c.DoRequest(ctx, http.MethodGet, url, nil, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// Fetch gets details of a customer on your integration.
func (c *Customers) Fetch(ctx context.Context, emailOrCode string) (*Customer, error) {
	resp := new(Customer)

	if err := c.c.DoRequest(ctx, http.MethodGet,
		fmt.Sprintf("customers/%s", emailOrCode), nil, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// Update a customer's details on your integration
func (c *Customers) Update(ctx context.Context, code string, opts *UpdateCustomerOptions) (*Customer, error) {
	resp := new(Customer)

	if err := c.c.DoRequest(ctx, http.MethodPut,
		fmt.Sprintf("customer/%s", code), opts, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Customers) Validate(ctx context.Context, code string, opts *ValidateCustomerOptions) (bool, error) {
	if err := c.c.DoRequest(ctx,
		http.MethodPost, fmt.Sprintf("customer/%s/identification", code), opts, nil); err != nil {
		return false, err
	}

	return true, nil
}

// Whitelist or blacklist a customer on your integration
func (c *Customers) WhitelistOrBlacklist(ctx context.Context, opts *WhitelistOrBlacklistCustomerOptions) (*Customer, error) {
	resp := new(Customer)

	if err := c.c.DoRequest(ctx,
		http.MethodPost, "customer/set_risk_action", opts, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// InitializeAuthorization Initiate a request to create a reusable
// authorization code for recurring transactions.
func (c *Customers) InitializeAuthorization(ctx context.Context, opts *InitializeAuthorizationOptions) (*InitializeAuthorizationResponse, error) {
	resp := new(InitializeAuthorizationResponse)

	if err := c.c.DoRequest(ctx,
		http.MethodPost, "customer/authorization/initialize", opts, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// VerifyAuthorization checks the status of an authorization request.
func (c *Customers) VerifyAuthorization(ctx context.Context, reference string) (*VerifyAuthorizationResponse, error) {
	resp := new(VerifyAuthorizationResponse)

	if err := c.c.DoRequest(ctx,
		http.MethodGet, fmt.Sprintf("customer/authorization/verify/%s", reference), nil, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Customers) DeactivateAuthorization(ctx context.Context, code string) error {
	type opts struct {
		Code string `json:"authorization_code"`
	}

	return c.c.DoRequest(ctx, http.MethodPost, "customer/authorization/deactivate", opts{Code: code}, nil)
}
