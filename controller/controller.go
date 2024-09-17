package controller

import (
	"CRUD/data/request"
	"CRUD/data/response"
	"CRUD/helper"
	"CRUD/service"
	"github.com/gin-gonic/gin"

	"net/http"
	"strconv"
)

type ComputadorController struct {
	computadorService service.ComputadorService
}

func NewComputadorController(service service.ComputadorService) *ComputadorController {
	return &ComputadorController{computadorService: service}
}

func (controller *ComputadorController) Create(ctx *gin.Context) {
	createComputadorRequest := request.CreateComputadorRequest{}
	err := ctx.ShouldBindJSON(&createComputadorRequest)
	helper.ErrorPanic(err)

	controller.computadorService.Create(createComputadorRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ComputadorController) Update(ctx *gin.Context) {
	updateComputadorRequest := request.UpdateComputadorRequest{}
	err := ctx.ShouldBindJSON(&updateComputadorRequest)
	helper.ErrorPanic(err)

	computadorId := ctx.Param("computadorId")
	id, err := strconv.Atoi(computadorId)
	helper.ErrorPanic(err)

	updateComputadorRequest.Id = id

	controller.computadorService.Update(updateComputadorRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ComputadorController) Delete(ctx *gin.Context) {
	computadorId := ctx.Param("computadorId")
	id, err := strconv.Atoi(computadorId)
	helper.ErrorPanic(err)
	controller.computadorService.Delete(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)

}

func (controller *ComputadorController) FindById(ctx *gin.Context) {
	computadorId := ctx.Param("computadorId")
	id, err := strconv.Atoi(computadorId)
	helper.ErrorPanic(err)

	computadorResponse := controller.computadorService.FindById(id)

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   computadorResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ComputadorController) FindAll(ctx *gin.Context) {
	comptadorResponse := controller.computadorService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "Ok",
		Data:   comptadorResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)

}
