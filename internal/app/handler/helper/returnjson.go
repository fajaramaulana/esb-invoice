package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Meta  Meta        `json:"meta"`
	Data  interface{} `json:"data"`
	Error interface{} `json:"error"`
}

type ResponsePagination struct {
	Meta MetaPagination `json:"meta"`
	Data interface{}    `json:"data"`
}

type Meta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type MetaPagination struct {
	Code          int    `json:"code"`
	Status        string `json:"status"`
	Message       string `json:"message"`
	TotalFiltered int    `json:"totalFiltered"`
	TotalRecords  int    `json:"totalRecords"`
	Page          int    `json:"page"`
	PerPage       int    `json:"perPage"`
}

func ReturnJSON(ctx *gin.Context, code int, message string, data interface{}) {
	meta := Meta{
		Code:    code,
		Status:  http.StatusText(code),
		Message: message,
	}

	response := Response{
		Meta: meta,
		Data: data,
	}

	ctx.JSON(code, response)
}

func ReturnJSONError(ctx *gin.Context, code int, message string, data interface{}, err interface{}) {
	meta := Meta{
		Code:    code,
		Status:  http.StatusText(code),
		Message: message,
	}

	response := ResponseError{
		Meta:  meta,
		Data:  data,
		Error: err,
	}

	ctx.JSON(code, response)
}

func ReturnJSONWithMeta(ctx *gin.Context, code int, message string, data interface{}, totalRecords int, totalFiltered int, page int, perPage int) {
	meta := MetaPagination{
		Code:          code,
		Status:        http.StatusText(code),
		Message:       message,
		TotalRecords:  totalRecords,
		TotalFiltered: totalFiltered,
		Page:          page,
		PerPage:       perPage,
	}

	response := ResponsePagination{
		Meta: meta,
		Data: data,
	}

	ctx.JSON(code, response)
}
