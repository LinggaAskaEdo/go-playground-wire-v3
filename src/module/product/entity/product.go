package entity

type Product struct {
	ID          int      `db:"product_id"`
	Name        string   `db:"product_name"`
	Description string   `db:"description"`
	Cost        float32  `db:"standard_cost"`
	Price       float32  `db:"list_price"`
	Category    Category `db:"category_id"`
}

type Category struct {
	ID   int    `db:"category_id"`
	Name string `db:"category_name"`
}
