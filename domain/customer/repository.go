package customer

import (
	"ddd/aggregate"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound = errors.New("a customer not found in the repository")
	ErrFailedAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer = errors.New("failed to update customer")
)

type Repository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer)error
	Update(aggregate.Customer) error
}