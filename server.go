package main

import (
	"fmt"
)

type Server interface {
	Serve(int) int
	String() string
}

type server struct {
	id int
}

func (s *server) Serve(fib int) int {
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
        f := make([]int, n+1, n+2)
	if n <= 1 {
		f = f[0:2]
	}
	f[0] = 0
	f[1] = 1
	for i :=2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
