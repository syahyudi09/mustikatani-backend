package handler

import (
	"pinjam-modal-app/config"
	"pinjam-modal-app/manager"
	"pinjam-modal-app/middleware"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager
	engine         *gin.Engine
}

func (s *server) Run() {

	NewUserHandler(s.engine, s.usecaseManager.GetUserUsecase())
	NewLoginHandler(s.engine, s.usecaseManager.GetLoginUsecase())
	NewCustomerHandler(s.engine, s.usecaseManager.GetCustomerUsecase())
	NewProductHandler(s.engine, s.usecaseManager.GetProductUsecase())
	NewCategoryProductHandler(s.engine, s.usecaseManager.GetCategoryProductUsecase())
	NewGoodsHandler(s.engine, s.usecaseManager.GetGoodsUsecase())
	NewCategoryLoanHandler(s.engine, s.usecaseManager.GetCategoryLoanUsecase())
	NewLoanApplicationHandler(s.engine, s.usecaseManager.GetLoanAppUsecase())

	s.engine.Run(":8080")
	s.engine.Use(middleware.LoggerMiddleware())
	s.engine.Use(middleware.RequireToken())
	s.engine.Use(middleware.AdminOnly())

}

func NewServer() Server {
	c, err := config.NewConfig()
	if err != nil {
		panic(err)
	}

	infra := manager.NewInfraManager(c)
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	engine := gin.Default()

	return &server{
		usecaseManager: usecase,
		engine:         engine,
	}
}
