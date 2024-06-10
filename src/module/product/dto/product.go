package dto

import "github.com/linggaaskaedo/go-playground-wire-v3/src/module/product/entity"

type ProductRespBody struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Cost        float32  `json:"cost"`
	Price       float32  `json:"price"`
	Category    Category `json:"category"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func CreateProductResp(data entity.Product) ProductRespBody {
	return ProductRespBody{
		ID:          data.ID,
		Name:        data.Name,
		Description: data.Description,
		Cost:        data.Cost,
		Price:       data.Price,
		Category: Category{
			ID:   data.Category.ID,
			Name: data.Category.Name,
		},
	}
}
