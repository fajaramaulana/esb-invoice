package filters

type ItemFilter struct {
	Name string `json:"name"`
}

func NewItemFilter() *ItemFilter {
	return &ItemFilter{}
}
