package paystack

type Metadata map[string]interface{}

type Paginator struct {
	PerPage int64 `url:"perPage,omitempty"`
	Page    int64 `url:"page,omitempty"`
}
