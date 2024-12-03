// Package aggregate holds our aggrets that combines many entities  into a full Object
package aggregate

import (
	"ddd/entities"
	"ddd/valueobject"
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("a customer has tobe a valid name")
)

type Customer struct{
	// person is the root entity of customer
	// whihc means the person.ID is the main identifies for the customer
	person *entities.Person
	products []*entities.Item
	transation []valueobject.Transtation
}


// NewCustomer is a factory to create a new customer aggregate
// It will valdate that the name is not empty

func NewCustomer(name string) (Customer, error){
	if name == "" {
		return Customer{}, ErrInvalidPerson 
	}

	c := &entities.Person{
		Name: name,
		ID: uuid.New(),
	}
	return Customer{
		person: c,
		products: make([]*entities.Item, 0),
		transation: make([]valueobject.Transtation, 0),
	}, nil
}

func(c  *Customer) GetID() uuid.UUID{
	return c.person.ID
}
func(c  *Customer) SetID(id uuid.UUID){
	if c.person == nil {
		c.person = &entities.Person{}
	}
	c.person.ID = id
}
func(c  *Customer) SetName(name string){
	if c.person == nil {
		c.person = &entities.Person{}
	}
	c.person.Name = name
}
func(c  *Customer) GetName(name string)string{
	return c.person.Name
}