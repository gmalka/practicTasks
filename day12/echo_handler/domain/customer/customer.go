package customer

import (
	"echohandlertest/models"
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidPerson = errors.New("a customer has to have an valid person")

type Customer struct {
	person   *models.Person
	products []*models.Item
}

func NewCustomer(name, ermail string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &models.Person{
		ID:    uuid.New(),
		Name:  name,
		Email: ermail,
	}

	return Customer{
		person:   person,
		products: make([]*models.Item, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}

func (c *Customer) SetId(id uuid.UUID) {
	if c.person == nil {
		c.person = &models.Person{}
	}
	c.person.ID = id
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &models.Person{}
	}
	c.person.Name = name
}

func (c *Customer) GetName() string {
	return c.person.Name
}

func (c *Customer) SetEmail(email string) {
	if c.person == nil {
		c.person = &models.Person{}
	}
	c.person.Email = email
}

func (c *Customer) GetEmail() string {
	return c.person.Email
}
