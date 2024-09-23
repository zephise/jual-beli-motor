package middleware

import (
	"fmt"
	"jual-beli-motor/models"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(level string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")

		tokenString := strings.Split(auth, " ")

		if tokenString[0] != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization Failed",
			})
			ctx.Abort()
		}

		if len(tokenString) < 2 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization Failed",
			})
			ctx.Abort()
		}

		claims := &models.ClaimJwt{}
		token, err := jwt.ParseWithClaims(tokenString[1], claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization Failed",
			})
			ctx.Abort()
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "Expired Token",
			})
			ctx.Abort()
		}
		ctx.Set("user", claims)
		ctx.Next()

	}
}
