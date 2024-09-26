package models

import "time"

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

type ResBike struct {
	Id           int        `json:"id"`
	BikeTypeId   int        `json:"bike_types_id"`
	BikeTypeName string     `json:"bike_types_name"`
	UserID       int        `json:"user_id"`
	UserName     string     `json:"user_name"`
	Name         string     `json:"name"`
	Description  string     `json:"description"`
	Price        int        `json:"price"`
	Status       int        `json:"status"`
	Image        string     `json:"image"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at"`
}
