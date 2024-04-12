package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"pinjam-modal-app/model"
	"pinjam-modal-app/utils"
	"time"
)

type GoodsRepo interface {
	InsertGoods(*model.GoodsModel) error
	GetCustomerById(int) (*model.ValidasiCustomerModel, error)
	GetAllTrxGoods(page, limit int) ([]*model.LoanGoodsModel, error)
	GetGoodsById(int) (*model.LoanGoodsModel, error)
	UpdateGoodsRepayment(int, *model.LoanRepaymentModel) error
	GetGooodsRepaymentStatus(page, limit int, repaymentStatus model.StatusEnum) ([]*model.LoanGoodsModel, error)
	GetLoanGoodsRepaymentsByDateRange(startDate time.Time, endDate time.Time) ([]*model.LoanRepaymentModel, error)
}

type goodsRepoImpl struct {
	db *sql.DB
}

func (goodsRepo *goodsRepoImpl) InsertGoods(goods *model.GoodsModel) error {
	insertQuery := utils.INSERT_GOODS
	goods.CreateAt = time.Now()
	goods.LoadDate = time.Now()

	tx, err := goodsRepo.db.Begin()
	if err != nil {
		return fmt.Errorf("gagal memulai transaksi: %w", err)
	}

	product := &model.ProductModel{}
	err = tx.QueryRow("SELECT price, stok FROM mst_product WHERE id = $1", goods.ProductId).Scan(&product.Price, &product.Stok)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("gagal mendapatkan harga produk: %w", err)
	}

	if product.Stok < goods.Quantity {
		tx.Rollback()
		return errors.New("stok produk tidak mencukupi")
	}

	goods.Amount = float64(goods.Quantity) * product.Price
	_, err = tx.Exec(insertQuery, goods.CustomerId, goods.LoadDate, goods.PaymentDate, goods.DueDate, goods.CategoryIdLoan, goods.ProductId, goods.Quantity, product.Price, goods.Amount, goods.CreateAt, goods.Status, goods.RepaymentStatus)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("gagal memasukkan data goods: %w", err)
	}

	newStock := product.Stok - goods.Quantity
	_, err = tx.Exec("UPDATE mst_product SET stok = $1 WHERE id = $2", newStock, goods.ProductId)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("gagal memperbarui stok produk: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("gagal melakukan commit transaksi: %w", err)
	}

	return nil
}

func (goodsRepo *goodsRepoImpl) GetCustomerById(id int) (*model.ValidasiCustomerModel, error) {
	qry := utils.GET_CUSTOMER_LOAN_BY_ID

	customer := &model.ValidasiCustomerModel{}
	err := goodsRepo.db.QueryRow(qry, id).Scan(
		&customer.Id, &customer.NIK, &customer.NoKK, &customer.EmergencyName, &customer.EmergencyContact, &customer.LastSalary)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("customer not found")
		}
		return nil, fmt.Errorf("error on GetCustomerById: %w", err)
	}

	return customer, nil
}

func (goodsRepo *goodsRepoImpl) GetGoodsById(id int) (*model.LoanGoodsModel, error) {
	qry := utils.GET_GOODS_BY_ID	
	loangoods := &model.LoanGoodsModel{}
	err := goodsRepo.db.QueryRow(qry, id).Scan(
					&loangoods.Id, &loangoods.CustomerId,&loangoods.LoanDate,
					&loangoods.DueDate, &loangoods.CategoryLoanID, &loangoods.ProductId, 
					&loangoods.ProductName,&loangoods.Quantity,&loangoods.Price, 
					&loangoods.Amount, &loangoods.Description,&loangoods.Status, 
					&loangoods.RepaymentStatus, &loangoods.CreatedAt,&loangoods.UpdatedAt,
					&loangoods.FullName, &loangoods.Address,&loangoods.PhoneNumber,
					&loangoods.NIK, &loangoods.NoKK,&loangoods.EmergencyName,
					&loangoods.EmergencyContact, &loangoods.LastSalary,
				)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("goods not found")
		}
		return nil, fmt.Errorf("error on goodsRepoImpl.GetGoodsById: %w", err)
	}

	return loangoods, nil
}

