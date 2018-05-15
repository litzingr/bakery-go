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
	done := make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				m.logger.Println("ending")
				wg.Wait()
				done <- struct{}{}
				return
			case c := <-m.customerList:
				s := <-m.serverList
				wg.Add(1)
				go func() {
					o := s.Serve(c.Order())
					m.serverList <- s;
					m.logger.Println("Served a customer the number: ", o)
					m.customerList <- c
					wg.Done()
				}()
			}
		}
	}()
}

func (m *manager) Customers(customers []Customer) {
	m.customers = customers
	m.customerList = make(chan Customer, len(customers)+1)
	for _, c := range m.customers {
		m.customerList <- c
	}
}

func (m *manager) Servers(servers []Server) {
	m.servers = servers
	m.serverList = make(chan Server, len(servers)+1)
	for _, b := range m.servers {
		m.serverList <- b
	}
}

func HireManager(logger *log.Logger) Manager {
	return &manager{logger: logger}
}
