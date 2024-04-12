package repository

import (
	"database/sql"
	"fmt"
	"pinjam-modal-app/model"
	"pinjam-modal-app/utils"
	"time"
)

type LoanApplicationRepo interface {
	CreateLoanApplication(application *model.LoanApplicationModel) error
	GetCustomerById(int) (*model.ValidasiCustomerModel, error)
	GetLoanApplications(page, limit int) ([]*model.LoanApplicationJoinModel, error)
	GetLoanApplicationById(id int) (*model.LoanApplicationJoinModel, error)
	LoanRepayment(id int, repayment *model.LoanRepaymentModel) error
	GetLoanApplicationRepaymentStatus(page, limit int, repaymentStatus model.StatusEnum) ([]*model.LoanApplicationJoinModel, error)
	GetLoanRepaymentsByDateRange(startDate time.Time, endDate time.Time) ([]*model.LoanRepaymentModel, error)
}

type loanApplicationRepo struct {
	db *sql.DB
}

func (r *loanApplicationRepo) CreateLoanApplication(application *model.LoanApplicationModel) error {
	insertStatement := utils.CREATE_APLICATION_LOAN_REPO
	application.LoanDate = time.Now()
	err := r.db.QueryRow(insertStatement, application.CustomerId, application.LoanDate, application.DueDate, application.CategoryLoanId, application.Amount, application.Description, application.Status, application.RepaymentStatus, application.CreatedAt, application.UpdatedAt).Scan(&application.Id)
	if err != nil {
		return fmt.Errorf("error on loanApplicationRepo.CreateLoanApplication: %w", err)
	}

	return nil
}

func (r *loanApplicationRepo) GetCustomerById(id int) (*model.ValidasiCustomerModel, error) {
	qry := utils.GET_CUSTOMER_LOAN_BY_ID

	customer := &model.ValidasiCustomerModel{}
	err := r.db.QueryRow(qry, id).Scan(
		&customer.Id, &customer.NIK, &customer.NoKK, &customer.EmergencyName, &customer.EmergencyContact, &customer.LastSalary)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("customer not found")
		}
		return nil, fmt.Errorf("error on GetCustomerById: %w", err)
	}

	return customer, nil
}

func (r *loanApplicationRepo) GetLoanApplications(page, limit int) ([]*model.LoanApplicationJoinModel, error) {
	offset := (page - 1) * limit

	selectStatement := utils.GET_LOAN_APLICATION
	

	rows, err := r.db.Query(selectStatement, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get loan applications: %w", err)
	}
	defer rows.Close()

	applications := []*model.LoanApplicationJoinModel{}
	for rows.Next() {
		application := &model.LoanApplicationJoinModel{}
		err := rows.Scan(
			&application.Id, &application.CustomerId, &application.LoanDate, &application.DueDate, &application.CategoryLoanID,
			&application.Amount, &application.Description, &application.Status, &application.RepaymentStatus, &application.CreatedAt, &application.UpdatedAt,
			&application.FullName, &application.Address, &application.NIK, &application.PhoneNumber, &application.NoKK, &application.EmergencyName,
			&application.EmergencyContact, &application.LastSalary,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan loan application: %w", err)
		}
		applications = append(applications, application)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get loan applications: %w", err)
	}

	return applications, nil
}

func (r *loanApplicationRepo) GetLoanApplicationById(id int) (*model.LoanApplicationJoinModel, error) {
	selectStatement := utils.GET_LOAN_APLICATION_BY_ID

	loan := &model.LoanApplicationJoinModel{}
	err := r.db.QueryRow(selectStatement, id).Scan(
		&loan.Id, &loan.CustomerId, &loan.LoanDate, &loan.DueDate, &loan.CategoryLoanID,
		&loan.Amount, &loan.Description, &loan.Status, &loan.RepaymentStatus, &loan.CreatedAt, &loan.UpdatedAt,
		&loan.FullName, &loan.Address, &loan.NIK, &loan.PhoneNumber, &loan.NoKK, &loan.EmergencyName,
		&loan.EmergencyContact, &loan.LastSalary,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("loan application not found")
		}
		return nil, fmt.Errorf("failed to get loan application: %w", err)
	}

	return loan, nil
}

func (r *loanApplicationRepo) GetLoanApplicationRepaymentStatus(page, limit int, repaymentStatus model.StatusEnum) ([]*model.LoanApplicationJoinModel, error) {
	offset := (page - 1) * limit

	selectStatement := utils.GET_LOAN_APLCATION_REPAYMENT_STATUS

	rows, err := r.db.Query(selectStatement, offset, limit, repaymentStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to get loan applications: %w", err)
	}
	defer rows.Close()

	applications := []*model.LoanApplicationJoinModel{}
	for rows.Next() {
		application := &model.LoanApplicationJoinModel{}
		err := rows.Scan(
			&application.Id, &application.CustomerId, &application.LoanDate, &application.DueDate, &application.CategoryLoanID,
			&application.Amount, &application.Description, &application.Status, &application.RepaymentStatus, &application.CreatedAt, &application.UpdatedAt,
			&application.FullName, &application.Address, &application.NIK, &application.PhoneNumber, &application.NoKK, &application.EmergencyName,
			&application.EmergencyContact, &application.LastSalary,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan loan application: %w", err)
		}
		applications = append(applications, application)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get loan applications: %w", err)
	}

	return applications, nil
}

func (r *loanApplicationRepo) LoanRepayment(id int, repayment *model.LoanRepaymentModel) error {
	updateStatment := utils.LOAN_REPAYMENT
	_, err := r.db.Exec(updateStatment, repayment.PaymentDate, repayment.Payment, model.StatusEnum(repayment.RepaymentStatus), repayment.UpdatedAt, id)
	if err != nil {
		return fmt.Errorf("error on loanApplicationRepo.LoanRepayment() : %w", err)
	}
	return nil
}

func (r *loanApplicationRepo) GetLoanRepaymentsByDateRange(startDate time.Time, endDate time.Time) ([]*model.LoanRepaymentModel, error) {
	selectStatement := utils.GET_LOAN_REPAYMENT_BY_DATE_RANGE

	rows, err := r.db.Query(selectStatement, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("error querying loan repayments: %w", err)
	}
	defer rows.Close()

	loanRepayments := []*model.LoanRepaymentModel{}
	for rows.Next() {
		loanRepayment := &model.LoanRepaymentModel{}
		err := rows.Scan(&loanRepayment.PaymentDate, &loanRepayment.Payment)
		if err != nil {
			return nil, fmt.Errorf("error scanning loan repayment: %w", err)
		}
		loanRepayments = append(loanRepayments, loanRepayment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error retrieving loan repayments: %w", err)
	}

	return loanRepayments, nil
}

func NewLoanApplicationRepository(db *sql.DB) LoanApplicationRepo {
	return &loanApplicationRepo{
		db: db,
	}
}
