package product

import (
	"ddd/aggregate"

	"github.com/google/uuid"
)

type ProductRepository interface{
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}