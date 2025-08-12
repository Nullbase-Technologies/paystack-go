package transactions

import (
	"context"
	"net/http"

	"github.com/Nullbase-Technologies/paystack-go"
)

type Transactions struct {
	c *paystack.Client
}

func New(client *paystack.Client) *Transactions {
	return &Transactions{
		c: client,
	}
}

func (t *Transactions) Initialize(ctx context.Context, opts *InitializeTransactionOptions) (*InitializeTransactionResponse, error) {
	resp := new(InitializeTransactionResponse)

	err := t.c.DoRequest(ctx, http.MethodPost, "transaction/initialize", opts, resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (t *Transactions) Verify(ctx context.Context, reference string) {

}
