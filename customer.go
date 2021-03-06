package main

import (
	"time"
	"math/rand"
)

type Customer interface {
	Order() int
	String() int
}

type customer struct {
	id int
}

func (c *customer) Order() int {
	return randInt(21)
}

func (c *customer) String() int {
	return c.id
}

func randInt(max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return rand.Intn(max)
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
