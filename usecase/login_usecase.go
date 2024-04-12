package usecase

import (
	"fmt"
	"pinjam-modal-app/repository"
	utils "pinjam-modal-app/utils/authutil"

	"github.com/gin-contrib/sessions"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type LoginUsecase interface {
	Login(email, password string) (string, error)
	Logout(ctx *gin.Context) error

}

type loginUsecase struct {
	loginRepo repository.LoginRepo
}

func (lu *loginUsecase) Login(email, password string) (string, error) {
	// Mengecek apakah pengguna dengan email tersebut ada di penyimpanan data
	user, err := lu.loginRepo.GetUserByEmail(email)
	if err != nil {
		return "", fmt.Errorf("failed to retrieve user: %v", err)
	}

	// Verifikasi password pengguna dengan menggunakan bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid email or password")
	}

	// Menghasilkan token JWT
	token, err := utils.GenerateToken(user)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return token, nil
}

func (lu *loginUsecase) Logout(ctx *gin.Context) error {
	session := sessions.Default(ctx)
	session.Clear()
	session.Save()
	return nil
}

func NewLoginUsecase(loginRepo repository.LoginRepo) LoginUsecase {
	return &loginUsecase{
		loginRepo: loginRepo,
	}
}
