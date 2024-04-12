package handler

import (
	"net/http"
	"pinjam-modal-app/usecase"

	"github.com/gin-gonic/gin"
)

type LogoutHandler struct {
	logoutUsecase usecase.LogoutUsecase
}

func (h *LogoutHandler) Logout(c *gin.Context) {
	email := c.GetString("email") // Ambil email pengguna dari konteks Gin, sesuaikan dengan implementasi Anda

	err := h.logoutUsecase.LogoutByEmail(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}
func NewLogoutHandler(router *gin.Engine, logoutUsecase usecase.LogoutUsecase) {
	logoutHandler := &LogoutHandler{
		logoutUsecase: logoutUsecase,
	}

	router.POST("/logout", logoutHandler.Logout)
}
