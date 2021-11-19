package dtos

type ProductDto struct {
	Id              string
	ProductCategory string
	Tag             string
	Title           string
	Subtitle        string
	Authors         []string
	AmountInStock   int
	NumPages        int
	Tags            []string
	CreatedAt       string
	UpdatedAt       string
}
