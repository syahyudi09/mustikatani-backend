package repository

import (
	"database/sql"
	"fmt"
	"pinjam-modal-app/model"
)

type ProductRepo interface {
	CreateProduct(newProduct *model.ProductModel) error
	GetAllProduct(page, limit int) ([]*model.ProductModel, error)
	GetProductById(id int) (*model.ProductModel, error)
	GetProductByName(nameProduct string) (*model.ProductModel, error)
	UpdateProduct(id int, updateProduct *model.ProductModel) error
	DeleteProduct(id int) error
}

type productRepo struct {
	db *sql.DB
}

func (p *productRepo) CreateProduct(newProduct *model.ProductModel) error {
	insertStatement := "INSERT INTO mst_product (product_name, description, price, stok, category_product_id, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"
	err := p.db.QueryRow(insertStatement, newProduct.ProductName, newProduct.Description, newProduct.Price, newProduct.Stok, newProduct.CategoryProductId, newProduct.Status, newProduct.CreatedAt, newProduct.UpdatedAt).Scan(&newProduct.Id)
	if err != nil {
		return err
	}
	return nil
}

func (p *productRepo) GetAllProduct(page, limit int) ([]*model.ProductModel, error) {
	offset := (page - 1) * limit
	selectStatement := `SELECT 
							mst_product.id, 
							mst_product.product_name, 
							mst_product.description, 
							mst_product.price, 
							mst_product.stok, 
							mst_product.category_product_id,
							category_product.category_product_name,
							mst_product.status, 
							mst_product.created_at, 
							mst_product.updated_at
						FROM 
							mst_product
						INNER JOIN 
							category_product ON mst_product.category_product_id = category_product.id
						ORDER BY 
							mst_product.id ASC
						OFFSET $1 LIMIT $2`

	rows, err := p.db.Query(selectStatement, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("GetAllProduct() : %w", err)
	}
	defer rows.Close()

	var products []*model.ProductModel
	for rows.Next() {
		product := &model.ProductModel{}
		err := rows.Scan(
			&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Stok, &product.CategoryProductId, &product.CategoryProductName,
			&product.Status, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("GetAllProduct() : %w", err)
		}
		products = append(products, product)
	}

	return products, nil
}

func (p *productRepo) GetProductById(id int) (*model.ProductModel, error) {
	selectStatement := `SELECT 
							mst_product.id,
							mst_product.product_name, 
							mst_product.description, 
							mst_product.price, 
							mst_product.stok, 
							mst_product.category_product_id,
							category_product.category_product_name,
							mst_product.status, 
							mst_product.created_at, 
							mst_product.updated_at
						FROM 
							mst_product
						INNER JOIN 
							category_product ON mst_product.category_product_id = category_product.id
						WHERE
							mst_product.id = $1
						ORDER BY 
							mst_product.id ASC`

	row := p.db.QueryRow(selectStatement, id)

	product := &model.ProductModel{}
	err := row.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Stok, &product.CategoryProductId, &product.CategoryProductName,
		&product.Status, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("GetProductById() : %w", err)
	}

	return product, nil
}

func (p *productRepo) GetProductByName(nameProduct string) (*model.ProductModel, error) {
	selectStatement := "SELECT id, product_name, description, price, stok, category_product_id, status, created_at, updated_at FROM mst_product WHERE product_name = $1"

	row := p.db.QueryRow(selectStatement, nameProduct)

	product := &model.ProductModel{}
	err := row.Scan(&product.Id, &product.ProductName, &product.Description, &product.Price, &product.Stok, &product.CategoryProductId, &product.Status, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("GetProductByName() : %w", err)
	}

	return product, nil
}

func (p *productRepo) UpdateProduct(id int, updateProduct *model.ProductModel) error {
	updateStatment := "UPDATE mst_product SET product_name = $1, description = $2, price = $3, stok = $4,  category_product_id = $5, status = $6, updated_at = $7 WHERE id = $8"
	_, err := p.db.Exec(updateStatment, updateProduct.ProductName, updateProduct.Description, updateProduct.Price, updateProduct.Stok, updateProduct.CategoryProductId, updateProduct.Status, updateProduct.UpdatedAt, id)
	if err != nil {
		return fmt.Errorf("error on productRepo.UpdateProduct() : %w", err)
	}
	return nil
}

func (p *productRepo) DeleteProduct(id int) error {
	deleteStatment := "DELETE FROM mst_product WHERE id = $1;"

	_, err := p.db.Exec(deleteStatment, id)
	if err != nil {
		return fmt.Errorf("DeleteProduct() : %w", err)
	}

	return nil
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return &productRepo{
		db: db,
	}
}
