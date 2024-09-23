package models

type ReqBike struct {
	BikeTypeId  int    `json:"bike_type_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Image       string `json:"image"`
}

type ReqParams struct {
	Search    string `form:"search"`
	Limit     int    `form:"limit"`
	PageIndex int    `form:"page_index"`
	Category  int    `form:"category"`
}

type ReqUpdateBike struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Price       int    `json:"price" validate:"required"`
}
