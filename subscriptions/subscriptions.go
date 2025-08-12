package subscriptions

import (
	"context"

	"github.com/Nullbase-Technologies/paystack-go"
)

type Subscriptions struct {
	c *paystack.Client
}

func New(client *paystack.Client) *Subscriptions {
	return &Subscriptions{
		c: client,
	}
}

func (s *Subscriptions) Create(ctx context.Context) {

}
