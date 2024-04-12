

package handler

import (
	// "fmt"
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

type GoodsHandler interface {
}

type goodsHandlerImpl struct {
	router  *gin.Engine
	goodsUsecase usecase.GoodsUsecase
}

func (goodsHandler *goodsHandlerImpl) InsertGoods(ctx *gin.Context) {
	goods := &model.GoodsModel{}
	err := ctx.ShouldBindJSON(&goods)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid data JSON",
		})
		return
	}

	err = goodsHandler.goodsUsecase.InsertGoods(goods)
	if err != nil{
		fmt.Printf("error an cpHandler.cpUsecase.InsertCategoryProduct : %v ", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika menyimpan data category product",
		})
		return
	}
	if goods.Status == "APPROVE" {
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

func (goodsHandler *goodsHandlerImpl) GetGoodsById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == ""{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id harus angka",
		})
		return
	}

	goods, err := goodsHandler.goodsUsecase.GetGoodsById(id)
	if err != nil {
		fmt.Printf(" cpHandler.cpUsecase.GetCategoryProductById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Terjadi kesalahan ketika mengambil data category product",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    goods,
	})
}

func(goodsHandler *goodsHandlerImpl) GetAllTrxGoods(ctx *gin.Context){

		page, err := strconv.Atoi(ctx.Query("page"))
		if err != nil {
			page = 1
		}
	
		limit, err := strconv.Atoi(ctx.Query("limit"))
		if err != nil {
			limit = 10
		}
	
		loanGoods, err := goodsHandler.goodsUsecase.GetAllTrxGoods(page, limit)
		if err != nil {
			log.Println("Failed to create loan application:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Failed to retrieve loan applications",
			})
			return
		}
	
		response := make([]model.LoanGoodsModel, 0)
		for _, loanGood := range loanGoods {
			response = append(response, *loanGood)
		}
	
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    response,
		})

}
func (goodsHandler *goodsHandlerImpl) GoodsRepayment(ctx *gin.Context) {
		goodsID := ctx.Param("id")
		if goodsID == "" {
			errResponse := apperror.NewAppError(http.StatusBadRequest, "ID cannot be empty")
			ctx.JSON(http.StatusBadRequest, errResponse)
			return
		}
	
		id, err := strconv.Atoi(goodsID)
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
	
		repaymentGoods := &model.LoanRepaymentModel{
			PaymentDate: req.PaymentDate,
			Payment:     req.Payment,
			UpdatedAt:   time.Now(),
		}
	
		err = goodsHandler.goodsUsecase.UpdateGoodsRepayment(id, repaymentGoods)
		if err != nil {
			log.Println("Failed to process loan repayment:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": fmt.Sprintf("Failed to process loan repayment: %v", err),
			})
			return
		}
		goods := &model.GoodsModel{}
	
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"Amount": goods.Amount,
			"message": "Loan repayment processed successfully",
		})
}

func (goodsHandler *goodsHandlerImpl) GetLoanGoodsByRepaymentStatus(ctx *gin.Context) {
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

	lunasApplications, err := goodsHandler.goodsUsecase.GetGooodsRepaymentStatus(page, limit, repaymentStatus)
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

func (goodsHandler *goodsHandlerImpl) generateIncomeReport(ctx *gin.Context) {
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

	loanRepayments, totalIncome, err := goodsHandler.goodsUsecase.GenerateIncomeReport(startDate, endDate)
	if err != nil {
		log.Println("err :", err)
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

	
func NewGoodsHandler(srv *gin.Engine,goodsUsecase usecase.GoodsUsecase) GoodsHandler {
	ghandler := goodsHandlerImpl{
		router: srv,
		goodsUsecase: goodsUsecase,
	}

	srv.POST("/goods", ghandler.InsertGoods)
	srv.GET("/goods/:id", ghandler.GetGoodsById)
	srv.GET("/goods", ghandler.GetAllTrxGoods)
	srv.PUT("/goods-update-payment/:id", ghandler.GoodsRepayment)
	srv.GET("/goods-repayment", ghandler.GetLoanGoodsByRepaymentStatus)
	srv.GET("/goods-income-report", ghandler.generateIncomeReport)
	return ghandler
}
