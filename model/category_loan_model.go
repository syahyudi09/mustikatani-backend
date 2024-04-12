package model

import (
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type CategoryLoanModel struct {
	Id               int       `json:"id"`
	CategoryLoanName string    `json:"category_loan_name" validate:"required"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type InsertCategoryLoan struct {
	Id               int    `json:"id"`
	CategoryLoanName string `json:"category_loan_name" validate:"required"`
}

func (c *CategoryLoanModel) Validate() error {
	validate := validator.New()
	err := validate.Struct(c)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		errMsg := ""
		for _, e := range errs {
			errMsg += fmt.Sprintf("Field %s: validation failed on tag '%s'\n", e.Field(), e.Tag())
		}
		return errors.New(errMsg)
	}

	return nil
}
