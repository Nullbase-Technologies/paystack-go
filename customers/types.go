package customers

import "time"

type CreateCustomerOptions struct {
	Email     string                 `json:"email"`
	FirstName string                 `json:"first_name"`
	LastName  string                 `json:"last_name"`
	Phone     string                 `json:"phone,omitempty"`
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
}

type CreateCustomerResponse struct {
	Email        string    `json:"email"`
	Integration  uint64    `json:"integration"`
	Domain       string    `json:"domain"`
	CustomerCode string    `json:"customer_code"`
	ID           uint64    `json:"id"`
	Identified   bool      `json:"identified"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}
