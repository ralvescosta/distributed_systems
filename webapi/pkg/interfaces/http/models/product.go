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

type CreateProductModel struct {
	ProductCategory string   `json:"product_category" validate:"required"`
	Tag             string   `json:"tag" validate:"required"`
	Title           string   `json:"title" validate:"required"`
	Subtitle        string   `json:"subtitle" validate:"required"`
	Authors         []string `json:"authors" validate:"required"`
	AmountInStock   int      `json:"amount_in_stock" validate:"required"`
	NumPages        int      `json:"num_pages" validate:"required"`
	Tags            []string `json:"tags" validate:"required"`
}

func (pst CreateProductModel) ToProductDto() dtos.ProductDto {
	return dtos.ProductDto{
		ProductCategory: pst.ProductCategory,
		Tag:             pst.Tag,
		Title:           pst.Title,
		Subtitle:        pst.Subtitle,
		Authors:         pst.Authors,
		AmountInStock:   pst.AmountInStock,
		NumPages:        pst.NumPages,
		Tags:            pst.Tags,
	}
}

func ToProductResponse(dto dtos.ProductDto) ProductModel {
	return ProductModel{
		Id:              dto.Id,
		ProductCategory: dto.ProductCategory,
		Tag:             dto.Tag,
		Title:           dto.Title,
		Subtitle:        dto.Subtitle,
		Authors:         dto.Authors,
		AmountInStock:   dto.AmountInStock,
		NumPages:        dto.NumPages,
		Tags:            dto.Tags,
		CreatedAt:       dto.CreatedAt,
		UpdatedAt:       dto.UpdatedAt,
	}
}
