package main

import (
	"fmt"
	"time"
)

type Server interface {
	Server(int) int
	String() string
}

type server struct {
	id int
}

func (s *server) Server(fib int) int {
	return fibonacci(fib)
}

func (s *server) String() string {
	return fmt.Sprintf("Server %d", s.id)
}

var NewServer = func(id int) Server {
	return &server{id: id}
}

var CreateNumberOfServers = func(n int) []Server {
	servers := make([]Server, n)
	for i := 0; i < n; i++ {
		servers[i] = NewServer(i+1)
	}
	return servers
}

var fibonacci = func(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
