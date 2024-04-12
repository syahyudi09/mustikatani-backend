package model

import "time"

type CategoryProductModel struct {
	Id 						int `json:"id"`
	CategoryProductName  	string `json:"category_product_name"`
	CreateAt				time.Time `json:"create_at"`
	UpdateAt 				time.Time `json:"update_at"`
}