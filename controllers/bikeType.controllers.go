package controllers

import (
	"jual-beli-motor/models"
	"jual-beli-motor/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ReqBikeType struct {
	Name string `json:"name"`
}

func CreatedBikeType(ctx *gin.Context) {
	var res Response
	var payload ReqBikeType

	if err := ctx.ShouldBind(&payload); err != nil {
		logrus.Println("Bad Request Error", err)
		logrus.Println("Error request: ", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	data := repository.BikeType{
		Name:      payload.Name,
		UpdatedAt: nil,
	}

	if err := repository.CreateBikeType(ctx, data); err != nil {
		logrus.Println("Error create bike type: ", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Failed Create Data"

		ctx.JSON(res.Code, res)
		return
	}

	res.Code = http.StatusOK
	res.Message = "Success Create Data"
	ctx.JSON(res.Code, res)
}

func GetAllBikeType(ctx *gin.Context) {
	var res Response
	var param models.ReqBikeParams

	if err := ctx.BindQuery(&param); err != nil {
		logrus.Println("Error bind query", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	data, err := repository.GetAllBikeType(ctx, param)

	if err != nil {
		logrus.Println("Error get data:", err)
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

func GetBikeTypeById(ctx *gin.Context) {
	var res Response

	id, _ := strconv.Atoi(ctx.Param("id"))

	data, err := repository.GetBikeTypeById(ctx, id)

	if err != nil {
		logrus.Println("Data not found")
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		ctx.JSON(res.Code, res)
		return
	}

	res.Code = http.StatusOK
	res.Message = "Success Get Data"
	res.Data = data
	ctx.JSON(res.Code, res)
}

func UpdateBikeType(ctx *gin.Context) {

	var res Response
	var payload ReqBikeType

	id, _ := strconv.Atoi(ctx.Param("id"))

	if err := ctx.ShouldBind(&payload); err != nil {
		logrus.Println("Bad Request")
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"
		ctx.JSON(res.Code, res)
		return
	}

	data, err := repository.GetBikeTypeById(ctx, id)

	if err != nil {
		logrus.Println("Data not found")
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"
		ctx.JSON(res.Code, res)
		return
	}

	data.Name = payload.Name

	if err := repository.UpdateBikeType(ctx, data, id); err != nil {
		logrus.Println("Failed Update Bike Type:", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Failed Update Data"

		ctx.JSON(res.Code, res)
		return
	}

	res.Code = http.StatusOK
	res.Message = "Success Update Data"
	ctx.JSON(res.Code, res)
}

func DeleteBikeTypeById(ctx *gin.Context) {
	var res Response
	id, _ := strconv.Atoi(ctx.Param("id"))

	_, err := repository.GetBikeTypeById(ctx, id)

	if err != nil {
		res.Code = http.StatusNotFound
		res.Message = "Data Not Found"

		ctx.JSON(res.Code, res)
		return
	}

	if err := repository.DeleteBikeType(ctx, id); err != nil {
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Failed Delete Data"

		ctx.JSON(res.Code, res)
		return
	}

	res.Code = http.StatusOK
	res.Message = "Success Delete Data"
	ctx.JSON(res.Code, res)
}
