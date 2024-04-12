package manager

import (
	"pinjam-modal-app/usecase"
	"sync"
)

type UsecaseManager interface {
	GetUserUsecase() usecase.UserUsecase
	GetLoginUsecase() usecase.LoginUsecase
	GetCustomerUsecase() usecase.CustomerUsecase
	GetProductUsecase() usecase.ProductUsecase
	GetCategoryLoanUsecase() usecase.CategoryLoanUsecase
	GetCategoryProductUsecase() usecase.CategoryProductUsecase
	GetGoodsUsecase() usecase.GoodsUsecase
	GetLoanAppUsecase() usecase.LoanApplicationUsecase
}

type usecaseManager struct {
	repoManager            RepoManager
	userUsecase            usecase.UserUsecase
	loginUsecase           usecase.LoginUsecase
	cstUsecase             usecase.CustomerUsecase
	categoryLoanUsecase    usecase.CategoryLoanUsecase
	categoryProductUsecase usecase.CategoryProductUsecase
	loanApp                usecase.LoanApplicationUsecase
	productUsecase         usecase.ProductUsecase
	trxGoodsUsecase        usecase.GoodsUsecase
}

var onceLoadUserUsecase sync.Once
var onceLoadLoginUsecase sync.Once
var onceLoadUsecase sync.Once
var onceLoadCustomerUsecase sync.Once
var onceLoadCategoryLoanUsecase sync.Once
var onceLoadGetCategoryProductUsecase sync.Once
var onceLoadGetGoodsUsecase sync.Once
var onceLoadProductUsecase sync.Once
var onceLoadLoanAppUsecase sync.Once

func (um *usecaseManager) GetUserUsecase() usecase.UserUsecase {
	onceLoadUserUsecase.Do(func() {
		um.userUsecase = usecase.NewUserUseCase(um.repoManager.GetUserRepo())
	})
	return um.userUsecase
}

func (um *usecaseManager) GetLoginUsecase() usecase.LoginUsecase {
	onceLoadLoginUsecase.Do(func() {
		um.loginUsecase = usecase.NewLoginUsecase(um.repoManager.GetLoginRepo())
	})
	return um.loginUsecase
}
func (um *usecaseManager) GetCustomerUsecase() usecase.CustomerUsecase {
	onceLoadCustomerUsecase.Do(func() {
		um.cstUsecase = usecase.NewCustomerUseCase(um.repoManager.GetCustomerRepo())
	})

	return um.cstUsecase
}

func (um *usecaseManager) GetCategoryLoanUsecase() usecase.CategoryLoanUsecase {
	onceLoadCategoryLoanUsecase.Do(func() {
		um.categoryLoanUsecase = usecase.NewCategoryLoanUsecase(um.repoManager.GetCategoryLoanRepo())
	})
	return um.categoryLoanUsecase
}

func (um *usecaseManager) GetProductUsecase() usecase.ProductUsecase {
	onceLoadProductUsecase.Do(func() {
		um.productUsecase = usecase.NewProductUseCase(um.repoManager.GetProductRepo())
	})
	return um.productUsecase
}

func (um *usecaseManager) GetLoanAppUsecase() usecase.LoanApplicationUsecase {
	onceLoadLoanAppUsecase.Do(func() {
		um.loanApp = usecase.NewLoanApplicationUseCase(um.repoManager.GetLoanApplicationRepo())
	})
	return um.loanApp
}
func (um *usecaseManager) GetCategoryProductUsecase() usecase.CategoryProductUsecase {
	onceLoadGetCategoryProductUsecase.Do(func() {
		um.categoryProductUsecase = usecase.NewCategoryProductUsecase(um.repoManager.GetCategoryProductRepo())
	})
	return um.categoryProductUsecase
}

func (um *usecaseManager) GetGoodsUsecase() usecase.GoodsUsecase {
	onceLoadGetGoodsUsecase.Do(func() {
		um.trxGoodsUsecase = usecase.NewGoodsUsecase(um.repoManager.GetGoodsRepo())
	})
	return um.trxGoodsUsecase

}

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
