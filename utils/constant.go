package utils

const (
	GET_ALL_PRODUCT = `SELECT 
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
							mst_product.id ASC`
	GET_PRODUCT_BY_ID = `SELECT 
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
	GET_PRODUCT_BY_NAME = "SELECT id, product_name, description, price, stok, category_product_id, status, created_at, updated_at FROM mst_product WHERE product_name = $1"
	UPDATE_PRODUCT      = "UPDATE mst_product SET product_name = $1, description = $2, price = $3, stok = $4,  category_product_id = $5, status = $6, updated_at = $7 WHERE id = $8"
	DELETE_PRODUCT      = "DELETE FROM mst_product WHERE id = $1;"
	ADD_PRODUCT         = "INSERT INTO mst_product (product_name, description, price, stok, category_product_id, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id"

	GET_CATEGORY_LOAN_BY_ID   = "SELECT id,category_loan_name, created_at, updated_at FROM category_loan WHERE id = $1"
	GET_CATEGORY_LOAN_BY_NAME = "SELECT id, category_loan_name, created_at, updated_at FROM category_loan WHERE category_loan_name = $1"
	GET_ALLCATEGORYLOAN       = "SELECT id, category_loan_name, created_at, updated_at FROM category_loan ORDER BY id ASC"
	INSERT_CATEGORY_LOAN      = "INSERT INTO category_loan ( category_loan_name, created_at, updated_at) VALUES($1, $2, $3) RETURNING id"
	UPDATE_CATEGORY_LOAN      = "UPDATE category_loan SET category_loan_name = $1, updated_at = $2, created_at = $3 WHERE id = $4"
	DELETE_CATEGORYLOAN       = "DELETE FROM category_loan WHERE id = $1"
	GET_CATEGORY_UPDATE_ID    = "SELECT id FROM category_loan WHERE id = $1"

	INSERT_CATEGORY_PRODUCT        = "INSERT INTO category_product(category_product_name, created_at, updated_at) VALUES ($1 ,$2, DEFAULT) RETURNING id"
	DELETE_CATEGORYPRODUCT         = "DELETE FROM category_product WHERE id = $1"
	UPDATE_CATEGORY_PRODUCT        = "UPDATE category_product SET category_product_name = $1, updated_at = $2 WHERE id = $3"
	GET_CATEGORY_PRODUCT_UPDATE_ID = "SELECT id FROM category_product WHERE id = $1"
	GET_ALLCATEGORYPRODUCT         = "SELECT  id, category_product_name, created_at, updated_at FROM category_product ORDER BY id"
	GET_CATEGORY_PRODUCT_BY_ID     = "SELECT id, category_product_name, created_at, updated_at FROM category_product WHERE id = $1"
	GET_CATEGORY_PRODUCT_BY_NAME   = "SELECT id, category_product_name, created_at, updated_at FROM category_product WHERE category_product_name = $1"

	ADD_CUSTOMER           = "INSERT INTO mst_customer(full_name, address, nik, phone, user_id, no_kk,emergency_contact, emergency_name,  last_salary, created_at) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id"
	GET_CUSTOMER_BY_ID     = "SELECT id, full_name, address, nik, phone, user_id, created_at, updated_at, no_kk, emergency_name, emergency_contact, last_salary FROM mst_customer WHERE id = $1"
	GET_ALL_CUSTOMER       = "SELECT * FROM mst_customer ORDER BY id"
	UPDATE_CUSTOMER        = "UPDATE mst_customer SET full_name=$1, address=$2, phone=$3, updated_at=$4,  emergency_name=$5, emergency_contact=$6, last_salary=$7 WHERE id=$8"
	DELETE_CUSTOMER        = "DELETE FROM mst_customer WHERE id=$1"
	GET_CUSTOMER_BY_NIK    = "SELECT id FROM mst_customer WHERE nik = $1"
	GET_CUSTOMER_BY_NUMBER = "SELECT id FROM mst_customer WHERE phone = $1"

	CREATE_APLICATION_LOAN_REPO = `INSERT INTO trx_loan (customer_id, loan_date, due_date, category_loan_id, amount, description, status, repayment_status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`
	GET_CUSTOMER_LOAN_BY_ID     = "SELECT id, nik, no_kk, emergency_name, emergency_contact, last_salary FROM mst_customer WHERE id = $1"
	GET_LOAN_APLICATION         = `SELECT la.id, la.customer_id, la.loan_date, la.due_date, la.category_loan_id, la.amount, la.description, la.status, la.repayment_status, la.created_at, la.updated_at, mc.full_name, mc.address, mc.nik, mc.phone, mc.no_kk, mc.emergency_name, mc.emergency_contact, mc.last_salary FROM trx_loan la INNER JOIN mst_customer mc ON la.customer_id = mc.id ORDER BY la.id ASC OFFSET $1 LIMIT $2`
	GET_LOAN_APLICATION_BY_ID   = `SELECT 
	la.id, 
	la.customer_id, 
	la.loan_date, 
	la.due_date, 
	la.category_loan_id, 
	la.amount, 
	la.description, 
	la.status,
	la.repayment_status, 
	la.created_at, 
	la.updated_at,
	mc.full_name, 
	mc.address, 
	mc.nik, 
	mc.phone, 
	mc.no_kk, 
	mc.emergency_name, 
	mc.emergency_contact, 
	mc.last_salary
FROM 
	trx_loan la
INNER JOIN 
	mst_customer mc ON la.customer_id = mc.id
WHERE
	la.id = $1
ORDER BY la.id`
	GET_LOAN_APLCATION_REPAYMENT_STATUS = `SELECT 
					la.id, 
					la.customer_id, 
					la.loan_date, 
					la.due_date, 
					la.category_loan_id, 
					la.amount, 
					la.description, 
					la.status, 
					la.repayment_status, 
					la.created_at, 
					la.updated_at,
			   		mc.full_name, 
					mc.address, 
					mc.nik, 
					mc.phone, 
					mc.no_kk, 
					mc.emergency_name, 
					mc.emergency_contact, 
					mc.last_salary
				FROM 
					trx_loan la
				INNER JOIN mst_customer mc ON la.customer_id = mc.id
				WHERE la.repayment_status = $3
				ORDER BY la.id ASC
				OFFSET $1 LIMIT $2`
	LOAN_REPAYMENT = "UPDATE trx_loan SET payment_date = $1, payment = $2, repayment_status = $3::status_enum, updated_at = $4 WHERE id = $5"
	GET_LOAN_REPAYMENT_BY_DATE_RANGE = "SELECT payment_date, payment FROM trx_loan WHERE payment_date >= $1 AND payment_date <= $2"

	INSERT_USER       = "INSERT INTO mst_user(id, user_name, email, password, roles_name, is_active, created_at, updated_at ) VALUES($1, $2, $3, $4, $5, $6, $7, $8)"
	UPDATE_USER       = "UPDATE mst_user SET user_name = $1,  email = $2, password = $3, roles_name = $4, is_active = $5 ,created_at = $6, updated_at = $7 WHERE id = $8"
	DELETE_USER       = "DELETE FROM mst_user WHERE id = $1"
	GET_ALL_USER      = "SELECT id,user_name, email, password, roles_name, is_active,created_at, updated_at FROM mst_user ORDER BY id ASC"
	GET_USER_BY_NAME  = "SELECT id,user_name, email, password, roles_name, is_active ,created_at, updated_at FROM mst_user WHERE user_name = $1"
	GET_USER_BY_ID    = "SELECT id, user_name, email, password, roles_name, is_active ,created_at, updated_at FROM mst_user WHERE id = $1"
	GET_USER_BY_EMAIL = "SELECT id,user_name, email, password, roles_name, is_active,created_at, updated_at  FROM mst_user WHERE email = $1"

	INSERT_GOODS                      = "INSERT INTO trx_goods (customer_id, loan_date, payment_date, due_date, category_loan_id, product_id, quantity, price, amount, created_at, status, repayment_status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)"
	GET_GOODS_BY_ID                   = "SELECT g.id, g.customer_id, g.loan_date, g.due_date, g.category_loan_id, g.product_id, p.product_name, g.quantity , p.price , g.amount, p.description, g.status, g.repayment_status, g.created_at, g.updated_at, c.full_name, c.address, c.phone, c.nik, c.no_kk, c.emergency_name, c.emergency_contact, c.last_salary FROM trx_goods g JOIN mst_customer c ON g.customer_id = c.id JOIN mst_product p ON g.product_id = p.id WHERE g.id = $1"
	GET_ALL_TRX_GOODS                 = `SELECT g.id, g.customer_id, g.loan_date, g.due_date, g.category_loan_id, g.product_id, p.product_name, g.quantity , p.price , g.amount, p.description, g.status, g.repayment_status, g.created_at, g.updated_at, c.full_name, c.address, c.phone, c.nik, c.no_kk, c.emergency_name, c.emergency_contact, c.last_salary FROM trx_goods g JOIN mst_customer c ON g.customer_id = c.id JOIN mst_product p ON g.product_id = p.id ORDER BY g.id ASC OFFSET $1 LIMIT $2`
	UPDATE_GOODS_REPAYMENT            = "UPDATE trx_goods SET payment_date = $1, payment = $2, repayment_status = $3, updated_at = $4 WHERE id = $5"
	GET_GOODS_REPAYMENT_STATUS        = `SELECT g.id, g.customer_id, g.loan_date, g.due_date, g.category_loan_id, g.product_id, p.product_name, g.quantity , p.price , g.amount, p.description, g.status, g.repayment_status, g.created_at, g.updated_at, c.full_name, c.address, c.phone, c.nik, c.no_kk, c.emergency_name, c.emergency_contact, c.last_salary FROM trx_goods g JOIN mst_customer c ON g.customer_id = c.id JOIN mst_product p ON g.product_id = p.id WHERE g.repayment_status = $3 ORDER BY g.id ASC OFFSET $1 LIMIT $2`
	GET_GOODS_REPAYMENT_BY_DATE_RANGE = ` SELECT payment_date, payment FROM trx_goods WHERE payment_date >= $1 AND payment_date <= $2`

)

