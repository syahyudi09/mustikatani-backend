package model

import "time"

type CategoryLoanModel struct {
	Id               int       `json:"id"`
	CategoryLoanName string    `json:"category_loan_name"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}

type InsertCategoryLoan struct {
	Id               int    `json:"id"`
	CategoryLoanName string `json:"category_loan_name"`
}
