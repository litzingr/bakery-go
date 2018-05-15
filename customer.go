package main

import (
	"fmt"
)

type Customer interface {
	
}

type customer struct {
	id int
}

func NewCustomer(id int) Customer {
	return &customer{id: id}
}

func CreateNumberOfCustomers(n int) []Customer {
	customers := make([]Customer, n)
	for i := 0; i < n; i++ {
		customers[i] = NewCustomer(i+1)
	}
	return customers
}
