package main

import (
	"context"
	"log"
	"sync"
)

type Manager interface {
	Start(context.Context)
	Customers([]Customer)
	Servers([]Server)
}

type manager struct {
	customers []Customer
	servers []Server
	serverList chan Server
	customerList chan Customer
	logger *log.Logger
}

func (m *manager) Start(ctx context.Context) {
	var wg sync.WaitGroup
		for {
			c := <-m.customerList
			s := <-m.serverList
			wg.Add(1)
			o := s.Serve(c.Order())
			m.logger.Println("Server", s.String(), "served customer", c.String(), " the number: ", o)
			m.serverList <- s;
			m.customerList <- c
			wg.Done()
		}
}

func (m *manager) Customers(customers []Customer) {
	m.logger.Println("finding", len(customers), "customers")
	m.customers = customers
	m.customerList = make(chan Customer, len(customers)+1)
	for _, c := range m.customers {
		m.customerList <- c
	}
}

func (m *manager) Servers(servers []Server) {
        m.logger.Println("hiring", len(servers), "servers")
	m.servers = servers
	m.serverList = make(chan Server, len(servers)+1)
	for _, b := range m.servers {
		m.serverList <- b
	}
}

func HireManager(logger *log.Logger) Manager {
	return &manager{logger: logger}
}
