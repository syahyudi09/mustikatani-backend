package model

import "time"

type GoodsModel struct {
	Id				int			`json:"id"`
	CustomerId      int			`json:"customer_id"`
	CategoryIdLoan  int			`json:"category_loan_id"`
	ProductId       int			`json:"product_id"`
	LoadDate		time.Time	`json:"loan_date"`
	PaymentDate		time.Time	`json:"payment_date"`
	DueDate			time.Time	`json:"due_date"`
	Quantity		int			`json:"quantity"`
	Price			float64		`json:"price"`
	Amount			float64		`json:"amount"`
	Status 			 string		`json:"status"`
	RepaymentStatus StatusEnum `json:"repayment_status"`
	CreateAt		time.Time	`json:"created_at"`
	UpdateAt		time.Time	`json:"updated_at"`
}

type LoanGoodsModel struct{
		Id               int        `json:"id"`
		CustomerId       int        `json:"customer_id"`
		LoanDate         time.Time  `json:"loan_date"`
		DueDate          time.Time  `json:"due_date"`
		CategoryLoanID   int        `json:"category_loan_id"`
		ProductId       int			`json:"product_id"`
		ProductName      string    `json:"product_name"`
		Quantity		int			`json:"quantity"`
		Price			float64		`json:"price"`
		Amount           float64    `json:"amount"`
		Description      string     `json:"description"`
		Status           string     `json:"status"`
		RepaymentStatus  StatusEnum `json:"repayment_status"`
		CreatedAt        time.Time  `json:"created_at"`
		UpdatedAt        time.Time  `json:"updated_at"`
		NoKK             string     `json:"nokk"`
		NIK              string     `json:"nik"`
		FullName         string     `json:"full_name"`
		Address          string     `json:"address"`
		PhoneNumber      string     `json:"phone"`
		EmergencyName    string     `json:"emergencyname"`
		EmergencyContact string     `json:"emergencycontact"`
		LastSalary       float64    `json:"last_salary"`
}

