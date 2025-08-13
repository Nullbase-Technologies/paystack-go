package split

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Nullbase-Technologies/paystack-go"
	"github.com/google/go-querystring/query"
)

type TransactionSplit struct {
	c *paystack.Client
}

func New(client *paystack.Client) *TransactionSplit {
	return &TransactionSplit{
		c: client,
	}
}

func (t *TransactionSplit) Create(ctx context.Context, opts *CreateSplitOptions) (*Split, error) {
	resp := new(Split)

	if err := t.c.DoRequest(ctx, http.MethodPost, "split", opts, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TransactionSplit) List(ctx context.Context, params *ListSplitParams) ([]Split, error) {
	url := "split"

	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	if !paystack.IsStringEmpty(queryParams.Encode()) {
		url = fmt.Sprintf("%s?%s", url, queryParams.Encode())
	}

	var resp []Split

	if err := t.c.DoRequest(ctx, http.MethodGet, url, nil, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TransactionSplit) Fetch(ctx context.Context, id uint64) (*Split, error) {
	resp := new(Split)

	if err := t.c.DoRequest(ctx,
		http.MethodGet, fmt.Sprintf("split/%d", id), nil, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TransactionSplit) Update(ctx context.Context,
	id uint64, opts *UpdateSplitOptions) (*Split, error) {
	resp := new(Split)

	if err := t.c.DoRequest(ctx,
		http.MethodPut, fmt.Sprintf("split/%d", id), opts, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TransactionSplit) AddOrUpdateSubaccount(ctx context.Context,
	id uint64, opts *AddOrUpdateSubaccountOptions) (*Split, error) {

	resp := new(Split)

	if err := t.c.DoRequest(ctx,
		http.MethodPost, fmt.Sprintf("split/%d/subaccount/add", id), opts, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (t *TransactionSplit) RemoveSubaccount(ctx context.Context,
	id uint64, opts *RemoveSubaccountOptions) (*Split, error) {

	resp := new(Split)

	if err := t.c.DoRequest(ctx,
		http.MethodPost, fmt.Sprintf("split/%d/subaccount/remove", id), opts, resp); err != nil {
		return nil, err
	}

	return resp, nil
}
