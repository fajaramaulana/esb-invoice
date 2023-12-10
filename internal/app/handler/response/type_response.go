package response

type TypeResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetByIdType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateTypeResponse struct {
	ID int `json:"id"`
}

type UpdateTypeResponse struct {
	ID int `json:"id"`
}

type DeleteTypeResponse struct {
	ID int `json:"id"`
}

type TypeOnItemResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
