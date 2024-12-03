package services

import (
	"ddd/aggregate"
	"ddd/domain/customer"
	"ddd/domain/customer/memory"
	"ddd/domain/product"
	prodmem "ddd/domain/product/memory"
	"fmt"

	"github.com/google/uuid"
)


type OrderConfiguration func(o *OrderService) error 

type OrderService struct{
	customers customer.Repository
	products product.ProductRepository
}


func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error){
	os := &OrderService{}
	// Loop through all the configs and apply them
	for _, cfg := range cfgs{
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

// WithCustomerRepository applies a customer repository to the OrderService
func WithCustomerRepository(cr customer.Repository) OrderConfiguration {
	// Return a function that matches the OrderConfiguration alias.
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration{
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration{
	
	return func(os *OrderService) error {
		pr := prodmem.New()

		for _,p := range products{
			if err := pr.Add(p); err != nil{
				return err
			}
		}

		os.products = pr 
		return nil
	}
}

func (o *OrderService) CreateOrder(CustomerId uuid.UUID, productIds []uuid.UUID)(float64, error) {
	//Fetch the customer
	c, err := o.customers.Get(CustomerId)
	if err != nil {
		return 0,err
	}
	var products []aggregate.Product
	var total float64

	for _, id := range productIds {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}

	fmt.Printf("Customer %v has ordered %v products\n", c.GetID(),len(products))
	return total,nil

}


