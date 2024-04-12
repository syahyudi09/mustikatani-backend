package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"pinjam-modal-app/apperror"
	"pinjam-modal-app/model"
	"pinjam-modal-app/usecase"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CategoryLoanHandler struct {
	router  *gin.Engine
	usecase usecase.CategoryLoanUsecase
}

func (ch *CategoryLoanHandler) InsertCategoryLoan(ctx *gin.Context) {
	var req model.InsertCategoryLoan
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	categoryLoan := model.CategoryLoanModel{
		CategoryLoanName: req.CategoryLoanName,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	err := ch.usecase.InsertCategoryLoan(&categoryLoan)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("CategoryLoanHandler.InsertCategoryLoan() 1: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("CategoryLoanHandler.InsertCategoryLoan() 2: %v", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Cannot insert category loan due to an error",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (ch *CategoryLoanHandler) GetAllCategoryLoan(ctx *gin.Context) {
	categoryLoans, err := ch.usecase.GetAllCategoryLoan()
	if err != nil {
		errResponse := apperror.NewAppError(http.StatusInternalServerError, "Failed to retrieve category loan data")
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	successResponse := gin.H{
		"success": true,
		"data":    categoryLoans,
	}
	ctx.JSON(http.StatusOK, successResponse)
}
func (ch *CategoryLoanHandler) GetCategoryLoanById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		err := apperror.NewAppError(http.StatusBadRequest, "ID cannot be empty")
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		err := apperror.NewAppError(http.StatusBadRequest, "ID must be a number")
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	CategoryLoans, err := ch.usecase.GetCategoryLoanById(id)
	if err != nil {
		log.Printf("CategoryLoanHandler.getCategoryLoanById(): %v", err.Error())
		err := apperror.NewAppError(http.StatusInternalServerError, "Failed to retrieve CategoryLoan data")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	successResponse := gin.H{
		"success": true,
		"data":    CategoryLoans,
	}
	ctx.JSON(http.StatusOK, successResponse)
}

func (ch *CategoryLoanHandler) GetCategoryLoanByName(ctx *gin.Context) {
	name := ctx.Param("name")
	if name == "" {
		err := apperror.NewAppError(http.StatusBadRequest, "Name cannot be empty")
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	categoryLoans, err := ch.usecase.GetCategoryLoanByName(name)
	if err != nil {
		log.Printf("CategoryLoanHandler.GetCategoryLoanByName(): %v", err.Error())
		err := apperror.NewAppError(http.StatusInternalServerError, "Failed to retrieve CategoryLoan data")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	successResponse := gin.H{
		"success": true,
		"data":    categoryLoans,
	}
	ctx.JSON(http.StatusOK, successResponse)
}

func (ch *CategoryLoanHandler) UpdateCategoryLoan(ctx *gin.Context) {
	cl := &model.CategoryLoanModel{}
	if err := ctx.ShouldBindJSON(&cl); err != nil {
		errResponse := apperror.NewAppError(http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	// Cek apakah kategori pinjaman dengan ID yang diberikan ada dalam database
	existingCategoryLoan, err := ch.usecase.GetCategoryLoanById(cl.Id)
	if err != nil {
		log.Printf("CategoryLoanHandler.UpdateCategoryLoan(): %v", err.Error())
		errResponse := apperror.NewAppError(http.StatusInternalServerError, "Failed to update category loan")
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	if existingCategoryLoan == nil {
		errResponse := apperror.NewAppError(http.StatusBadRequest, "Category loan not found")
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	err = ch.usecase.UpdateCategoryLoan(cl.Id, cl)
	if err != nil {
		log.Printf("CategoryLoanHandler.UpdateCategoryLoan(): %v", err.Error())
		errResponse := apperror.NewAppError(http.StatusInternalServerError, "Failed to update category loan")
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	successResponse := gin.H{
		"success": true,
	}
	ctx.JSON(http.StatusOK, successResponse)
}

func (ch *CategoryLoanHandler) DeleteCategoryLoan(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		err := apperror.NewAppError(http.StatusBadRequest, "ID cannot be empty")
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		err := apperror.NewAppError(http.StatusBadRequest, "ID must be a number")
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	categoryLoan, err := ch.usecase.GetCategoryLoanById(id)
	if err != nil {
		log.Printf("CategoryLoanHandler.DeleteCategoryLoan(): %v", err.Error())
		err := apperror.NewAppError(http.StatusInternalServerError, "Failed to delete CategoryLoan")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if categoryLoan == nil {
		err := apperror.NewAppError(http.StatusNotFound, "CategoryLoan not found")
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	err = ch.usecase.DeleteCategoryLoan(categoryLoan)
	if err != nil {
		log.Printf("CategoryLoanHandler.DeleteCategoryLoan(): %v", err.Error())
		err := apperror.NewAppError(http.StatusInternalServerError, "Failed to delete CategoryLoan")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	successResponse := gin.H{
		"success": true,
	}
	ctx.JSON(http.StatusOK, successResponse)
}

func NewCategoryLoanHandler(r *gin.Engine, usecase usecase.CategoryLoanUsecase) *CategoryLoanHandler {
	handler := CategoryLoanHandler{
		router:  r,
		usecase: usecase,
	}

	r.POST("/category-loan", handler.InsertCategoryLoan)
	r.GET("/category-loan", handler.GetAllCategoryLoan)
	r.GET("/category-loan/:name", handler.GetCategoryLoanByName)
	r.GET("/category-loan/id/:id", handler.GetCategoryLoanById)
	r.PUT("/category-loan", handler.UpdateCategoryLoan)
	r.DELETE("/category-loan/:id", handler.DeleteCategoryLoan)
	return &handler
}
