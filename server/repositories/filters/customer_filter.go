package filters

type CustomerFilter struct {
	Name string `json:"name"`
}

func NewCustomerFilter() *CustomerFilter {
	return &CustomerFilter{}
}
