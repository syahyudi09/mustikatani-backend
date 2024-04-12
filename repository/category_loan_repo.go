package repository

import (
	"database/sql"
	"fmt"
	"pinjam-modal-app/model"
	"pinjam-modal-app/utils"
)

type CategoryLoanRepo interface {
	GetCategoryLoanById(int) (*model.CategoryLoanModel, error)
	GetCategoryLoanByName(string) (*model.CategoryLoanModel, error)
	GetAllCategoryLoan() ([]*model.CategoryLoanModel, error)
	InsertCategoryLoan(*model.CategoryLoanModel) error
	UpdateCategoryLoan(int, *model.CategoryLoanModel) error
	DeleteCategoryLoan(*model.CategoryLoanModel) error
}

type CategoryLoanRepoImpl struct {
	db *sql.DB
}

func (repo *CategoryLoanRepoImpl) GetCategoryLoanById(id int) (*model.CategoryLoanModel, error) {
	qry := utils.GET_CATEGORY_LOAN_BY_ID

	ctr := &model.CategoryLoanModel{}
	err := repo.db.QueryRow(qry, id).Scan(&ctr.Id, &ctr.CategoryLoanName, &ctr.CreatedAt, &ctr.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on CategoryLoanRepoImpl.GetCategoryLoanById(): %w", err)
	}
	return ctr, nil
}

func (repo *CategoryLoanRepoImpl) GetCategoryLoanByName(name string) (*model.CategoryLoanModel, error) {
	qry := utils.GET_CATEGORY_LOAN_BY_NAME

	ctr := &model.CategoryLoanModel{}
	err := repo.db.QueryRow(qry, name).Scan(&ctr.Id, &ctr.CategoryLoanName, &ctr.CreatedAt, &ctr.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on CategoryLoanRepoImpl.GetCategoryLoanByName(): %w", err)
	}
	return ctr, nil
}

func (repo *CategoryLoanRepoImpl) GetAllCategoryLoan() ([]*model.CategoryLoanModel, error) {
	qry := utils.GET_ALLCATEGORYLOAN

	rows, err := repo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("error on CategoryLoanRepoImpl.GetAllCategoryLoan(): %w", err)
	}
	defer rows.Close()

	var categoryLoans []*model.CategoryLoanModel
	for rows.Next() {
		ctr := &model.CategoryLoanModel{}
		err := rows.Scan(&ctr.Id, &ctr.CategoryLoanName, &ctr.CreatedAt, &ctr.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("error on CategoryLoanRepoImpl.GetAllCategoryLoan(): %w", err)
		}
		categoryLoans = append(categoryLoans, ctr)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error on CategoryLoanRepoImpl.GetAllCategoryLoan(): %w", err)
	}

	return categoryLoans, nil
}

func (repo *CategoryLoanRepoImpl) InsertCategoryLoan(ctr *model.CategoryLoanModel) error {
	qry := utils.INSERT_CATEGORY_LOAN

	err := repo.db.QueryRow(qry, ctr.CategoryLoanName, ctr.CreatedAt, ctr.UpdatedAt).Scan(&ctr.Id)
	if err != nil {
		return fmt.Errorf("error on CategoryLoanRepoImpl.InsertCategoryLoan(): %w", err)
	}
	return nil
}

func (repo *CategoryLoanRepoImpl) UpdateCategoryLoan(id int, ctr *model.CategoryLoanModel) error {
	qryid := utils.GET_CATEGORY_UPDATE_ID
	err := repo.db.QueryRow(qryid, ctr.Id).Scan(&ctr.Id)
	if err != nil {
		return fmt.Errorf("error on CategoryLoanRepoImpl.UpdateCategoryLoan(): %w", err)
	}

	qry := utils.UPDATE_CATEGORY_LOAN

	_, err = repo.db.Exec(qry, ctr.CategoryLoanName, ctr.UpdatedAt, &ctr.CreatedAt, ctr.Id)
	if err != nil {
		return fmt.Errorf("error on CategoryLoanRepoImpl.UpdateCategoryLoan(): %w", err)
	}
	return nil
}

func (repo *CategoryLoanRepoImpl) DeleteCategoryLoan(categoryLoan *model.CategoryLoanModel) error {
	qry := utils.DELETE_CATEGORYLOAN

	_, err := repo.db.Exec(qry, categoryLoan.Id)
	if err != nil {
		return fmt.Errorf("error on CategoryLoanRepoImpl.DeleteCategoryLoan(): %w", err)
	}
	return nil
}

func NewCategoryLoanRepo(db *sql.DB) CategoryLoanRepo {
	return &CategoryLoanRepoImpl{
		db: db,
	}
}
