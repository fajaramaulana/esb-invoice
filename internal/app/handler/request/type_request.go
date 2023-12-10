package request

type CreateTypeRequest struct {
	Name string `form:"name" validate:"required,max=255" binding:"required" example:"Fajar Agus"`
}

type UpdateTypeRequest struct {
	Name string `form:"name" validate:"required,min=3,max=255" binding:"required" example:"Fajar Agus"`
}
