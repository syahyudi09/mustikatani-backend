package model

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type CustomerModel struct {
	Id               int       `json:"id"`
	UserId           int       `json:"user_id"`
	FullName         string    `json:"full_name" validate:"required,min=3"`
	Address          string    `json:"address"`
	NIK              string    `json:"nik" validate:"required,min=16,max=16"`
	Phone            string    `json:"phone" validate:"required"`
	NoKK             string    `json:"no_kk" validate:"required,min=16,max=16"`
	EmergencyName    string    `json:"emergency_name" validate:"required"`
	EmergencyContact string    `json:"emergency_contact" validate:"required"`
	LastSalary       float64   `json:"last_salary" validate:"required,gte=0"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type ValidasiCustomerModel struct {
	Id               int
	NIK              sql.NullString
	NoKK             sql.NullString
	EmergencyName    sql.NullString
	EmergencyContact sql.NullString
	LastSalary       sql.NullFloat64
}

func (c *CustomerModel) Validate() error {
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
