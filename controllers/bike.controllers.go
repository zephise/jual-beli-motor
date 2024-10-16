package controllers

import (
	"fmt"
	"jual-beli-motor/helper"
	"jual-beli-motor/models"
	"jual-beli-motor/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateBike(ctx *gin.Context) {
	res := Response{}
	payload := models.ReqBike{}

	if err := ctx.ShouldBind(&payload); err != nil {
		logrus.Println("Bad Request", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	data := repository.Bike{
		BikeTypeId:  payload.BikeTypeId,
		Name:        payload.Name,
		Description: payload.Description,
		Price:       payload.Price,
	}

	if err := repository.CreateBike(ctx, data); err != nil {
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Failed Create Bike"

		ctx.JSON(res.Code, res)
		return
	}

	res.Code = http.StatusOK
	res.Message = "Successfully Create Bike"

	ctx.JSON(res.Code, res)
}

func GetAllBike(ctx *gin.Context) {
	res := Response{}
	params := models.ReqParams{}

	if err := ctx.ShouldBindQuery(&params); err != nil {
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	data, err := repository.GetAllBike(ctx, params)

	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = "Server Error"

		ctx.JSON(res.Code, res)
		return
	}

	res.Code = http.StatusOK
	res.Message = "Success Get Data"
	res.Data = data

	ctx.JSON(res.Code, res)
}

func UpdateBike(ctx *gin.Context) {

	var res Response
	var payload models.ReqUpdateBike

	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBind(&payload); err != nil {
		fmt.Println("Bad Request", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"
		ctx.JSON(res.Code, res)
		return
	}

	if err := helper.Validate(payload); err != nil {
		res.Code = http.StatusBadRequest
		res.Message = err.Error()

		ctx.JSON(res.Code, res)
		return
	}

	data, err := repository.GetBikeById(ctx, id)

	if err != nil {
		fmt.Println("Data not found", err)
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		ctx.JSON(res.Code, res)
		return
	}

	data.Name = payload.Name
	data.Description = payload.Description
	data.Price = payload.Price

	if err := repository.UpdateBike(ctx, data, id); err != nil {
		fmt.Println("Failed Update Bike:", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Failed Update Data"

		ctx.JSON(res.Code, res)
		return
	}

	res.Code = http.StatusOK
	res.Message = "Success Update Data"
	ctx.JSON(res.Code, res)
}

func GetBikeDetail(ctx *gin.Context) {
	var res Response

	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := repository.GetBikeById(ctx, id)

	if err != nil {
		fmt.Println("Data not found", err)
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		ctx.JSON(res.Code, res)
		return
	}

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = data

	ctx.JSON(res.Code, res)
}

func DeleteBikeById(ctx *gin.Context) {
	var res Response
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err := repository.GetBikeById(ctx, id)

	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		ctx.JSON(res.Code, res)
		return
	}

	if err := repository.DeleteBike(ctx, id); err != nil {
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Failed Delete Data"
		ctx.JSON(res.Code, res)
		return
	}

	res.Code = http.StatusOK
	res.Message = "Success Delete Data"
	ctx.JSON(res.Code, res)
}
