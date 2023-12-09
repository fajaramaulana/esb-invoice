package filters

type CustomerFilter struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func NewCustomerFilter() *CustomerFilter {
	return &CustomerFilter{}
}
