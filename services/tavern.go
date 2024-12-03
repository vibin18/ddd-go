package services

import (

	"github.com/google/uuid"
	"fmt"
)

type TavernConfiguration func(t *Tavern) error
 

type Tavern struct{
	OrderService *OrderService

	BillingService interface{}
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error){
	t := &Tavern{}

	for _, cfg := range cfgs{
		if err := cfg(t); err != nil {
			return nil, err
		}
	}

	return t, nil
} 

func WithOrderService(os *OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}

}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error{
	price, err := t.OrderService.CreateOrder(customer,products)
	if err != nil {
		return err
	}

	fmt.Printf("Bill the customer: %0.0f\n", price)
	return nil

}