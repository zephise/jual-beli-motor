package repository

import (
	"time"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id        int       `gorm:"column:id" json:"id"`
	Username  string    `gorm:"column:username" json:"username"`
	Password  string    `gorm:"column:password" json:"password"`
	Email     string    `gorm:"column:email" json:"email"`
	Type      int       `gorm:"column:type" json:"type"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (u *User) GetTypeUser() string {
	userType := map[int]string{
		0: "user",
		1: "admin",
	}

	return userType[u.Type]
}

func (User) TableName() string {
	return "user"
}

func GetUserByEmail(ctx *gin.Context, email string) (User, error) {
	var data User
	query := DB.Model(&data)
	query = query.Where("email = ?", email)
	query.First(&data)

	return data, query.Error

}

func CreateUser(ctx *gin.Context, data User) error {
	query := DB.Model(&data)
	query.Save(&data)

	return query.Error
}
