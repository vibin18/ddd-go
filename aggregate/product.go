package aggregate

import (
	"ddd/entities"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrMissingValue = errors.New("missing important value")
)

type Product struct{
	item entities.Item
	price float64
	quantity int
}


func NewProduct(name, description string, price float64)(Product, error){
	if name == "" || description == "" {
		return Product{}, ErrMissingValue
	}

	return Product{
		item: entities.Item{
			ID: uuid.New(),
			Name: name,
			Description: description,

		},
		price: price,
		quantity: 0,
	},nil
}

func (p Product) GetID() uuid.UUID{
	return p.item.ID
}
func (p Product) GetItem() *entities.Item{
	return &p.item
}
func (p Product) GetPrice() float64{
	return p.price
}