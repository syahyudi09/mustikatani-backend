package repository

import (
	"database/sql"
	"fmt"
	"pinjam-modal-app/model"
	"pinjam-modal-app/utils"
	"time"
)

type CustomerRepo interface {
	AddCustomer(*model.CustomerModel) error
	GetAllCustomer() ([]model.CustomerModel, error)
	GetCustomerById(int) (*model.CustomerModel, error)
	UpdateCustomer(*model.CustomerModel) error
	DeleteCustomer(int) error
	GetCustomerByNIK(string) (int, error)
	GetCustomerByNumber(string) (int, error)
}

type customerRepoImpl struct {
	db *sql.DB
}

func (cstRepo *customerRepoImpl) AddCustomer(cst *model.CustomerModel) error {

	qry := utils.ADD_CUSTOMER

	err := cstRepo.db.QueryRow(qry, cst.FullName, cst.Address, cst.NIK, cst.Phone, cst.UserId, cst.NoKK, cst.EmergencyContact, cst.EmergencyName, cst.LastSalary, time.Now()).Scan(&cst.Id)
	if err != nil {
		return fmt.Errorf("error on customerRepoImpl.AddCustomer() : %w", err)
	}
	return nil
}

func (cstRepo *customerRepoImpl) GetCustomerById(id int) (*model.CustomerModel, error) {
	qry := utils.GET_CUSTOMER_BY_ID

	cst := &model.CustomerModel{}
	err := cstRepo.db.QueryRow(qry, id).Scan(&cst.Id, &cst.FullName, &cst.Address, &cst.NIK, &cst.Phone, &cst.UserId, &cst.CreatedAt, &cst.UpdatedAt, &cst.NoKK, &cst.EmergencyName, &cst.EmergencyContact, &cst.LastSalary)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on customerRepoImpl.getCustomerById() : %w", err)
	}
	return cst, nil
}

func (cstRepo *customerRepoImpl) UpdateCustomer(cst *model.CustomerModel) error {
	qry := utils.UPDATE_CUSTOMER

	result, err := cstRepo.db.Exec(qry, cst.FullName, cst.Address, cst.Phone, time.Now(), cst.EmergencyName, cst.EmergencyContact, cst.LastSalary, cst.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return fmt.Errorf("error on customerRepoImpl.UpdateCustomer() : %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("updateService(): failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ID tidak di temukan, ID : %d", cst.Id)
	}
	return nil
}

func (cusRepo *customerRepoImpl) GetAllCustomer() ([]model.CustomerModel, error) {
	qry := utils.GET_ALL_CUSTOMER

	cst := &model.CustomerModel{}
	rows, err := cusRepo.db.Query(qry)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("error on customerRepoImpl.GetAllCustomer() : %w", err)
	}

	arrCus := []model.CustomerModel{}
	for rows.Next() {
		rows.Scan(&cst.Id, &cst.FullName, &cst.Address, &cst.NIK, &cst.Phone, &cst.UserId, &cst.CreatedAt, &cst.UpdatedAt, &cst.NoKK, &cst.EmergencyContact, &cst.EmergencyName, &cst.LastSalary)
		arrCus = append(arrCus, *cst)
	}
	return arrCus, nil
}

func (cstRepo *customerRepoImpl) DeleteCustomer(id int) error {

	qry := utils.DELETE_CUSTOMER
	result, err := cstRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("deleteCustomer() : %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("deleteCustomer(): failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("ID tidak di temukan, ID : %d", id)
	}
	return nil
}

func (cstRepo *customerRepoImpl) GetCustomerByNIK(nik string) (int, error) {
	qry := utils.GET_CUSTOMER_BY_NIK

	var cstID int
	err := cstRepo.db.QueryRow(qry, nik).Scan(&cstID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, nil
	}
	return cstID, nil
}

func (cstRepo *customerRepoImpl) GetCustomerByNumber(Phone string) (int, error) {
	qry := utils.GET_CUSTOMER_BY_NUMBER

	var cstID int
	err := cstRepo.db.QueryRow(qry, Phone).Scan(&cstID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, nil
	}
	return cstID, nil
}

func NewCustomerRepo(db *sql.DB) CustomerRepo {
	return &customerRepoImpl{
		db: db,
	}
}
