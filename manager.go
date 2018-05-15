package main

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"
)

type Manager interface {
	Open(context.Context) (<-chan struct{}, error)
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

func (m *manager) Open(ctx context.Context) (<-chan struct{}, error) {
	var wg sync.WaitGroup
	done := make(chan struct{}, 1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				m.logger.Println("ending")
				wg.Wait()
				cancel()
				done <- struct{}{}
				return
			case c := <-m.customerList:
				s := <-m.serverList
				wg.Add(1)
				go func() {
					o := s.Serve(c.Order())
					m.serverList <- b
					m.logger.Println("Served a customer")
					m.customerList <- c
					wg.Done()
				}()
			}
		}
	}()
	return done, nil
}

func (m *manager) Customers(customers []Customer) {
	m.customers = customers
	m.customerList = make(chan Customer, len(customers)+1)
	for _, c := range m.customers {
		m.customerLIst <- c
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