func (goodsRepo *goodsRepoImpl) GetAllTrxGoods(page, limit int) ([]*model.LoanGoodsModel, error) {
	offset := (page - 1) * limit
	query := utils.GET_ALL_TRX_GOODS
	rows, err := goodsRepo.db.Query(query, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get loan goods: %w", err)
	}
	defer rows.Close()

	loangoods := []*model.LoanGoodsModel{}
	for rows.Next() {
		loangood := &model.LoanGoodsModel{}
		err := rows.Scan(&loangood.Id, &loangood.CustomerId, &loangood.LoanDate, &loangood.DueDate, &loangood.CategoryLoanID, &loangood.ProductId, &loangood.ProductName, &loangood.Quantity, &loangood.Price, &loangood.Amount, &loangood.Description, &loangood.Status, &loangood.RepaymentStatus, &loangood.CreatedAt, &loangood.UpdatedAt, &loangood.FullName, &loangood.Address, &loangood.PhoneNumber, &loangood.NIK, &loangood.NoKK, &loangood.EmergencyName, &loangood.EmergencyContact, &loangood.LastSalary)
		if err != nil {
			return nil, fmt.Errorf("failed to scan loan application: %w", err)
		}
		loangoods = append(loangoods, loangood)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get loan applications: %w", err)
	}
	return loangoods, nil
}

func (goodsRepo *goodsRepoImpl) UpdateGoodsRepayment(id int, repayment *model.LoanRepaymentModel) error {
	updateStatment := utils.UPDATE_GOODS_REPAYMENT
	_, err := goodsRepo.db.Exec(updateStatment, repayment.PaymentDate, repayment.Payment, model.StatusEnum(repayment.RepaymentStatus), repayment.UpdatedAt, id)

	if err != nil {
		return fmt.Errorf("error on loanApplicationRepo.LoanRepayment() : %w", err)
	}
	return nil
}

func (goodsRepo *goodsRepoImpl) GetGooodsRepaymentStatus(page, limit int, repaymentStatus model.StatusEnum) ([]*model.LoanGoodsModel, error) {
	offset := (page - 1) * limit
	query := utils.GET_GOODS_REPAYMENT_STATUS
	rows, err := goodsRepo.db.Query(query, offset, limit, repaymentStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to get loan applications: %w", err)
	}
	defer rows.Close()
	loangoods := []*model.LoanGoodsModel{}
	for rows.Next() {
		loangood := &model.LoanGoodsModel{}
		err := rows.Scan(&loangood.Id, &loangood.CustomerId, &loangood.LoanDate, &loangood.DueDate, &loangood.CategoryLoanID, &loangood.ProductId, &loangood.ProductName, &loangood.Quantity, &loangood.Price, &loangood.Amount, &loangood.Description, &loangood.Status, &loangood.RepaymentStatus, &loangood.CreatedAt, &loangood.UpdatedAt, &loangood.FullName, &loangood.Address, &loangood.PhoneNumber, &loangood.NIK, &loangood.NoKK, &loangood.EmergencyName, &loangood.EmergencyContact, &loangood.LastSalary)
		if err != nil {
			return nil, fmt.Errorf("failed to scan loan application: %w", err)
		}
		loangoods = append(loangoods, loangood)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to get loan applications: %w", err)
	}
	return loangoods, nil
}

func (goodsRepo *goodsRepoImpl) GetLoanGoodsRepaymentsByDateRange(startDate time.Time, endDate time.Time) ([]*model.LoanRepaymentModel, error) {
	selectStatement := `
	SELECT payment_date, payment, repayment_status
	FROM trx_goods
	WHERE payment_date >= $1 AND payment_date <= $2
`

	rows, err := goodsRepo.db.Query(selectStatement, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("error querying loan repayments: %w", err)
	}
	defer rows.Close()

	loanRepayments := []*model.LoanRepaymentModel{}
	for rows.Next() {
		loanRepayment := &model.LoanRepaymentModel{}
		err := rows.Scan(&loanRepayment.PaymentDate, &loanRepayment.Payment, &loanRepayment.RepaymentStatus)
		if err != nil {
			return nil, fmt.Errorf("error scanning loan repayment: %w", err)
		}
		loanRepayments = append(loanRepayments, loanRepayment)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error retrieving loan repayments: %w", err)
	}

	return loanRepayments, nil
}

func NewGoodsRepo(db *sql.DB) GoodsRepo {
	return &goodsRepoImpl{
		db: db,
	}
}
