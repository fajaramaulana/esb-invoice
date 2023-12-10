package request

type CreateCustomerRequest struct {
	Name    string `json:"name" form:"name" validate:"required,min=3,max=255" example:"Fajar Agus" binding:"required"`
	Email   string `json:"email" form:"email" validate:"required,email" example:"fajar@gmail.com" binding:"required"`
	Address string `json:"address" form:"address"  validate:"required,min=5" example:"Jl. Raya" binding:"required"`
}

type UpdateCustomerRequest struct {
	Name    string `json:"name" form:"name" validate:"required,min=3,max=255" example:"Fajar Agus" binding:"required"`
	Email   string `json:"email" form:"email" validate:"required,email" example:"fajar@gmail.com" binding:"required"`
	Address string `json:"address" form:"address"  validate:"required,min=5" example:"Jl. Raya" binding:"required"`
}
