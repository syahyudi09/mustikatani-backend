package manager

import (
	"pinjam-modal-app/repository"
	"sync"
)

type RepoManager interface {
	GetUserRepo() repository.UserRepo
	GetLoginRepo() repository.LoginRepo
	GetCustomerRepo() repository.CustomerRepo
	GetProductRepo() repository.ProductRepo
	GetLoanApplicationRepo() repository.LoanApplicationRepo
	GetCategoryLoanRepo() repository.CategoryLoanRepo
	GetCategoryProductRepo() repository.CategoryProductRepo
	GetGoodsRepo() repository.GoodsRepo
}

type repoManager struct {
	infraManager     InfraManager
	userRepo         repository.UserRepo
	loginRepo        repository.LoginRepo
	cstRepo          repository.CustomerRepo
	productRepo      repository.ProductRepo
	loan             repository.LoanApplicationRepo
	categoryLoanRepo repository.CategoryLoanRepo
	CategoryProduct  repository.CategoryProductRepo
	TrxGoods         repository.GoodsRepo
}

var onceLoadUserRepo sync.Once
var onceLoadLoginRepo sync.Once
var onceLoadCustomerRepo sync.Once
var onceLoadCategoryLoan sync.Once
var onceLoadCategoryProductRepo sync.Once
var onceLoadGoodsRepo sync.Once
var onceLoadProductRepo sync.Once
var onceLoadLoanAppRepo sync.Once
var onceLoadRepo sync.Once

func (rm *repoManager) GetUserRepo() repository.UserRepo {
	onceLoadUserRepo.Do(func() {
		rm.userRepo = repository.NewUserRepo(rm.infraManager.GetDB())
	})
	return rm.userRepo
}
func (rm *repoManager) GetLoginRepo() repository.LoginRepo {
	onceLoadLoginRepo.Do(func() {
		rm.loginRepo = repository.NewLoginRepo(rm.infraManager.GetDB())
	})
	return rm.loginRepo

}

func (rm *repoManager) GetCustomerRepo() repository.CustomerRepo {
	onceLoadCustomerRepo.Do(func() {
		rm.cstRepo = repository.NewCustomerRepo(rm.infraManager.GetDB())
	})
	return rm.cstRepo
}

func (rm *repoManager) GetCategoryLoanRepo() repository.CategoryLoanRepo {
	onceLoadCategoryLoan.Do(func() {
		rm.categoryLoanRepo = repository.NewCategoryLoanRepo(rm.infraManager.GetDB())
	})
	return rm.categoryLoanRepo
}

func (rm *repoManager) GetProductRepo() repository.ProductRepo {
	onceLoadProductRepo.Do(func() {
		rm.productRepo = repository.NewProductRepo(rm.infraManager.GetDB())
	})
	return rm.productRepo
}

func (rm *repoManager) GetLoanApplicationRepo() repository.LoanApplicationRepo {
	onceLoadLoanAppRepo.Do(func() {
		rm.loan = repository.NewLoanApplicationRepository(rm.infraManager.GetDB())
	})
	return rm.loan
}
func (rm *repoManager) GetCategoryProductRepo() repository.CategoryProductRepo {
	onceLoadCategoryProductRepo.Do(func() {
		rm.CategoryProduct = repository.NewCategoryProductRepo(rm.infraManager.GetDB())
	})
	return rm.CategoryProduct
}

func (rm *repoManager) GetGoodsRepo() repository.GoodsRepo {
	onceLoadGoodsRepo.Do(func() {
		rm.TrxGoods = repository.NewGoodsRepo(rm.infraManager.GetDB())
	})
	return rm.TrxGoods
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
