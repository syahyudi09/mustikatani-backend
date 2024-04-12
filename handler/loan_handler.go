package handler

import (
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

type LoanHandler struct {
	router  *gin.Engine
	usecase usecase.LoanApplicationUsecase
}

func (lh *LoanHandler) createLoanApplication(ctx *gin.Context) {
	var req model.LoanJoinRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	loan := model.LoanApplicationModel{
		CustomerId:      req.CustomerId,
		LoanDate:        req.LoanDate,
		DueDate:         req.DueDate,
		CategoryLoanId:  req.CategoryLoanId,
		Amount:          req.Amount,
		Description:     req.Description,
		Status:          req.Status,
		RepaymentStatus: req.RepaymentStatus,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err = lh.usecase.CreateLoanApplication(&loan)
	if err != nil {
		log.Println("Failed to create loan application:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Failed to create loan application",
		})
		return
	}

	if loan.Status == model.LoanStatusApprove {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  model.LoanStatusApprove,
			"message": "Loan application created successfully",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status":  model.LoanStatusDenied,
			"message": "Failed to create loan application",
		})
	}
}

func (lh *LoanHandler) getLoanApplications(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	loanApplications, err := lh.usecase.GetLoanApplications(page, limit)
	if err != nil {
		log.Println("Failed to create loan application:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Failed to retrieve loan applications",
		})
		return
	}

	response := make([]model.LoanApplicationJoinModel, 0)
	for _, loanApplication := range loanApplications {
		response = append(response, *loanApplication)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

func (lh *LoanHandler) getLoanApplicationById(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid ID",
		})
		return
	}

	loanApplication, err := lh.usecase.GetLoanApplicationById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Failed to retrieve loan application",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    loanApplication,
	})
}

func (lh *LoanHandler) loanRepayment(ctx *gin.Context) {
	loanID := ctx.Param("id")
	if loanID == "" {
		errResponse := apperror.NewAppError(http.StatusBadRequest, "ID cannot be empty")
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	id, err := strconv.Atoi(loanID)
	if err != nil {
		errResponse := apperror.NewAppError(http.StatusBadRequest, "ID must be a number")
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	var req model.LoanRepaymentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errResponse := apperror.NewAppError(http.StatusBadRequest, "Invalid JSON data")
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	repayment := &model.LoanRepaymentModel{
		PaymentDate: req.PaymentDate,
		Payment:     req.Payment,
		UpdatedAt:   time.Now(),
	}

	err = lh.usecase.LoanRepayment(id, repayment)
	if err != nil {
		log.Println("Failed to process loan repayment:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": fmt.Sprintf("Failed to process loan repayment: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Loan repayment processed successfully",
	})
}

func (lh *LoanHandler) getLoanApplicationsByRepaymentStatus(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	limitStr := ctx.Query("limit")
	repaymentStatusStr := ctx.Query("repayment_status")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page value"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit value"})
		return
	}

	var repaymentStatus model.StatusEnum
	if repaymentStatusStr == "lunas" {
		repaymentStatus = model.RepaymentStatusLunas
	} else if repaymentStatusStr == "belum lunas" {
		repaymentStatus = model.RepaymentStatusBelumLunas
	} else {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid repayment status value"})
		return
	}

	lunasApplications, err := lh.usecase.GetLoanApplicationRepaymentStatus(page, limit, repaymentStatus)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": fmt.Sprintf("Failed to get loan applications: %v", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    lunasApplications,
	})
}

func (lh *LoanHandler) generateIncomeReport(ctx *gin.Context) {
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start date"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end date"})
		return
	}

	loanRepayments, totalIncome, err := lh.usecase.GenerateIncomeReport(startDate, endDate)
	if err != nil {
		errResponse := apperror.NewAppError(http.StatusInternalServerError, "Failed to generate income report")
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	successResponse := gin.H{
		"success":        true,
		"loanRepayments": loanRepayments,
		"totalIncome":    totalIncome,
	}
	ctx.JSON(http.StatusOK, successResponse)
}

func NewLoanApplicationHandler(r *gin.Engine, usecase usecase.LoanApplicationUsecase) *LoanHandler {
	handler := LoanHandler{
		router:  r,
		usecase: usecase,
	}
	r.POST("/loan", handler.createLoanApplication)
	r.GET("/loan", handler.getLoanApplications)
	r.GET("/loan/:id", handler.getLoanApplicationById)
	r.PUT("/loan/:id", handler.loanRepayment)
	r.GET("/loan-application", handler.getLoanApplicationsByRepaymentStatus)
	r.GET("/income-report", handler.generateIncomeReport)

	return &handler
}
