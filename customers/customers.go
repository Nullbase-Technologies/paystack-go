package customers

import (
	"context"
	"net/http"

	"github.com/Nullbase-Technologies/paystack-go"
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
