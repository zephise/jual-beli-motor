package controllers

import (
	"jual-beli-motor/helper"
	"jual-beli-motor/models"
	"jual-beli-motor/repository"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type ReqGetUserByEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context) {
	var payload ReqGetUserByEmail
	var res Response
	if err := ctx.BindJSON(&payload); err != nil {
		logrus.Println("Bad Request", err)
		res.Code = http.StatusOK
		res.Message = "Bad Request"
		ctx.JSON(res.Code, res)
		return
	}

	user, err := repository.GetUserByEmail(ctx, payload.Email)

	if err != nil {
		logrus.Println("User Not Found")
		res.Code = http.StatusNotFound
		res.Message = "User Not Found"
		ctx.JSON(res.Code, res)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
		logrus.Println("Wrong Password")
		res.Code = http.StatusUnauthorized
		res.Message = "Wrong Password"

		ctx.JSON(res.Code, res)
		return
	}

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := models.ClaimJwt{
		Id:   user.Id,
		Name: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_KEY")))

	res.Code = http.StatusOK
	res.Message = "Success"
	res.Data = map[string]string{
		"token": tokenString,
	}

	ctx.JSON(res.Code, res)

}

func CreateUserNonAdmin(ctx *gin.Context) {
	res := Response{}
	payload := models.ReqUser{}

	if err := ctx.ShouldBind(&payload); err != nil {
		logrus.Println("Bad Request", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	if err := helper.Validate(payload); err != nil {
		logrus.Println("Bad Request", err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()

		ctx.JSON(res.Code, res)
		return
	}

	existEmail, _ := repository.GetUserByEmail(ctx, payload.Email)

	if existEmail.Id != 0 {
		logrus.Println("Email Already Registered")
		res.Code = http.StatusBadRequest
		res.Message = "Email Already Registered"

		ctx.JSON(res.Code, res)
		return
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Email), bcrypt.DefaultCost)

	if err != nil {
		logrus.Println("Encryption Error", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Internal Server"

		ctx.JSON(res.Code, res)
		return
	}

	newUser := repository.User{
		Username: payload.Name,
		Email:    payload.Email,
		Password: string(newPassword),
		Type:     0,
	}

	if err := repository.CreateUser(ctx, newUser); err != nil {
		logrus.Println("Failed Create User", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Failed Create User"

		ctx.JSON(res.Code, res)
	}
	res.Code = http.StatusCreated
	res.Message = "Success Create User"
	ctx.JSON(res.Code, res)
}

func CreateUserAdmin(ctx *gin.Context) {
	res := Response{}
	payload := models.ReqUser{}

	if err := ctx.ShouldBind(&payload); err != nil {
		logrus.Println("Bad Request", err)
		res.Code = http.StatusBadRequest
		res.Message = "Bad Request"

		ctx.JSON(res.Code, res)
		return
	}

	if err := helper.Validate(payload); err != nil {
		logrus.Println("Bad Request", err)
		res.Code = http.StatusBadRequest
		res.Message = err.Error()

		ctx.JSON(res.Code, res)
		return
	}

	existEmail, _ := repository.GetUserByEmail(ctx, payload.Email)

	if existEmail.Id != 0 {
		logrus.Println("Email Already Registered")
		res.Code = http.StatusBadRequest
		res.Message = "Email Already Registered"

		ctx.JSON(res.Code, res)
		return
	}

	newPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Email), bcrypt.DefaultCost)

	if err != nil {
		logrus.Println("Encryption Error", err)
		res.Code = http.StatusInternalServerError
		res.Message = "Internal Server"

		ctx.JSON(res.Code, res)
		return
	}

	logrus.Println(payload.Password)

	newUser := repository.User{
		Username: payload.Name,
		Email:    payload.Email,
		Password: string(newPassword),
		Type:     1,
	}

	if err := repository.CreateUser(ctx, newUser); err != nil {
		logrus.Println("Failed Create User", err)
		res.Code = http.StatusUnprocessableEntity
		res.Message = "Failed Create User"

		ctx.JSON(res.Code, res)
	}
	res.Code = http.StatusCreated
	res.Message = "Success Create User"
	ctx.JSON(res.Code, res)
}
