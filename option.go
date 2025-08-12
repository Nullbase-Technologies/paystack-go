package paystack

import "net/http"

type Option func(c *Client)

func WithSecretKey(key string) Option {
	return func(c *Client) {
		c.secretKey = key
	}
}

func WithHTTPClient(client *http.Client) Option {
	return func(c *Client) {
		c.c = client
	}
}
