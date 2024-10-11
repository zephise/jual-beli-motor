package controllers

import (
	"jual-beli-motor/models"
	"jual-beli-motor/repository"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func CreateCoupon(ctx *gin.Context) {
	res := Response{}
	payload := models.ReqCreateCoupon{}

	if err := ctx.ShouldBind(&payload); err != nil {
		logrus.Println("Bad Request", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	data := repository.Coupon{
		Name:        payload.Name,
		Percent:     payload.Percent,
		MinPurchase: payload.MinPurchase,
		MaxPurchase: payload.MaxPurchase,
		Quota:       payload.Quota,
	}

	if err := repository.CreateCoupon(ctx, data); err != nil {
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Failed Create Coupon"

		ctx.JSON(res.Code, res)
		return
	}

	res.Code = http.StatusOK
	res.Message = "Successfully Create Bike"

	ctx.JSON(res.Code, res)
}
