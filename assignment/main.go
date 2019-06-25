package main

import (
	"fmt"

	"Golang_Tutorial/assignment/domain"
	"Golang_Tutorial/assignment/mapstore"
)

type CustomerController struct {
	// Explicit dependency and declarative programming that hides dependent logic
	store domain.CustomerStore // It can be any Store including MapStore
}

func (c CustomerController) Add(d domain.Customer) {
	err := c.store.Create(d)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("New Customer has been created")
}

func (c CustomerController) GetAll() {
	fmt.Println("Customer list is:\n")
	c.store.GetAll()

}

func (c CustomerController) Update(key string, d domain.Customer) {
	err := c.store.Update(key, d)
	if err != nil {
		fmt.Println(err)
	}
}

// GetByID
func (c CustomerController) GetById(Id string) error {
	err := c.store.GetById(Id)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	fmt.Println("New Customer has been created")
	return nil
}

// main
func main() {
	controller := CustomerController{
		store: mapstore.NewMapStore(), // Inject the dependency
	}

	customer1 := domain.Customer{
		ID:   "cust101",
		Name: "JPM",
	}
	customer2 := domain.Customer{
		ID:   "cust102",
		Name: "Morgan",
	}
	controller.Add(customer1) // Create new Customer
	controller.Add(customer2)
	controller.Update("cust102", customer1)
	controller.GetById("102")
	// Printing all the customers
	//controller.GetAll()
}
