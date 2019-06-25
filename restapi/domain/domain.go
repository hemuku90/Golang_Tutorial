package domain

// A struct for customer

type Customer struct {
	ID, Name, Email string
}

// An interface for CRUD operation

type CustomerStore interface {
	Create(Customer) error
	Update(string, Customer) error
	Delete(string) error
	GetById(string) (Customer, error)
	GetAll()
}
