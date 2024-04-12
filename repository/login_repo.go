package repository

import (
	"database/sql"
	"fmt"
	"pinjam-modal-app/model"
	"pinjam-modal-app/utils"
)

type LoginRepo interface {
	GetUserByEmail(email string) (*model.UserModel, error)
}

type loginRepoImpl struct {
	db *sql.DB
}

func (LoginRepo *loginRepoImpl) GetUserByEmail(email string) (*model.UserModel, error) {
	query := utils.GET_USER_BY_EMAIL

	user := &model.UserModel{}
	err := LoginRepo.db.QueryRow(query, email).Scan(
		&user.Id,
		&user.UserName,
		&user.Email,
		&user.Password,
		&user.RolesName,
		&user.IsActive,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error in GetUserByEmail(): %w", err)
	}
	return user, nil
}

func NewLoginRepo(db *sql.DB) LoginRepo {
	return &loginRepoImpl{
		db: db,
	}
}
