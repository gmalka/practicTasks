package customermemoryrepository

import (
	"echohandlertest/domain/customer"
	"sync"

	"github.com/google/uuid"
)

type MemoryRepository struct {
	customers map[uuid.UUID]customer.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]customer.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (customer.Customer, error) {
	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()
	if v, ok := mr.customers[id]; ok {
		return v, nil
	}

	return customer.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c customer.Customer) error {
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]customer.Customer)
		mr.Unlock()
	}

	mr.Lock()
	if _, ok := mr.customers[c.GetID()]; ok {
		mr.Unlock()
		return customer.ErrFailedToAddCustomer
	}

	mr.customers[c.GetID()] = c
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) GetAll() ([]customer.Customer, error) {
	var customers []customer.Customer
	mr.Mutex.Lock()
	defer mr.Mutex.Unlock()
	for _, customer := range mr.customers {
		customers = append(customers, customer)
	}
	return customers, nil
}