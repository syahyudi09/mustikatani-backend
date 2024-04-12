package handler

import (
	"log"
	"net/http"
	"pinjam-modal-app/apperror"
	"pinjam-modal-app/model"
	"pinjam-modal-app/usecase"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	router  *gin.Engine
	usecase usecase.ProductUsecase
}

func (ph *ProductHandler) createProduct(ctx *gin.Context) {
	var req model.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errResponse := apperror.NewAppError(http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	product := model.ProductModel{
		ProductName:       req.ProductName,
		Description:       req.Description,
		Price:             req.Price,
		Stok:              req.Stok,
		CategoryProductId: req.CategoryProductId,
		Status:            req.Status,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	err := ph.usecase.CreateProduct(&product)
	if err != nil {
		log.Println("err : ", err)
		errResponse := apperror.NewAppError(http.StatusInternalServerError, "Cannot insert product due to an error")
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	successResponse := gin.H{
		"success": true,
	}
	ctx.JSON(http.StatusOK, successResponse)
}

func (ph *ProductHandler) getAllProduct(ctx *gin.Context) {
	pageStr := ctx.Query("page")
	limitStr := ctx.Query("limit")

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
	products, err := ph.usecase.GetAllProduct(page, limit)
	if err != nil {
		log.Println("err :",err)
		errResponse := apperror.NewAppError(http.StatusInternalServerError, "Failed to retrieve product data")
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	successResponse := gin.H{
		"success": true,
		"data":    products,
	}
	ctx.JSON(http.StatusOK, successResponse)
}

func (ph *ProductHandler) getProductById(ctx *gin.Context) {
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

	product, err := ph.usecase.GetProductById(id)
	if err != nil {
		log.Printf("ProductHandler.getProductById(): %v", err.Error())
		err := apperror.NewAppError(http.StatusInternalServerError, "Failed to retrieve product data")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	successResponse := gin.H{
		"success": true,
		"data":    product,
	}
	ctx.JSON(http.StatusOK, successResponse)
}

func (ph *ProductHandler) updateProduct(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		errResponse := apperror.NewAppError(http.StatusBadRequest, "ID cannot be empty")
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		errResponse := apperror.NewAppError(http.StatusBadRequest, "ID must be a number")
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	var req model.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		errResponse := apperror.NewAppError(http.StatusBadRequest, err.Error())
		ctx.JSON(http.StatusBadRequest, errResponse)
		return
	}

	product := model.ProductModel{
		ProductName:       req.ProductName,
		Description:       req.Description,
		Price:             req.Price,
		Stok:              req.Stok,
		CategoryProductId: req.CategoryProductId,
		Status:            req.Status,
		UpdatedAt:         time.Now(),
	}

	err = ph.usecase.UpdateProduct(id, &product)
	if err != nil {
		errResponse := apperror.NewAppError(http.StatusInternalServerError, "Failed to update product")
		ctx.JSON(http.StatusInternalServerError, errResponse)
		return
	}

	successResponse := gin.H{
		"success": true,
	}
	ctx.JSON(http.StatusOK, successResponse)
}

func (ph *ProductHandler) deleteProduct(ctx *gin.Context) {
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

	err = ph.usecase.DeleteProduct(id)
	if err != nil {
		log.Printf("ProductHandler.deleteProduct(): %v", err.Error())
		err := apperror.NewAppError(http.StatusInternalServerError, "Failed to delete product")
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	successResponse := gin.H{
		"success": true,
	}
	ctx.JSON(http.StatusOK, successResponse)
}

func NewProductHandler(r *gin.Engine, usecase usecase.ProductUsecase) *ProductHandler {
	handler := ProductHandler{
		router:  r,
		usecase: usecase,
	}
	r.POST("/product", handler.createProduct)
	r.GET("/product", handler.getAllProduct)
	r.GET("/product/:id", handler.getProductById)
	r.PUT("/product/:id", handler.updateProduct)
	r.DELETE("/product/:id", handler.deleteProduct)
	return &handler
}
