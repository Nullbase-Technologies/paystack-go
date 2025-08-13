package split

import "github.com/Nullbase-Technologies/paystack-go"

type CreateSplitOptions struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Currency    string `json:"currency"`
	Subaccounts []struct {
		Subaccount string `json:"subaccount"`
		Share      int    `json:"share"`
	} `json:"subaccounts"`
	BearerType    string `json:"bearer_type"`
	BearerAccount string `json:"bearer_account"`
}

type Split struct {
	ID          uint64 `json:"id"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Currency    string `json:"currency"`
	Integration uint64 `json:"integration"`
	Domain      string `json:"domain"`
	SplitCode   string `json:"split_code"`
	Active      bool   `json:"active"`
	BearerType  string `json:"bearer_type"`
	IsDynamic   bool   `json:"is_dynamic"`
	Subaccount  []struct {
		Subaccount []struct {
			ID                  uint64  `json:"id"`
			SubaccountCode      string  `json:"subaccount_code"`
			BusinessName        string  `json:"business_name"`
			Description         string  `json:"description"`
			PrimaryContactName  *string `json:"primary_contact_name"`
			PrimaryContactEmail *string `json:"primary_contact_email"`
			PrimaryContactPhone *string `json:"primary_contact_phone"`
			Metadata            any     `json:"metadata"`
			SettlementBank      string  `json:"settlement_bank"`
			Currency            string  `json:"currency"`
			AccountNumber       string  `json:"account_number"`
		} `json:"subaccount"`
		Share int `json:"share"`
	} `json:"subaccounts"`
	TotalSubaccounts int `json:"total_subaccounts"`
}

type ListSplitParams struct {
	paystack.Paginator
	paystack.Period
	Name   string `url:"name"`
	Active bool   `url:"active"`
	SortBy string `url:"sort_by"`
}

type UpdateSplitOptions struct {
	Name             string `json:"name"`
	Active           bool   `json:"active"`
	BearerType       string `json:"bearer_type,omitempty"`
	BearerSubaccount string `json:"bearer_subaccount,omitempty"`
}

type AddOrUpdateSubaccountOptions struct {
	Subaccount string `json:"subaccount"`
	Share      int    `json:"share"`
}

type RemoveSubaccountOptions struct {
	Subaccount string `json:"subaccount"`
}
