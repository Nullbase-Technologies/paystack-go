package paystack

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// defaultBaseUrl is the base endpoint for all requests
	defaultBaseUrl = "https://api.paystack.co"
	// defaultUserAgent is the default user agent, can be changed with options
	defaultUserAgent = "Nullbase-Technologies/paystack-go-sdk"
	// defaultTimeout is the default timeout on the http client
	defaultTimeout = 15 * time.Second
)

type AuthTransport struct {
	originalTransport http.RoundTripper
	secret            string
}

func (a *AuthTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", a.secret))
	return a.originalTransport.RoundTrip(r)
}

type Client struct {
	c         *http.Client
	secretKey string
	userAgent string

	// exporting this for use in other sub-packages
	BaseUrl *url.URL
}

func New(opts ...Option) (*Client, error) {
	c := &Client{}

	for _, opt := range opts {
		opt(c)
	}

	if IsStringEmpty(c.secretKey) {
		return nil, errors.New("please provide a valid secret key")
	}

	// ignoring err because... what error can come out from a static value :)
	base, _ := url.Parse(defaultBaseUrl)
	c.BaseUrl = base

	if IsStringEmpty(c.userAgent) {
		c.userAgent = defaultUserAgent
	}

	if c.c == nil {
		c.c = &http.Client{
			Transport: &AuthTransport{
				originalTransport: http.DefaultTransport,
				secret:            c.secretKey,
			},
			Timeout: defaultTimeout,
		}
	}

	c.secretKey = ""

	return c, nil
}

type Response struct {
	Status  bool            `json:"status"`
	Message string          `json:"message"`
	Data    json.RawMessage `json:"data,omitempty"`
}

func (c *Client) DoRequest(ctx context.Context, method, uri string, body, result any) error {
	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return err
		}
	}

	u, err := c.BaseUrl.Parse(uri)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", c.userAgent)
	req = req.WithContext(ctx)

	res, err := c.c.Do(req)
	if err != nil {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		return err
	}
	defer res.Body.Close()

	var resp Response

	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return err
	}

	if res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("paystack: %s", resp.Message)
	}

	if resp.Data != nil && result != nil {
		if err := json.Unmarshal(resp.Data, &result); err != nil {
			return err
		}
	}

	return nil
}

func IsStringEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
