package response

type ItemResponse struct {
	ID     uint               `json:"id"`
	Name   string             `json:"name"`
	Price  float64            `json:"price"`
	TypeID uint               `json:"typeId"`
	Type   TypeOnItemResponse `json:"type"`
}
