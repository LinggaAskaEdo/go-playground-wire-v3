package repository

const (
	GetProductByID = `
		SELECT
			p.product_id, p.product_name , p.description, p.standard_cost, p.list_price, pc.category_id, pc.category_name
		FROM
			products p INNER JOIN product_categories pc on p.category_id = pc.category_id
		WHERE
			p.product_id = $1
	`
)
