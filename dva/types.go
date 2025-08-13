package dva

type CreateVirtualAccountOptions struct {
	// Customer ID or code
	Customer      string `json:"customer"`
	PreferredBank string `json:"preferred_bank,omitempty"`
	Subaccount    string `json:"subaccount,omitempty"`
	SplitCode     string `json:"split_code,omitempty"`
	FirstName     string `json:"first_name,omitempty"`
	LastName      string `json:"last_name,omitempty"`
	Phone         string `json:"phone,omitempty"`
}
