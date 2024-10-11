package repository

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Coupon struct {
	Id          int        `gorm:"column:id" json:"id"`
	Name        string     `gorm:"column:name" json:"name"`
	Percent     int        `gorm:"column:persen" json:"persen"`
	MinPurchase int        `gorm:"column:min_purchase" json:"min_purchase"`
	MaxPurchase int        `gorm:"column:max_purchase" json:"max_purchase"`
	Quota       int        `gorm:"column:quota" json:"quota"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Coupon) TableName() string {
	return "coupon"
}

func CreateCoupon(ctx *gin.Context, data Coupon) error {
	query := DB.Model(&data)
	query.Create(&data)

	return query.Error
}

func GetAllCoupon(ctx *gin.Context) ([]Coupon, error) {
	var data []Coupon
	query := DB.Model(&data)
	query.Find(&data)

	return data, query.Error
}
