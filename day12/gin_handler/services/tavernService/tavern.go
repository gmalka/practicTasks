package tavernservice

import (
	"echohandlertest/domain/customer"
	"echohandlertest/domain/product"
	orderservice "echohandlertest/services/orderService"
	"log"

	"github.com/google/uuid"
)

type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	orderService *orderservice.OrderService

	customers customer.CustomerRepository
	products  product.ProductRepository
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	tavern := &Tavern{}

	for _, cfg := range cfgs {
		err := cfg(tavern)
		if err != nil {
			return nil, err
		}
	}

	return tavern, nil
}

func WithOrderService(os *orderservice.OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.orderService = os
		return nil
	}
}

func WithCustomerRepository(cr customer.CustomerRepository) TavernConfiguration {
	return func(t *Tavern) error {
		t.customers = cr
		return nil
	}
}

func WithProductRepository(pr product.ProductRepository) TavernConfiguration {
	return func(t *Tavern) error {
		t.products = pr
		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.orderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}
	log.Printf("Bill the Customer: %0.0f", price)

	return nil
}

func (t *Tavern) NewCustomer(customer customer.Customer) error {
	return t.customers.Add(customer)
}

func (t *Tavern) AllCustomers() ([]customer.Customer, error) {
	return t.customers.GetAll()
}

func (t *Tavern) AllProducts() ([]product.Product, error) {
	return t.products.GetAll()
}

func (t *Tavern) NewProduct(product product.Product) error {
	return t.products.Add(product)
}