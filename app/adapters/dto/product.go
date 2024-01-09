package dto

import "github.com/marielifm/hexagonal-architecture/app"

type Product struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Status string  `json:"status"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) Bind(product *app.Product) (*app.Product, error) {
	if p.Id != "" {
		product.ID = p.Id
	}
	product.Name = p.Name
	product.Price = p.Price
	product.Status = p.Status
	_, err := product.IsValid()
	if err != nil {
		return &app.Product{}, err
	}

	return product, nil
}
