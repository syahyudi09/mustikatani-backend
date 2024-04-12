package usecase

import (
	"fmt"
	"pinjam-modal-app/apperror"
	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"
)

type CategoryProductUsecase interface {
	InsertCategoryProduct(*model.CategoryProductModel) error
	GetCategoryProductById(int) (*model.CategoryProductModel, error)
	GetAllCategoryProduct() ([]model.CategoryProductModel, error)
	UpdateCategoryProduct(int, *model.CategoryProductModel) error
	DeleteCategoryProduct(int) error
}

type categoryProductUsecaseImpl struct {
	cpRepo repository.CategoryProductRepo
}

func (cpUsecase *categoryProductUsecaseImpl) InsertCategoryProduct(cp *model.CategoryProductModel) error{
	cpDB, err := cpUsecase.cpRepo.GetCategoryProductByName(cp.CategoryProductName)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.InsertService() : %w", err)
	}

	if cpDB != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("data category product dengan nama %v sudah ada", cp.CategoryProductName),
		}
	}
	return cpUsecase.cpRepo.InsertCategoryProduct(cp)
}

func (cpUsecase *categoryProductUsecaseImpl) GetCategoryProductById(id int)(*model.CategoryProductModel, error){
	return cpUsecase.cpRepo.GetCategoryProductById(id)
}

func (cpUsecase *categoryProductUsecaseImpl) GetAllCategoryProduct() ([]model.CategoryProductModel, error){
	return cpUsecase.cpRepo.GetAllCategoryProduct()
}

func (cpUsecase *categoryProductUsecaseImpl) UpdateCategoryProduct(id int,cp *model.CategoryProductModel) error{
	cpDB, err := cpUsecase.cpRepo.GetCategoryProductByName(cp.CategoryProductName)
	if err != nil {
		return fmt.Errorf("serviceUsecaseImpl.InsertService() : %w", err)
	}

	if cpDB != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMessage: fmt.Sprintf("data category product dengan nama %v sudah ada", cp.CategoryProductName),
		}
	}
	return cpUsecase.cpRepo.UpdateCategoryProduct(id, cp)
}

func (cpUsecase *categoryProductUsecaseImpl) DeleteCategoryProduct(id int) error {
	return cpUsecase.cpRepo.DeleteCategoryProduct(id)
}

func NewCategoryProductUsecase(cpRepo repository.CategoryProductRepo) CategoryProductUsecase {
	return &categoryProductUsecaseImpl{
		cpRepo: cpRepo,
	}
}