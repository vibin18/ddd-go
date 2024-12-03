package services

import (
	"ddd/aggregate"
	"testing"

	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product{
	mem, err := aggregate.NewProduct("Memory", "16GB DDR", 150.00)
	if err != nil {
		t.Fatal(err)
	}

	cpu, err := aggregate.NewProduct("CPU", "AMD Ryzen 7", 450.00)
	if err != nil {
		t.Fatal(err)
	}

	gpu, err := aggregate.NewProduct("GPU", "NVIDEA 4070", 550.00)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{
	mem, cpu, gpu,
	}
}

func TestOrder_NewOrderService(t *testing.T){
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("vibin")
	if err != nil {
		t.Error(err)
	}

	err = os.customers.Add(cust)
	if err != nil {
		t.Error(err)
	}

	// Odering CPU
	order := []uuid.UUID{
		products[1].GetID(),
	}

	_, err = os.CreateOrder(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}


}