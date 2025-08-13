package miscellaneous

import "github.com/Nullbase-Technologies/paystack-go"

type Bank struct {
	Name        string `json:"name"`
	Slug        string `json:"slug"`
	Code        string `json:"code"`
	LongCode    string `json:"longcode"`
	Gateway     string `json:"gateway"`
	PayWithBank bool   `json:"pay_with_bank"`
	Active      bool   `json:"active"`
	IsDeleted   bool   `json:"is_deleted"`
	Country     string `json:"country"`
	Currency    string `json:"currency"`
	Type        string `json:"type"`
	ID          uint64 `json:"id"`
}

type ListBankParams struct {
	paystack.Paginator
	Country                string `url:"country"`
	PayWithBankTransfer    bool   `url:"pay_with_bank_transfer,omitempty"`
	PayWithBank            bool   `url:"pay_with_bank,omitempty"`
	EnabledForVerification bool   `url:"enabled_for_verification,omitempty"`
	Gateway                string `url:"gateway,omitempty"`
	Type                   string `url:"type,omitempty"`
	Currency               string `url:"currency,omitempty"`
	IncludeNipSortCode     bool   `url:"include_nip_sort_code,omitempty"`
}
