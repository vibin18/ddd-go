package memory

import (
	"ddd/aggregate"
	"errors"
	"sync"

	"github.com/google/uuid"
)

var(
	ErrProductNotfound = errors.New("no such product")
	ErrProductAlreadyExist = errors.New("product already exist")
)

type MemoryProductRepository struct {
	products map[uuid.UUID]aggregate.Product
	mu sync.Mutex
}


func New() *MemoryProductRepository {
	return &MemoryProductRepository{
		products: make(map[uuid.UUID]aggregate.Product),
	}
}


func (m *MemoryProductRepository) GetAll()([]aggregate.Product, error){
	var products []aggregate.Product
	for _, prd := range m.products {
		products = append(products, prd)
	}
	return products, nil
}

func (m *MemoryProductRepository) GetByID(id uuid.UUID) (aggregate.Product, error){
	if prd, ok := m.products[id]; ok{
		return prd, nil
	}
	return aggregate.Product{}, ErrProductNotfound
}
func (m *MemoryProductRepository) Add(product aggregate.Product) error {

	if _, ok := m.products[product.GetID()]; ok{
		return ErrProductAlreadyExist
	}
	m.mu.Lock()
	m.products[product.GetID()] = product
	m.mu.Unlock()
	return nil
}
func (m *MemoryProductRepository) Update(product aggregate.Product) error {
	if _, ok := m.products[product.GetID()]; !ok{
		return ErrProductNotfound
	}

	m.mu.Lock()
	m.products[product.GetID()] = product
	m.mu.Unlock()
	return nil
}
func (m *MemoryProductRepository) Delete(id uuid.UUID) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.products[id]; !ok{
		return ErrProductNotfound
	}
	
	delete(m.products, id)
	return nil
}