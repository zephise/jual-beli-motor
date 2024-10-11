package repository

import (
	"jual-beli-motor/models"
	"time"

	"github.com/gin-gonic/gin"
)

type Bike struct {
	Id          int        `gorm:"column:id" json:"id"`
	BikeTypeId  int        `gorm:"column:bike_types_id" json:"bike_types_id"`
	BikeType    BikeType   `gorm:"reference:bike_types_id;foreignKey:id" json:"bike_types"`
	UserID      int        `gorm:"column:user_id" json:"user_id"`
	User        User       `gorm:"references:user_id; foreignKey:id" json:"user"`
	Name        string     `gorm:"column:name" json:"name"`
	Description string     `gorm:"column:description" json:"description"`
	Price       int        `gorm:"column:price" json:"price"`
	Status      int        `gorm:"column:status" json:"status"`
	Image       string     `gorm:"column:image" json:"image"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Bike) TableName() string {
	return "bikes"
}

func CreateBike(ctx *gin.Context, data Bike) error {
	query := DB.Model(&data)
	query.Create(&data)

	return query.Error
}

func GetAllBike(ctx *gin.Context, params models.ReqParams) ([]Bike, error) {
	var data []Bike
	query := DB.Model(&data)
	query = query.Joins("BikeType")
	query = query.Joins("User")
	if params.Search != "" {
		query = query.Where("bikes.name like ?", "%"+params.Search+"%")
	}

	if params.Category != 0 {
		query = query.Where("BikeType.id = ?", params.Category)
	}

	query = query.Where("bikes.status = 1")

	query.Find(&data)

	return data, query.Error
}

func GetBikeById(ctx *gin.Context, id int) (Bike, error) {
	var data Bike

	query := DB.Model(&data)
	query = query.Where("id = ?", id)
	query.First(&data)

	return data, query.Error
}

func UpdateBike(ctx *gin.Context, data Bike, id int) error {
	query := DB.Model(&data)
	query = query.Where("id = ?", id)
	query.Updates(&data)

	return query.Error

}

func DeleteBike(ctx *gin.Context, id int) error {
	query := DB.Table("bikes")
	query = query.Where("id = ?", id)
	query.Delete(&Bike{})

	return query.Error
}

func PurchaseBike(ctx *gin.Context, id int) error {
	query := DB.Table("bikes")
	query = query.Where("id = ?", id)
	query.Update("status", "0")

	return query.Error
}
