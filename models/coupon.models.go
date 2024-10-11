package models

type ReqCreateCoupon struct {
	Name        string `json:"name" validate:"required"`
	Quota       int    `json:"quota" validate:"required,numeric"`
	Percent     int    `json:"percent" validate:"required,numeric"`
	MinPurchase int    `json:"min_purchase" validate:"required"`
	MaxPurchase int    `json:"max_purchase" validate:"required"`
}
