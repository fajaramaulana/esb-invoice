package request

type CreateItemRequest struct {
	NameItem string  `json:"name_item" name:"Name_Item" form:"name_item" validate:"required,min=3,max=255" example:"VGA" binding:"required"`
	TypeId   int     `json:"type_id" form:"type_id" validate:"required" example:"1" binding:"required"`
	Price    float64 `json:"price" form:"price" validate:"required" example:"1000000" binding:"required"`
}

type UpdateItemRequest struct {
	NameItem string  `json:"name_item" form:"name_item" validate:"required,min=3,max=255" example:"VGA" binding:"required"`
	TypeId   int     `json:"type_id" form:"type_id" validate:"required" example:"1" binding:"required"`
	Price    float64 `json:"price" form:"price" validate:"required" example:"1000000" binding:"required"`
}
