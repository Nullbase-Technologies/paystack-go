package transactions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Nullbase-Technologies/paystack-go"
	"github.com/google/go-querystring/query"
)

type Transactions struct {
	c *paystack.Client
}

func New(client *paystack.Client) *Transactions {
	return &Transactions{
		c: client,
	}
}

// Initialize starts a new transaction from your backend
func (t *Transactions) Initialize(ctx context.Context, opts *InitializeTransactionOptions) (*InitializeTransactionResponse, error) {
	resp := new(InitializeTransactionResponse)

	err := t.c.DoRequest(ctx, http.MethodPost, "transaction/initialize", opts, resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// Verify confirms the status of a transaction
func (t *Transactions) Verify(ctx context.Context, reference string) (*Transaction, error) {
	resp := new(Transaction)

	if err := t.c.DoRequest(ctx,
		http.MethodGet, fmt.Sprintf("transaction/verify/%s", reference), nil, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// List transactions carried out on your integration
func (t *Transactions) List(ctx context.Context, params *ListTransactionsParams) ([]Transaction, error) {
	url := "transaction"

	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	if !paystack.IsStringEmpty(queryParams.Encode()) {
		url = fmt.Sprintf("%s?%s", url, queryParams.Encode())
	}

	var resp []Transaction

	if err := t.c.DoRequest(ctx, http.MethodGet, url, nil, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

// Fetch gets details of a transaction carried out on your integration
func (t *Transactions) Fetch(ctx context.Context, id uint64) (*Transaction, error) {
	resp := new(Transaction)

	if err := t.c.DoRequest(ctx,
		http.MethodGet, fmt.Sprintf("transaction/%d", id), nil, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *Transactions) ChargeAuthorization(ctx context.Context, opts *ChargeAuthorizationOptions) (*Transaction, error) {
	resp := new(Transaction)

	err := t.c.DoRequest(ctx, http.MethodPost,
		"transaction/charge_authorization", opts, resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (t *Transactions) GetTransactionTimeline(ctx context.Context, ref string) (*TransactionTimeline, error) {
	resp := new(TransactionTimeline)

	err := t.c.DoRequest(ctx, http.MethodGet,
		fmt.Sprintf("transaction/timeline/%s", ref), nil, resp)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
