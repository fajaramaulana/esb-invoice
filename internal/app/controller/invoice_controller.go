package controller

import (
	"esb-invoice/internal/app/handler/helper"
	"esb-invoice/internal/app/handler/request"
	"esb-invoice/internal/domain/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type InvoiceController struct {
	invoiceService *service.InvoiceService
}

func NewInvoiceController(invoiceService *service.InvoiceService) *InvoiceController {
	return &InvoiceController{
		invoiceService: invoiceService,
	}
}

// CreateInvoice godoc
// @Summary Create invoice
// @Description Create invoice
// @Tags invoice
// @Accept  json
// @Produce  json
// @Param invoice body request.CreateInvoiceRequest
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ReturnResponseError
// @Failure 500 {object} response.ReturnResponseError
// @Router /api/v1/invoice [post]
func (c *InvoiceController) CreateInvoice(ctx *gin.Context) {
	var req request.CreateInvoiceRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		message, data := helper.GlobalCheckingErrorBindJson(err.Error())
		log.Println(fmt.Sprintf("Error: %s", message))
		helper.ReturnJSONError(ctx, http.StatusBadRequest, message, nil, data)
		return
	}

	res := helper.DoValidation(req)

	if len(res) > 0 {
		log.Println(fmt.Sprintf("Error: %s", "Validation error"))
		helper.ReturnJSONError(ctx, http.StatusBadRequest, "Validation error", nil, res)
		return
	}

	idInvoice, err := c.invoiceService.Create(&req)
	if err != nil {
		fmt.Printf("%# v\n", "masuk error insert")
		log.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	helper.ReturnJSON(ctx, http.StatusCreated, "Invoice created", idInvoice)
}
