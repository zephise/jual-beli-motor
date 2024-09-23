package repository

import (
	"jual-beli-motor/models"
	"time"

	"github.com/gin-gonic/gin"
)

type BikeType struct {
	Id        int        `gorm:"column:id" json:"id"`
	Name      string     `gorm:"column:name" json:"name"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (BikeType) TableName() string {
	return "bike_types"
}

func CreateBikeType(ctx *gin.Context, data BikeType) error {
	query := DB.Model(&data)
	query.Create(&data)

	return query.Error
}

func GetAllBikeType(ctx *gin.Context, param models.ReqBikeParams) ([]BikeType, error) {
	var data []BikeType
	query := DB.Model(&data)

	if param.Search != "" {
		query = query.Where("name like ?", "%"+param.Search+"%")
	}

	query.Find(&data)
	return data, query.Error
}

func GetBikeTypeById(ctx *gin.Context, id int) (BikeType, error) {
	var data BikeType

	query := DB.Model(&data)
	query = query.Where("id = ?", id)
	query.First(&data)

	return data, query.Error
}

func UpdateBikeType(ctx *gin.Context, data BikeType, id int) error {
	query := DB.Model(&data)
	query = query.Where("id = ?", id)
	query.Updates(&data)

	return query.Error
}

func DeleteBikeType(ctx *gin.Context, id int) error {
	query := DB.Table("bike_types")
	query = query.Where("id = ?", id)
	query.Delete(&BikeType{})

	return query.Error
}
