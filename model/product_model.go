package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type ProductModel struct {
	Id                  int       `json:"id"`
	ProductName         string    `json:"product_name" validate:"required,min=3" `
	Description         string    `json:"description"`
	Price               float64   `json:"price" validate:"required,gte=0"`
	Stok                int       `json:"stok" validate:"required"`
	CategoryProductId   int       `json:"category_product_id" validate:"required"`
	CategoryProductName string    `json:"category_product_name"`
	Status              bool      `json:"status" validate:"required"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type CreateProductRequest struct {
	Id                int     `json:"id"`
	ProductName       string  `json:"product_name" validate:"required,min=3"`
	Description       string  `json:"description"`
	Price             float64 `json:"price" validate:"required,gte=0"`
	Stok              int     `json:"stok" validate:"required"`
	CategoryProductId int     `json:"category_product_id" validate:"required"`
	Status            bool    `json:"status" validate:"required"`
}

func (p *ProductModel) ValidateUpdate() error {
	validate := validator.New()
	err := validate.Struct(p)
	if err != nil {
		// Mengembalikan error dengan pesan validasi yang lebih spesifik
		errs := err.(validator.ValidationErrors)
		errMsg := ""
		for _, e := range errs {
			errMsg += fmt.Sprintf("Field %s: validation failed on tag '%s'\n", e.Field(), e.Tag())
		}
		return errors.New(errMsg)
	}

	return nil
}
