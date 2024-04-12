package usecase

import (
	"fmt"
	"pinjam-modal-app/repository"
)

type LogoutUsecase interface {
	LogoutByEmail(email string) error
}

type logoutUsecaseImpl struct {
	userRepository repository.UserRepo
}

func NewLogoutUsecase(userRepository repository.UserRepo) LogoutUsecase {
	return &logoutUsecaseImpl{
		userRepository: userRepository,
	}
}

func (uc *logoutUsecaseImpl) LogoutByEmail(email string) error {
	err := uc.userRepository.DeleteTokenByEmail(email)
	if err != nil {
		return fmt.Errorf("failed to logout user: %w", err)
	}
	return nil
}
