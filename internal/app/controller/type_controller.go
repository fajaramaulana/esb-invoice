package controller

import (
	"esb-invoice/internal/app/handler/helper"
	"esb-invoice/internal/app/handler/request"
	"esb-invoice/internal/domain/model"
	"esb-invoice/internal/domain/service"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TypeController struct {
	typeService *service.TypeService
}

func NewTypeController(typeService *service.TypeService) *TypeController {
	return &TypeController{
		typeService: typeService,
	}
}

// FindAll godoc
// @Summary Get All Type
// @Description Get All Type
// @Tags Type
// @Produce  json
// @Success 200 {object} []response.Response
// @Failure 500 {object} response.ReturnResponseError
// @Router api/v1/type [get]
func (c *TypeController) FindAll(ctx *gin.Context) {
	types, err := c.typeService.FindAll()

	if err != nil {
		helper.ReturnJSON(ctx, 500, err.Error(), nil)
		return
	}

	helper.ReturnJSON(ctx, 200, "Success Get Data", types)
}

// Create godoc
// @Summary Create Type
// @Description Create Type
// @Tags Type
// @Accept  json
// @Produce  json
// @Param user body request.CreateTypeRequest true "Type"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.ReturnResponseError
// @Failure 500 {object} response.ReturnResponseError
// @Router api/v1/type [post]
func (c *TypeController) Create(ctx *gin.Context) {
	var req request.CreateTypeRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Printf("%# v\n", err.Error())
		if err.Error() == "EOF" {
			log.Println("Error: Request body is empty")
			helper.ReturnJSON(ctx, http.StatusBadRequest, "Request body is empty", nil)
			return
		}
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

	// check if type already exist
	findByName, err := c.typeService.FindByName(req.Name)
	if err != nil {
		if err.Error() != "record not found" {
			helper.ReturnJSON(ctx, 500, err.Error(), nil)
			return
		}
	}

	if findByName != nil {
		helper.ReturnJSON(ctx, 400, "Type already exist", nil)
		return
	}

	typeItem := &model.Type{
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Time{},
	}

	typeId, err := c.typeService.Create(typeItem)
	if err != nil {
		helper.ReturnJSON(ctx, 500, err.Error(), nil)
		return
	}

	helper.ReturnJSON(ctx, 201, "Success Create Data", typeId)
}

// Update godoc
// @Summary Update Type
// @Description Update Type
// @Tags Type
// @Accept  form-data
// @Produce  json
// @Param id path int true "Type ID"
// @Param user body request.UpdateTypeRequest true "Type"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ReturnResponseError
// @Failure 500 {object} response.ReturnResponseError
// @Router api/v1/type/{id} [put]
func (c *TypeController) Update(ctx *gin.Context) {
	var req request.UpdateTypeRequest

	id := ctx.Param("id")
	intId, err := helper.ConvertStringToInt(id)
	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSON(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Printf("%# v\n", err.Error())
		if err.Error() == "EOF" {
			log.Println("Error: Request body is empty")
			helper.ReturnJSON(ctx, http.StatusBadRequest, "Request body is empty", nil)
			return
		}
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

	// check if type already exist
	findByName, err := c.typeService.FindByNameAndNotId(req.Name, intId)
	if err != nil {
		if err.Error() != "record not found" {
			helper.ReturnJSON(ctx, 500, err.Error(), nil)
			return
		}
	}

	if findByName != nil {
		helper.ReturnJSON(ctx, 400, "Type already exist", nil)
		return
	}

	typeItem := &model.Type{
		Name:      req.Name,
		UpdatedAt: time.Now(),
	}

	typeId, err := c.typeService.UpdateById(intId, typeItem)
	if err != nil {
		if err.Error() == "record not found" {
			helper.ReturnJSON(ctx, http.StatusNotFound, "Type not found", nil)
			return
		}
		helper.ReturnJSON(ctx, 500, err.Error(), nil)
		return
	}

	helper.ReturnJSON(ctx, 200, "Success Update Data", typeId.ID)
}

// Delete godoc
// @Summary Delete Type
// @Description Delete Type
// @Tags Type
// @Produce  json
// @Param id path int true "Type ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.ReturnResponseError
// @Failure 500 {object} response.ReturnResponseError
// @Router api/v1/type/{id} [delete]

func (c *TypeController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	intId, err := helper.ConvertStringToInt(id)
	if err != nil {
		log.Println("Error:", err)
		helper.ReturnJSON(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = c.typeService.SoftDelete(intId)
	if err != nil {
		if err.Error() == "record not found" {
			helper.ReturnJSON(ctx, http.StatusNotFound, "Type not found", nil)
			return
		}
		helper.ReturnJSON(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	helper.ReturnJSON(ctx, http.StatusOK, "Success delete type", intId)
}
