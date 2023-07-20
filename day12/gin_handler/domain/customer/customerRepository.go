package customer

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("the customer was not found in the repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	ErrUpdateCustomer      = errors.New("failed to update the customer in the repository")
)

type CustomerRepository interface {
	Get(uuid.UUID) (Customer, error)
	Add(Customer) error
	GetAll() ([]Customer, error)
}
