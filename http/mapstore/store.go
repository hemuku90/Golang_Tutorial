package mapstore

import (
	"errors"
	"fmt"
	"http/domain"
	"net/http"

	"Golang_Tutorial/http/domain"
)

// Create a Map store with key as string and value as interface customer of package domain

type MapStore struct {
	store map[string]domain.Customer // An in-memory store with a map
}

// Factory method gives a new instance of MapStore
// This is for caller packages, not for mapstore

func NewMapStore() *MapStore {
	return &MapStore{store: make(map[string]domain.Customer)}
}

//MapStore implementation of Create
func (m *MapStore) Create(c domain.Customer) error {
	if _, ok := m.store[c.ID]; ok {
		return errors.New("Customer already exists")
	}
	m.store[c.ID] = c
	return nil
}

func (m *MapStore) GetAll(http.ResponseWriter,
	*http.Request) {

	fmt.Println(m)

}

func (m *MapStore) Update(key string, value domain.Customer) error {

	if _, ok := m.store[value.ID]; ok {

		m.store[key] = value
		fmt.Print(m)
		return nil
	}

	return errors.New("Customer does not exist")

}

func (m *MapStore) Delete(key string) error {

	if _, ok := m.store[key]; ok {
		delete(m.store, key)
		fmt.Println("After Deletion :", m)
		return nil
	} else {
		return errors.New("Customer does not exist")
	}

}

func (m *MapStore) GetById(key string) (domain.Customer, error) {
	if cus, ok := m.store[key]; ok {

		fmt.Println("Element is :", m.store[key])
		fmt.Println("Map is :", m)

		return cus, nil

	} else {
		return cus, errors.New("Element not found")
	}

}
