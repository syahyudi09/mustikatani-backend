package usecase

import (
	"fmt"
	"log"
	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"
	"time"
)

type GoodsUsecase interface {
	InsertGoods(*model.GoodsModel) error
	GetAllTrxGoods(page, limit int) ([]*model.LoanGoodsModel, error)
	GetGoodsById(int) (*model.LoanGoodsModel, error)
	UpdateGoodsRepayment(int, *model.LoanRepaymentModel) error
	GetGooodsRepaymentStatus(page, limit int, repaymentStatus model.StatusEnum) ([]*model.LoanGoodsModel, error)
	GenerateIncomeReport(startDate time.Time, endDate time.Time) ([]*model.LoanRepaymentModel, float64, error)
}

type goodsUsecaseImpl struct {
	goodsRepo repository.GoodsRepo
}

func (goodsUsecase *goodsUsecaseImpl) InsertGoods(goods *model.GoodsModel) error {
	customerDB, err := goodsUsecase.goodsRepo.GetCustomerById(goods.CustomerId)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	if customerDB.NIK.Valid && customerDB.NIK.String != "" &&
		customerDB.NoKK.Valid && customerDB.NoKK.String != "" &&
		customerDB.EmergencyName.Valid && customerDB.EmergencyName.String != "" &&
		customerDB.EmergencyContact.Valid && customerDB.EmergencyContact.String != "" &&
		customerDB.LastSalary.Valid && customerDB.LastSalary.Float64 != 0 {
		goods.Status = "APPROVE"
		goods.RepaymentStatus = model.StatusEnum(model.RepaymentStatusBelumLunas)
		goods.DueDate = time.Now().AddDate(0, 2, 0)
	} else {
		goods.Status = "DENIED"
		fmt.Println("Silakan lengkapi data customer")
	}

	err = goodsUsecase.goodsRepo.InsertGoods(goods)
	if err != nil {
		return fmt.Errorf("failed to insert goods: %v", err)
	}

	return nil
}

func (goodsUsecase *goodsUsecaseImpl) GetAllTrxGoods(page, limit int) ([]*model.LoanGoodsModel, error){
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}
	return goodsUsecase.goodsRepo.GetAllTrxGoods(page, limit)
}

func (goodUsecase *goodsUsecaseImpl) UpdateGoodsRepayment(id int, repayment *model.LoanRepaymentModel) error{
	goods, err := goodUsecase.goodsRepo.GetGoodsById(id)
	if err != nil {
		return fmt.Errorf("failed to get loan application: %w", err)
	}

	if repayment.Payment < goods.Amount {
		log.Printf("Payment amount: %v, Loan amount: %v", repayment.Payment, goods.Amount)
		return fmt.Errorf("payment amount is less than the loan amount")
	}

	if repayment.Payment == goods.Amount {
		repayment.RepaymentStatus = model.RepaymentStatusLunas
	}

	if repayment.PaymentDate.Before(goods.DueDate) {
		return fmt.Errorf("payment date must be on or after due date")
	}

	err = goodUsecase.goodsRepo.UpdateGoodsRepayment(id, repayment)
	if err != nil {
		return fmt.Errorf("failed to update loan repayment: %w", err)
	}

	return nil
}

func (goodsUsecase *goodsUsecaseImpl)  GetGooodsRepaymentStatus(page, limit int, repaymentStatus model.StatusEnum) ([]*model.LoanGoodsModel, error){
	if page <= 0 {
		page = 1
	}
	if limit <= 0 {
		limit = 10
	}

	return goodsUsecase.goodsRepo.GetGooodsRepaymentStatus(page, limit, repaymentStatus)
}

func (goodsUsecase *goodsUsecaseImpl) GenerateIncomeReport(startDate time.Time, endDate time.Time) ([]*model.LoanRepaymentModel, float64, error) {
	loanRepayments, err := goodsUsecase.goodsRepo.GetLoanGoodsRepaymentsByDateRange(startDate, endDate)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to retrieve loan repayments: %w", err)
	}

	totalIncome := 0.0
	for _, repayment := range loanRepayments {
		totalIncome += repayment.Payment
	}

	return loanRepayments, totalIncome, nil
}

func (goodsUsecase *goodsUsecaseImpl) GetGoodsById(id int) (*model.LoanGoodsModel, error){
	return goodsUsecase.goodsRepo.GetGoodsById(id)
}


func NewGoodsUsecase(goodsRepo repository.GoodsRepo) GoodsUsecase {
	return &goodsUsecaseImpl{
		goodsRepo: goodsRepo,
	}
}
