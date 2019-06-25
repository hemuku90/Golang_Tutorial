package main

import (
	"fmt"
	"http/domain"
	"http/mapstore"
)

type CustomerController struct {
	// Explicit dependency and declarative programming that hides dependent logic
	store domain.CustomerStore // It can be any Store including MapStore
}

func (c CustomerController) Add(cus domain.Customer) {
	err := c.store.Create(cus)

	if err != nil {
		fmt.Println("Error : ", err)
		return
	}

	fmt.Println("New Customer has been created")
}
func (c CustomerController) GetAll() {

	fmt.Println("Customer List")
	c.store.GetAll()

}

func (c CustomerController) Update(key string, value domain.Customer) error {
	err := c.store.Update(key, value)

	if err != nil {
		fmt.Println("Error : ", err)
		return err
	}

	return nil
}
func (c CustomerController) Delete(key string) error {
	err := c.store.Delete(key)

	if err != nil {
		fmt.Println("Error : ", err)
		return err
	}

	return nil
}

func (c CustomerController) GetById(key string) (domain.Customer, error) {
	cus, err := c.store.GetById(key)

	if err != nil {
		fmt.Println("Error : ", err)
		return cus, err
	}

	return cus, nil
}

func main() {

	// EntryPoint

	controller := CustomerController{
		store: mapstore.NewMapStore(), // Inject the dependency
	}

	customer1 := domain.Customer{
		ID:    "1",
		Email: "MK",
		Name:  "12",
	}
	customer2 := domain.Customer{
		ID:    "2",
		Email: "MK",
		Name:  "12",
	}
	controller.Add(customer1)
	controller.Add(customer2)
	controller.GetAll()
	controller.GetById(customer1.ID)

}
