package response

type MetaPagination struct {
	Code          int    `json:"code"`
	Status        string `json:"status"`
	Message       string `json:"message"`
	TotalFiltered int    `json:"totalFiltered"`
	TotalRecords  int    `json:"totalRecords"`
	Page          int    `json:"page"`
	PerPage       int    `json:"perPage"`
}

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponsePagination struct {
	Meta MetaPagination `json:"meta"`
	Data interface{}    `json:"data"`
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ReturnResponseError struct {
	Meta  Meta        `json:"meta"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}
