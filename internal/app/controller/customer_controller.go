package controller

import (
	"esb-invoice/internal/app/handler/filters"
	"esb-invoice/internal/app/handler/helper"
	"esb-invoice/internal/app/handler/request"
	"esb-invoice/internal/app/handler/response"
	"esb-invoice/internal/domain/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService *service.CustomerService
}

func NewCustomerController(customerService *service.CustomerService) *CustomerController {
	return &CustomerController{
		customerService: customerService,
	}
}

// CreateUser godoc
// @Summary Create User
// @Description Create User
// @Tags Customer
// @Accept  form-data
// @Produce  json
// @Param user body request.CreateCustomerRequest true "Customer"
// @Success 201 {object} response.ReturnResponseCreate
// @Failure 400 {object} response.ReturnResponseError
// @Failure 500 {object} response.ReturnResponseError
// @Router api/v1/customer [post]
func (c *CustomerController) CreateCustomer(ctx *gin.Context) {
	var req request.CreateCustomerRequest

	// Bind request body to struct
	// If the structure of the body is wrong, return an HTTP error with status code 400
	// request should json or form-data
	if err := ctx.ShouldBind(&req); err != nil {
		returnDataErrorCheck := helper.ExtractFieldNameFromError(err.Error())
		log.Println("Error: Validation error")
		helper.ReturnJSON(ctx, http.StatusBadRequest, "Validation error", returnDataErrorCheck)
		return
	}

	res := helper.DoValidation(req)

	if len(res) > 0 {
		log.Println("Error: Validation error")
		helper.ReturnJSON(ctx, http.StatusBadRequest, "Validation error", res)
		return
	}

	// check if email already exist
	checkByEmail, err := c.customerService.FindByEmail(req.Email)
	if err != nil {
		if err.Error() != "record not found" {
			log.Println("Error:", err)
			helper.ReturnJSON(ctx, http.StatusInternalServerError, err.Error(), nil)
			return
		}
	}

	if checkByEmail != nil {
		log.Println("Error: Email already exist")
		helper.ReturnJSON(ctx, http.StatusBadRequest, "Email already exist", nil)
		return
	}

	customerId, err := c.customerService.Create(&req)

	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSON(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ReturnJSON(ctx, http.StatusCreated, "Customer created", response.CreateCustomerResponse{Id: customerId})
}

// UpadateCustomer godoc
// @Summary Update Customer
// @Description Update Customer
// @Tags Customer
// @Accept  form-data
// @Produce  json
// @Param id path int true "Customer ID"
// @Param user body request.UpdateCustomerRequest true "Customer"
// @Success 200 {object} response.ReturnResponseUpdate
// @Failure 400 {object} response.ReturnResponseError
// @Failure 500 {object} response.ReturnResponseError.
// @Router api/v1/customer/{id} [put]

func (c *CustomerController) UpdateCustomer(ctx *gin.Context) {
	var req request.UpdateCustomerRequest
	id := ctx.Param("id")
	intId, err := helper.ConvertStringToInt(id)
	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSON(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	// Bind request body to struct
	// If the structure of the body is wrong, return an HTTP error with status code 400
	// request should json or form-data
	if err := ctx.ShouldBind(&req); err != nil {
		returnDataErrorCheck := helper.ExtractFieldNameFromError(err.Error())
		log.Println("Error: Validation error")
		helper.ReturnJSON(ctx, http.StatusBadRequest, "Validation error", returnDataErrorCheck)
		return
	}

	res := helper.DoValidation(req)

	if len(res) > 0 {
		log.Println("Error: Validation error")
		helper.ReturnJSON(ctx, http.StatusBadRequest, "Validation error", res)
		return
	}

	// check if email already exist
	checkByEmail, err := c.customerService.FindByEmail(req.Email)
	if err != nil {
		log.Println("Error:", err)
		if err.Error() == "record not found" {
			helper.ReturnJSON(ctx, http.StatusNotFound, "Customer not found", nil)
			return
		}
	}

	if checkByEmail != nil {
		log.Println("Error: Email already exist")
		helper.ReturnJSON(ctx, http.StatusBadRequest, "Email already exist", nil)
		return
	}

	customerId, err := c.customerService.UpdateById(intId, &req)

	if err != nil {
		log.Println("Error:", err)
		if err.Error() == "record not found" {
			helper.ReturnJSON(ctx, http.StatusNotFound, "Customer not found", nil)
			return
		}
		helper.ReturnJSON(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ReturnJSON(ctx, http.StatusOK, "Customer updated", response.UpdateCustomerResponse{Id: customerId.ID})
}

// DeleteCustomer godoc
// @Summary Soft Delete Customer
// @Description Soft Delete Customer
// @Tags Customer
// @Produce  json
// @Param id path int true "Customer ID"
// @Success 200 {object} response.ReturnResponseDelete
// @Failure 400 {object} response.ReturnResponseError
// @Failure 500 {object} response.ReturnResponseError.
// @Router api/v1/customer/{id} [delete]
func (c *CustomerController) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := helper.ConvertStringToInt(id)
	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSON(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = c.customerService.SoftDelete(intId)

	if err != nil {
		log.Println("Error:", err)
		if err.Error() == "record not found" {
			helper.ReturnJSON(ctx, http.StatusNotFound, "Customer not found", nil)
			return
		}
		helper.ReturnJSON(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ReturnJSON(ctx, http.StatusOK, "Customer deleted", response.DeleteCustomerResponse{Id: intId})
}

// FindCustomerById godoc
// @Summary Find Customer By Id
// @Description Find Customer By Id
// @Tags Customer
// @Produce  json
// @Param id path int true "Customer ID"
// @Success 200 {object} response.ReturnResponseGetById
// @Failure 400 {object} response.ReturnResponseError
// @Failure 500 {object} response.ReturnResponseError
// @Router api/v1/customer/{id} [get]

func (c *CustomerController) FindCustomerById(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := helper.ConvertStringToInt(id)
	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSON(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	customer, err := c.customerService.FindById(intId)

	if err != nil {
		log.Println("Error:", err)
		if err.Error() == "record not found" {
			helper.ReturnJSON(ctx, http.StatusNotFound, "Customer not found", nil)
			return
		}
		helper.ReturnJSON(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ReturnJSON(ctx, http.StatusOK, "Customer found", response.GetByIdCustomer{
		Id:      customer.ID,
		Name:    customer.Name,
		Email:   customer.Email,
		Address: customer.Address,
	})
}

// FindAllCustomer godoc
// @Summary Find All Customer
// @Description Find All Customer
// @Tags Customer
// @Produce  json
// @Param page query int false "Page"
// @Param page_size query int false "Page Size"
// @Param name query string false "Name"
// @Param email query string false "Email"
// @Param address query string false "Address"
// @Success 200 {object} response.ResponsePagination
// @Failure 400 {object} responseReturnResponseError
// @Failure 500 {object} responseReturnResponseError
// @Router api/v1/customer [get]
func (c *CustomerController) FindAllCustomer(ctx *gin.Context) {
	var returnMessage string
	var page int
	var pageSize int
	var name string
	var email string
	var address string

	pageQuery := ctx.DefaultQuery("page", "1")
	pageSizeQuery := ctx.DefaultQuery("page_size", "10")
	nameQuery := ctx.DefaultQuery("name", "")
	emailQuery := ctx.DefaultQuery("email", "")
	addressQuery := ctx.DefaultQuery("address", "")

	// min length nameQuery is 3
	if helper.MinLengthQueryParam(nameQuery, 3) {
		log.Println("Error: Name min length is 3")
		helper.ReturnJSON(ctx, http.StatusBadRequest, "You must enter at least 3 characters for name", nil)
		return
	}

	// min length emailQuery is 8
	if helper.MinLengthQueryParam(emailQuery, 3) {
		log.Println("Error: Email min length is 3")
		helper.ReturnJSON(ctx, http.StatusBadRequest, "You must enter at least 8 characters for email", nil)
		return
	}

	// min length addressQuery is 8
	if helper.MinLengthQueryParam(addressQuery, 3) {
		log.Println("Error: Address min length is 3")
		helper.ReturnJSON(ctx, http.StatusBadRequest, "You must enter at least 8 characters for address", nil)
		return
	}

	// convert string to int
	page, err := helper.ConvertStringToInt(pageQuery)
	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSON(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	pageSize, err = helper.ConvertStringToInt(pageSizeQuery)
	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSON(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	name = nameQuery
	email = emailQuery
	address = addressQuery

	filter := filters.CustomerFilter{
		Name:    name,
		Email:   email,
		Address: address,
	}

	customers, totalRecords, err := c.customerService.FindAll(&filter, page, pageSize)

	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSON(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	totalData, err := c.customerService.CountAll()

	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSON(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	if len(customers) == 0 {
		returnMessage = fmt.Sprintf("Customer not found")
	} else {
		returnMessage = fmt.Sprintf("Customer found")
	}

	helper.ReturnJSONWithMeta(ctx, http.StatusOK, returnMessage, customers, int(totalData), int(totalRecords), page, pageSize)
}
