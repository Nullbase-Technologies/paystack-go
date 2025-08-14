package plans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Nullbase-Technologies/paystack-go"
	"github.com/google/go-querystring/query"
)

type Plans struct {
	c *paystack.Client
}

func New(client *paystack.Client) *Plans {
	return &Plans{
		c: client,
	}
}

func (p *Plans) Create(ctx context.Context,
	opts *CreateOrUpdatePlanOptions) (*CreatePlanResponse, error) {
	resp := new(CreatePlanResponse)

	if err := p.c.DoRequest(ctx, http.MethodPost,
		"plan", opts, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *Plans) List(ctx context.Context, params *ListPlanParams) ([]Plan, error) {
	url := "plan"

	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	if !paystack.IsStringEmpty(queryParams.Encode()) {
		url = fmt.Sprintf("%s?%s", url, queryParams.Encode())
	}

	var resp []Plan

	if err := p.c.DoRequest(ctx, http.MethodGet, url, nil, &resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *Plans) Fetch(ctx context.Context, code string) (*Plan, error) {
	resp := new(Plan)

	if err := p.c.DoRequest(ctx,
		http.MethodGet, fmt.Sprintf("plan/%s", code), nil, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *Plans) Update(ctx context.Context,
	code string, opts *CreateOrUpdatePlanOptions) error {
	return p.c.DoRequest(ctx, http.MethodPut, fmt.Sprintf("plan/%s", code), opts, nil)
}
