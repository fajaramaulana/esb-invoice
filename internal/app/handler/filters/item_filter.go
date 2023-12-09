package filters

type ItemFilter struct {
	ItemName string `json:"item_name"`
	TypeName string `json:"type_name"`
}

func NewItemFilter() *ItemFilter {
	return &ItemFilter{}
}
