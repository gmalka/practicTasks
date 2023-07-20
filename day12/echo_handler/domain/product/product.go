package product

import (
	"echohandlertest/models"
	"errors"

	"github.com/google/uuid"
)

var ErrMissingValues = errors.New("missing values")

type Product struct {
	item     *models.Item
	price    float64
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &models.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}

func (p Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p Product) GetItem() *models.Item {
	return p.item
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p Product) SetPrice(price float64) {
	p.price = price
}
