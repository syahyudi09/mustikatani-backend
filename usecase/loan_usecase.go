package usecase

import (
	"fmt"
	"log"
	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"
	"time"
)

type LoanApplicationUsecase interface {
	CreateLoanApplication(application *model.LoanApplicationModel) error
	GetLoanApplications(page, limit int) ([]*model.LoanApplicationJoinModel, error)
	GetLoanApplicationById(id int) (*model.LoanApplicationJoinModel, error)
	LoanRepayment(id int, repayment *model.LoanRepaymentModel) error
	GetLoanApplicationRepaymentStatus(page, limit int, repaymentStatus model.StatusEnum) ([]*model.LoanApplicationJoinModel, error)
	GenerateIncomeReport(startDate time.Time, endDate time.Time) ([]*model.LoanRepaymentModel, float64, error)
}

type loanApplicationUsecase struct {
	repo repository.LoanApplicationRepo
}

func (uc *loanApplicationUsecase) CreateLoanApplication(application *model.LoanApplicationModel) error {

	customerDB, err := uc.repo.GetCustomerById(application.CustomerId)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	if customerDB.NIK.Valid && customerDB.NIK.String != "" &&
		customerDB.NoKK.Valid && customerDB.NoKK.String != "" &&
		customerDB.EmergencyName.Valid && customerDB.EmergencyName.String != "" &&
		customerDB.EmergencyContact.Valid && customerDB.EmergencyContact.String != "" &&
		customerDB.LastSalary.Valid && customerDB.LastSalary.Float64 != 0 {
		application.Status = model.LoanStatusApprove
		application.RepaymentStatus = model.StatusEnum(model.RepaymentStatusBelumLunas)
		application.DueDate = time.Now().AddDate(0, 2, 0)
	} else {
		application.Status = model.LoanStatusDenied
		fmt.Println("Silakan lengkapi data customer")
	}

	err = uc.repo.CreateLoanApplication(application)
	if err != nil {
		return fmt.Errorf("failed to insert loan: %v", err)
	}

	return nil
}

func (uc *loanApplicationUsecase) GetLoanApplications(page, limit int) ([]*model.LoanApplicationJoinModel, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	return uc.repo.GetLoanApplications(page, limit)
}

func (uc *loanApplicationUsecase) GetLoanApplicationById(id int) (*model.LoanApplicationJoinModel, error) {
	return uc.repo.GetLoanApplicationById(id)
}

func (uc *loanApplicationUsecase) LoanRepayment(id int, repayment *model.LoanRepaymentModel) error {
	loan, err := uc.repo.GetLoanApplicationById(id)
	if err != nil {
		return fmt.Errorf("failed to get loan application: %w", err)
	}

	if repayment.Payment < loan.Amount {
		log.Printf("Payment amount: %v, Loan amount: %v", repayment.Payment, loan.Amount)
		return fmt.Errorf("payment amount is less than the loan amount")
	}

	if repayment.Payment == loan.Amount {
		repayment.RepaymentStatus = model.RepaymentStatusLunas
	}

	if repayment.PaymentDate.Before(loan.DueDate) {
		return fmt.Errorf("payment date must be on or after due date")
	}

	err = uc.repo.LoanRepayment(id, repayment)
	if err != nil {
		return fmt.Errorf("failed to update loan repayment: %w", err)
	}

	return nil
}

func (uc *loanApplicationUsecase) GetLoanApplicationRepaymentStatus(page, limit int, repaymentStatus model.StatusEnum) ([]*model.LoanApplicationJoinModel, error) {
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	return uc.repo.GetLoanApplicationRepaymentStatus(page, limit, repaymentStatus)
}

func (uc *loanApplicationUsecase) GenerateIncomeReport(startDate time.Time, endDate time.Time) ([]*model.LoanRepaymentModel, float64, error) {
	loanRepayments, err := uc.repo.GetLoanRepaymentsByDateRange(startDate, endDate)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to retrieve loan repayments: %w", err)
	}

	totalIncome := 0.0
	for _, repayment := range loanRepayments {
		totalIncome += repayment.Payment
	}

	return loanRepayments, totalIncome, nil
}

func NewLoanApplicationUseCase(repo repository.LoanApplicationRepo) LoanApplicationUsecase {
	return &loanApplicationUsecase{
		repo: repo,
	}
}
