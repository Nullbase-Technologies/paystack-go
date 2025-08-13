package miscellaneous

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Nullbase-Technologies/paystack-go"
	"github.com/google/go-querystring/query"
)

type Miscellaneous struct {
	c *paystack.Client
}

func New(client *paystack.Client) *Miscellaneous {
	return &Miscellaneous{
		c: client,
	}
}

func (m *Miscellaneous) ListBanks(ctx context.Context,
	params *ListBankParams) ([]Bank, error) {
	url := "bank"

	queryParams, err := query.Values(params)
	if err != nil {
		return nil, err
	}

	if !paystack.IsStringEmpty(queryParams.Encode()) {
		url = fmt.Sprintf("%s?%s", url, queryParams.Encode())
	}

	var banks []Bank

	if err := m.c.DoRequest(ctx, http.MethodGet, url,
		nil, &banks); err != nil {
		return nil, err
	}

	return banks, nil
}
