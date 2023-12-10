package controller

import (
	"esb-invoice/internal/app/handler/filters"
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

// GetInvoice by id
// @Summary Get invoice by id
// @Description Get invoice by id
// @Tags invoice
// @Produce  json
// @Param id path int true "Invoice ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.ResponseError
// @Router /api/v1/invoice/{id} [get]
func (c *InvoiceController) FindInvoiceById(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := helper.ConvertStringToInt(id)
	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSONError(ctx, http.StatusBadRequest, err.Error(), nil, map[string]interface{}{"error": err.Error()})
		return
	}

	invoice, err := c.invoiceService.GetInvoice(intId)
	if err != nil {
		log.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	helper.ReturnJSON(ctx, http.StatusOK, "Invoice found", invoice)
}

// FindAllInvoice godoc
// @Summary Find all invoice With Filter
// @Description Find all invoice
// @Tags invoice
// @Produce  json
// @Param page query int false "Page"
// @Param page_size query int false "Page Size"
// @Param invoice_id query int false "Invoice ID"
// @Param issue_date query string false "Issue Date"
// @Param due_date query string false "Due Date"
// @Param subject query string false "Subject"
// @Param total_items query int false "Total Items"
// @Param customer_name query string false "Customer Name"
// @Param payment_status query int false "Payment Status"
// @Success 200 {object} response.ResponsePagination
// @Failure 400 {object} response.ResponseError
// @Failure 500 {object} response.Response
// @Router /api/v1/invoice [get]
func (c *InvoiceController) FindAllInvoice(ctx *gin.Context) {
	var req filters.InvoiceFilter
	var returnMessage string
	if err := ctx.ShouldBind(&req); err != nil {
		message, data := helper.GlobalCheckingErrorBindJson(err.Error())
		log.Println(fmt.Sprintf("Error: %s", message))
		helper.ReturnJSONError(ctx, http.StatusBadRequest, message, nil, data)
		return
	}

	// check query param
	if err := ctx.ShouldBind(&req); err != nil {
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

	page, pageSize, err := helper.PaginationHelper(ctx)
	if err != nil {
		log.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusBadRequest, err.Error(), nil, nil)
		return
	}

	invoices, totalRecords, err := c.invoiceService.FindAll(&req, page, pageSize)
	if err != nil {
		log.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	// get total data
	totalData, err := c.invoiceService.CountAll()

	if err != nil {
		log.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	if len(invoices) == 0 {
		returnMessage = fmt.Sprintf("Data not found")
	} else {

		returnMessage = fmt.Sprintf("Data found")
	}

	helper.ReturnJSONWithMeta(ctx, http.StatusOK, returnMessage, invoices, int(totalData), int(totalRecords), page, pageSize)
}
