package request

type CreateCustomerRequest struct {
	Name    string `json:"name" form:"name" validate:"required,max=255" example:"Fajar Agus"`
	Email   string `json:"email" form:"email" validate:"required,email" example:"fajar@gmail.com"`
	Address string `json:"address"  form:"address"  validate:"required,min=5" example:"Jl. Raya"`
}

type UpdateCustomerRequest struct {
	Name    string `json:"name" form:"name" validate:"required,min=3,max=255" example:"Fajar Agus"`
	Email   string `json:"email" form:"email" validate:"required,email" example:"fajar@gmail.com"`
	Address string `json:"address"  form:"address"  validate:"required,min=5" example:"Jl. Raya"`
}
