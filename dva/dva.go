package dva

import (
	"context"

	"github.com/Nullbase-Technologies/paystack-go"
)

type VirtualAccount struct {
	c *paystack.Client
}

func New(client *paystack.Client) *VirtualAccount {
	return &VirtualAccount{
		c: client,
	}
}

func (d *VirtualAccount) Create(ctx context.Context, opts *CreateVirtualAccountOptions) {

}
