// Package memory is an in-memory implimentation of Customer Repo.
package memory

import (
	"ddd/aggregate"
	"ddd/domain/customer"
	"fmt"
	"sync"

	"github.com/google/uuid"
)


var mu sync.Mutex

type Repository struct {
	customers map[uuid.UUID]aggregate.Customer

}

func New() *Repository {
	return &Repository{
		customers: make(map[uuid.UUID]aggregate.Customer),
	}
}

func(m *Repository)Get(id uuid.UUID) (aggregate.Customer, error){
	if customer, ok := m.customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func(m *Repository)Add(c aggregate.Customer)error {
	if m.customers == nil {
		mu.Lock()
		m.customers = make(map[uuid.UUID]aggregate.Customer)
		mu.Unlock()		
	}

	// Makesure customer not in the repo
	mu.Lock()
	if _,ok := m.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exist! : %w", customer.ErrFailedAddCustomer)
	}

	m.customers[c.GetID()] = c
	mu.Unlock()
 	return nil

}
func(m *Repository)Update(c aggregate.Customer) error{

	if _,ok := m.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exist! : %w", customer.ErrUpdateCustomer)
	}

	mu.Lock()
	m.customers[c.GetID()] = c
	mu.Unlock()
 	return nil
	
}