package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type ClaimJwt struct {
	jwt.RegisteredClaims
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ReqUser struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
