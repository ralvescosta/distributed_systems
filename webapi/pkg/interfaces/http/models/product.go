package models

import "webapi/pkg/domain/dtos"

type ProductModel struct {
	Id              string   `json:"id"`
	ProductCategory string   `json:"product_category"`
	Tag             string   `json:"tag"`
	Title           string   `json:"title"`
	Subtitle        string   `json:"subtitle"`
	Authors         []string `json:"authors"`
	AmountInStock   int      `json:"amount_in_stock"`
	NumPages        int      `json:"num_pages"`
	Tags            []string `json:"tags"`
	CreatedAt       string   `json:"created_at"`
	UpdatedAt       string   `json:"updated_at"`
}

func ToGetByIdResponse(dto dtos.ProductDto) ProductModel {
	return ProductModel{}
}
