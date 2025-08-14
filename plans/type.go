package plans

import (
	"time"

	"github.com/Nullbase-Technologies/paystack-go"
)

type CreateOrUpdatePlanOptions struct {
	Name                        string `json:"name"`
	Amount                      uint64 `json:"amount"`
	Interval                    string `json:"interval"`
	Description                 string `json:"description,omitempty"`
	SendInvoices                bool   `json:"send_invoices,omitempty"`
	SendSMS                     bool   `json:"send_sms,omitempty"`
	Currency                    string `json:"currency,omitempty"`
	InvoiceLimit                int    `json:"invoice_limit,omitempty"`
	UpdateExistingSubscriptions bool   `json:"update_existing_subscriptions,omitempty"`
}

type CreatePlanResponse struct {
	Name        string `json:"name"`
	Amount      uint64 `json:"amount"`
	Interval    string `json:"interval"`
	Integration uint64 `json:"integration"`
	Domain      string `json:"domain"`
	PlanCode    string `json:"plan_code"`
	SendInvoice bool   `json:"send_invoices"`
	SendSMS     bool   `json:"send_sms"`
	HostedPage  bool   `json:"hosted_page"`
	Currency    string `json:"currency"`
	ID          uint64 `json:"id"`
}

type Plan struct {
	Subscriptions []struct {
		Customer         uint64                 `json:"customer"`
		Plan             uint64                 `json:"plan"`
		Integration      uint64                 `json:"integration"`
		Domain           string                 `json:"domain"`
		Start            uint                   `json:"start"`
		Status           string                 `json:"status"`
		Quantity         uint                   `json:"quantity"`
		Amount           uint64                 `json:"amount"`
		SubscriptionCode string                 `json:"subscription_code"`
		EmailToken       string                 `json:"email_token"`
		Authorization    paystack.Authorization `json:"authorization"`
		EasyCronID       *uint64                `json:"easy_cron_id"`
		CronExpression   string                 `json:"cron_expression"`
		NextPaymentDate  time.Time              `json:"next_payment_date"`
		ID               uint64                 `json:"id"`
	} `json:"subscriptions"`
	Name              string  `json:"name"`
	Description       string `json:"description"`
	Amount            uint64  `json:"amount"`
	Interval          string  `json:"interval"`
	Integration       uint64  `json:"integration"`
	Domain            string  `json:"domain"`
	PlanCode          string  `json:"plan_code"`
	SendInvoice       bool    `json:"send_invoices"`
	SendSMS           bool    `json:"send_sms"`
	HostedPage        bool    `json:"hosted_page"`
	HostedPageURL     *string `json:"hosted_page_url"`
	HostedPageSummary *string `json:"hosted_page_summary"`
	Currency          string  `json:"currency"`
	ID                uint64  `json:"id"`
}

type ListPlanParams struct {
	paystack.Paginator
	Status   string `url:"status,omitempty"`
	Interval string `url:"interval,omitempty"`
	Amount   uint64 `url:"amount,omitempty"`
}
