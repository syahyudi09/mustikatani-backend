package usecase

import (
	"errors"
	"fmt"
	"time"

	"pinjam-modal-app/model"
	"pinjam-modal-app/repository"
	"pinjam-modal-app/utils"
	"pinjam-modal-app/utils/authutil"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	RegisterUser(user *model.UserModel) error
	Login(email, password string) (string, error)
	GetUserById(int) (*model.UserModel, error)
	GetAllUser() (*[]model.UserModel, error)
	DeleteUser(*model.UserModel) error
	UpdateUser(user *model.UserModel) error
}

type userUsecaseImpl struct {
	userRepo repository.UserRepo
}

func (uc *userUsecaseImpl) RegisterUser(user *model.UserModel) error {
	// Validasi kredensial
	existingUser, err := uc.userRepo.GetUserByUsername(user.UserName)
	if err == nil && existingUser != nil {
		return errors.New("Username is already taken")
	}

	// Validasi email unik
	existingUser, err = uc.userRepo.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		return errors.New("Email is already taken")
	}

	// Validasi password keamanan
	if !utils.IsValidPassword(user.Password) {
		return errors.New("Password is not strong enough")
	}

	// Validasi data pengguna
	if !utils.IsValidEmail(user.Email) {
		return errors.New("Invalid email address")
	}

	// Generate hash password
	passHash, err := GeneratePasswordHash(user.Password)
	if err != nil {
		return fmt.Errorf("Failed to generate password hash: %w", err)
	}
	user.Password = passHash

	// Set nilai default
	user.RolesName = "User"
	user.IsActive = true
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Buat pengguna baru
	err = uc.userRepo.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}
func (uc *userUsecaseImpl) UpdateUser(user *model.UserModel) error {
	// Validate email uniqueness
	existingUser, err := uc.userRepo.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil && existingUser.Id != user.Id {
		return errors.New("Email is already taken")
	}

	// Validate password strength
	if !utils.IsValidPassword(user.Password) {
		return errors.New("Password is not strong enough")
	}

	// Validate email address
	if !utils.IsValidEmail(user.Email) {
		return errors.New("Invalid email address")
	}

	// Generate password hash if a new password is provided
	if user.Password != "" {
		passHash, err := GeneratePasswordHash(user.Password)
		if err != nil {
			return fmt.Errorf("Failed to generate password hash: %w", err)
		}
		user.Password = passHash
	}

	// Update the user
	err = uc.userRepo.UpdateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func GeneratePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (uc *userUsecaseImpl) Login(email, password string) (string, error) {
	// Cari pengguna berdasarkan email
	user, err := uc.userRepo.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	// Verifikasi password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Invalid email or password")
	}

	// Generate token
	token, err := authutil.GenerateToken(user)
	if err != nil {
		return "", fmt.Errorf("Failed to generate token: %w", err)
	}

	return token, nil
}

func (usrUsecase *userUsecaseImpl) GetUserById(id int) (*model.UserModel, error) {
	return usrUsecase.userRepo.GetUserById(id)
}

func (usrUsecase *userUsecaseImpl) GetAllUser() (*[]model.UserModel, error) {
	return usrUsecase.userRepo.GetAllUser()
}

func (usrUsecase *userUsecaseImpl) DeleteUser(usr *model.UserModel) error {
	return usrUsecase.userRepo.DeleteUser(usr)
}

func NewUserUseCase(userRepo repository.UserRepo) UserUsecase {
	return &userUsecaseImpl{
		userRepo: userRepo,
	}
}
