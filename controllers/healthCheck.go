package controllers

import (
	"jual-beli-motor/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

var repo = repository.Database{}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func HealthCheck(ctx *gin.Context) {
	data := repo.GetHealthCheck(ctx)
	res := Response{
		Code:    http.StatusOK,
		Message: data,
	}
	ctx.JSON(res.Code, res)

}
