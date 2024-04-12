package handler

import (
	"errors"
	"fmt"
	"net/http"
	"pinjam-modal-app/apperror"
	"pinjam-modal-app/model"
	"pinjam-modal-app/usecase"
	"pinjam-modal-app/middleware"
	"strconv"
	"github.com/gin-gonic/gin"
)

type CustomerHandler interface {
}

type customerHandlerImpl struct {
	router  *gin.Engine
	cstUsecase usecase.CustomerUsecase
}

func (cstHandler customerHandlerImpl) AddCustomer(ctx *gin.Context) {
	cst := &model.CustomerModel{}
	err := ctx.ShouldBindJSON(&cst)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Invalid JSON data",
		})
		return
	}

	if len(cst.FullName) > 15 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Name is to long max at 15 character",
		})
		return
	}

	err = cstHandler.cstUsecase.AddCustomer(cst)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("CustomerHandler.AddCustomer() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("CustomerHandler.AddCustomer() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Terjadi kesalahan ketika menyimpan data customer",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})

}

func (cstHandler customerHandlerImpl) GetCustomerById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "Id harus angka",
		})
		return
	}

	cst, err := cstHandler.cstUsecase.GetCustomerById(id)
	if err != nil {
		apperror := &apperror.AppError{}
		if errors.As(err, apperror) {
			fmt.Printf("customerHandlerImpl.GetCustomerById():%v", err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success":      false,
				"errorMassage": apperror.Error(),
			})
			return
		}
		fmt.Printf("customerHandlerImpl.GetCustomerById(() : %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "Terjadi kesalahan ketika mengambil data customer",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    cst,
	})
}

func (cstHandler customerHandlerImpl) GetAllCustomer(ctx *gin.Context) {

	cst, err := cstHandler.cstUsecase.GetAllCustomer()
	if err != nil {
		fmt.Printf("serviceHandlerImpl.GetAllCustomer(() : %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "Terjadi kesalahan ketika mengambil data service",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    cst,
	})
}

func (cstHandler customerHandlerImpl) UpdateCustomer(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "Id harus angka",
		})
		return
	}

	cst := &model.CustomerModel{}
	err = ctx.ShouldBind(&cst)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "JSON salah",
		})
		return
	}
	cst.Id = id
	err = cstHandler.cstUsecase.UpdateCustomer(cst)
	if err != nil {
		fmt.Printf("customerHandlerImpl.UpdateCustomer(() : %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "Terjadi kesalahan ketika mengubah data service",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    cst,
	})
}

func (cstHandler customerHandlerImpl) DeleteCustomer(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "Id harus angka",
		})
		return
	}

	err = cstHandler.cstUsecase.DeleteCustomer(id)
	if err != nil {
		fmt.Printf("customerHandlerImpl.DeleteCustomer(() : %v", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMassage": "Terjadi kesalahan ketika menghapus data customer",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"Massage": "Customer telah di hapus",
		"ID":      id,
	})
}

func NewCustomerHandler(router  *gin.Engine, cstUsecase usecase.CustomerUsecase) CustomerHandler {
	cstHandler := &customerHandlerImpl{
		router: router,
		cstUsecase: cstUsecase,
	}
	router.POST("/Customer", cstHandler.AddCustomer)
	router.GET("/Customer",middleware.RequireToken(),middleware.AdminOnly(), cstHandler.GetAllCustomer)
	router.GET("/Customer/:id",middleware.RequireToken(),middleware.AdminOnly(), cstHandler.GetCustomerById)
	router.PUT("/Customer/:id",middleware.RequireToken(),middleware.AdminOnly(), cstHandler.UpdateCustomer)
	router.DELETE("/Customer/:id",middleware.RequireToken(), middleware.AdminOnly(), cstHandler.DeleteCustomer)

	return cstHandler
}
