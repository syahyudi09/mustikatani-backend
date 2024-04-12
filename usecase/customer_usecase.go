package usecase

import (
	"fmt"
	"pinjam-modal-app/apperror"
	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"
)

type CustomerUsecase interface {
	AddCustomer(*model.CustomerModel) error
	GetAllCustomer() ([]model.CustomerModel, error)
	GetCustomerById(int) (*model.CustomerModel, error)
	UpdateCustomer(*model.CustomerModel) error
	DeleteCustomer(int) error
}

type customerUsecaseImpl struct {
	cstRepo repository.CustomerRepo
}

func (cstUsecase *customerUsecaseImpl) AddCustomer(cst *model.CustomerModel) error {
	cstDb, err := cstUsecase.cstRepo.GetCustomerByNIK(cst.NIK)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.InsertService() : %w", err)
	}

	if cstDb != 0 {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("NIK %v telah terdaftar!", cst.NIK),
		}
	}
	numDb, err := cstUsecase.cstRepo.GetCustomerByNumber(cst.Phone)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.InsertService() : %w", err)
	}

	if numDb != 0 {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("Nomor telpon %v telah terdaftar!", cst.Phone),
		}
	}

	return cstUsecase.cstRepo.AddCustomer(cst)
}

func (cstUsecase *customerUsecaseImpl) GetAllCustomer() ([]model.CustomerModel, error) {
	return cstUsecase.cstRepo.GetAllCustomer()

}

func (cstUsecase *customerUsecaseImpl) GetCustomerById(id int) (*model.CustomerModel, error) {
	cstDb, err := cstUsecase.cstRepo.GetCustomerById(id)
	if err != nil {
		return cstDb, fmt.Errorf("customerUsecaseImpl.AddCustomer() : %w", err)
	}

	if cstDb == nil {
		return cstDb, apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("data customer dengan id %v tidak ada", id),
		}
	}
	return cstDb, err
}

func (cstUsecase *customerUsecaseImpl) UpdateCustomer(cst *model.CustomerModel) error {
	return cstUsecase.cstRepo.UpdateCustomer(cst)

}

func (cstUsecase *customerUsecaseImpl) DeleteCustomer(id int) error {
	return cstUsecase.cstRepo.DeleteCustomer(id)
}

func NewCustomerUseCase(cstRepo repository.CustomerRepo) CustomerUsecase {
	return &customerUsecaseImpl{
		cstRepo: cstRepo,
	}
}
