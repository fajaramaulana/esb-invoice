package controller

import (
	"esb-invoice/internal/app/handler/filters"
	"esb-invoice/internal/app/handler/helper"
	"esb-invoice/internal/app/handler/request"
	"esb-invoice/internal/domain/model"
	"esb-invoice/internal/domain/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	itemService *service.ItemService
}

func NewItemController(itemService *service.ItemService) *ItemController {
	return &ItemController{
		itemService: itemService,
	}
}

// CreateItem godoc
// @Summary Create a new item
// @Description Create a new item
// @Tags Item
// @Accept  json
// @Produce  json
// @Param item body request.CreateItemRequest true "Item"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.ReturnResponseError
// @Failure 500 {object} response.ReturnResponseError
// @Router api/v1/item [post]
func (c *ItemController) CreateItem(ctx *gin.Context) {
	var req request.CreateItemRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		message, data := helper.GlobalCheckingErrorBindJson(err.Error())
		log.Println(fmt.Sprintf("Error: %s", message))
		helper.ReturnJSONError(ctx, http.StatusBadRequest, message, nil, data)
		return
	}

	res := helper.DoValidation(req)

	if len(res) > 0 {
		log.Println("Error: Validation error")
		helper.ReturnJSONError(ctx, http.StatusBadRequest, "Validation error", nil, res)
		return
	}

	item := model.Item{
		Name:   req.NameItem,
		TypeID: uint(req.TypeId),
		Price:  req.Price,
	}

	itemId, err := c.itemService.Create(&item)

	if err != nil {
		log.Println("Error: Create item error")
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, "Create item error", nil, err.Error())
		return
	}

	helper.ReturnJSON(ctx, http.StatusCreated, "Create item success", itemId)
}

// FindAllItem godoc
// @Summary Get All Item
// @Description Get All Item
// @Tags Item
// @Produce  json
// @Success 200 {object} response.ResponsePagination
// @Failure 500 {object} response.ReturnResponseError
// @Router api/v1/item [get]
func (c *ItemController) FindAllItem(ctx *gin.Context) {
	var returnMessage string
	var page int
	var pageSize int
	var name string

	pageQuery := ctx.DefaultQuery("page", "1")
	pageSizeQuery := ctx.DefaultQuery("page_size", "10")
	nameQuery := ctx.DefaultQuery("name", "")

	// min length nameQuery is 3
	if helper.MinLengthQueryParam(nameQuery, 3) {
		log.Println("Error: Name min length is 3")
		errMessage := "You must enter at least 3 characters for name"
		helper.ReturnJSONError(ctx, http.StatusBadRequest, errMessage, nil, map[string]interface{}{"error": errMessage})
		return
	}

	// convert string to int
	page, err := helper.ConvertStringToInt(pageQuery)
	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSONError(ctx, http.StatusBadRequest, err.Error(), nil, map[string]interface{}{"error": err.Error()})
		return
	}

	pageSize, err = helper.ConvertStringToInt(pageSizeQuery)
	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSONError(ctx, http.StatusBadRequest, err.Error(), nil, map[string]interface{}{"error": err.Error()})
		return
	}

	name = nameQuery

	filter := filters.ItemFilter{
		Name: name,
	}

	items, totalRecords, err := c.itemService.FindAll(&filter, page, pageSize)

	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, err.Error(), nil, map[string]interface{}{"error": err.Error()})
		return
	}

	totalData, err := c.itemService.CountAll()

	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, err.Error(), nil, map[string]interface{}{"error": err.Error()})
		return
	}

	if len(items) == 0 {
		returnMessage = "Data not found"
	} else {
		returnMessage = "Success get data"
	}

	helper.ReturnJSONWithMeta(ctx, http.StatusOK, returnMessage, items, int(totalData), int(totalRecords), page, pageSize)
}

// FindByIdItem godoc
// @Summary Get Item By Id
// @Description Get Item By Id
// @Tags Item
// @Produce  json
// @Param id path int true "Id Item"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ReturnResponseError
// @Failure 500 {object} response.ReturnResponseError
// @Router api/v1/item/{id} [get]

func (c *ItemController) FindByIdItem(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := helper.ConvertStringToInt(id)
	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSONError(ctx, http.StatusBadRequest, err.Error(), nil, map[string]interface{}{"error": err.Error()})
		return
	}

	item, err := c.itemService.FindById(intId)

	if err != nil {
		log.Println("Error:", err)
		if err.Error() == "record not found" {
			helper.ReturnJSONError(ctx, http.StatusNotFound, "Item not found", nil, map[string]interface{}{"error": "Item not found"})
			return
		}
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, err.Error(), nil, map[string]interface{}{"error": err.Error()})
		return
	}

	helper.ReturnJSON(ctx, http.StatusOK, "Item found", item)
}

// UpdateItem godoc
// @Summary Update Item
// @Description Update Item
// @Tags Item
// @Accept  json
// @Produce  json
// @Param id path int true "Id Item"
// @Param item body request.UpdateItemRequest true "Item"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ReturnResponseError
// @Failure 500 {object} response.ReturnResponseError
// @Router api/v1/item/{id} [put]
func (c *ItemController) UpdateItem(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := helper.ConvertStringToInt(id)
	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSONError(ctx, http.StatusBadRequest, err.Error(), nil, map[string]interface{}{"error": err.Error()})
		return
	}

	var req request.UpdateItemRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		message, data := helper.GlobalCheckingErrorBindJson(err.Error())
		log.Println(fmt.Sprintf("Error: %s", message))
		helper.ReturnJSONError(ctx, http.StatusBadRequest, message, nil, data)
		return
	}

	res := helper.DoValidation(req)

	if len(res) > 0 {
		log.Println("Error: Validation error")
		helper.ReturnJSONError(ctx, http.StatusBadRequest, "Validation error", nil, res)
		return
	}

	item := model.Item{
		Name:   req.NameItem,
		TypeID: uint(req.TypeId),
		Price:  req.Price,
	}

	itemUpdate, err := c.itemService.UpdateById(intId, &item)

	if err != nil {
		log.Println("Error: Update item error")
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, "Update item error", nil, err.Error())
		return
	}

	helper.ReturnJSON(ctx, http.StatusOK, "Update item success", itemUpdate.ID)
}

// DeleteItem godoc
// @Summary Delete Item
// @Description Delete Item
// @Tags Item
// @Produce  json
// @Param id path int true "Id Item"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ReturnResponseError
// @Failure 500 {object} response.ReturnResponseError
// @Router api/v1/item/{id} [delete]
func (c *ItemController) DeleteItem(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := helper.ConvertStringToInt(id)
	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSONError(ctx, http.StatusBadRequest, "Error Parsing Id", nil, map[string]interface{}{"error": err.Error()})
		return
	}

	err = c.itemService.SoftDelete(intId)

	if err != nil {
		log.Println("Error:", err)
		if err.Error() == "record not found" {
			helper.ReturnJSONError(ctx, http.StatusNotFound, "Item not found", nil, map[string]interface{}{"error": "Item not found"})
			return
		}
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, err.Error(), nil, map[string]interface{}{"error": err.Error()})
		return
	}

	helper.ReturnJSON(ctx, http.StatusOK, "Delete item success", intId)
}
