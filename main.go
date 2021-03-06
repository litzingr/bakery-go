package main

import (
	"io"
	"os"
	"context"
	"log"
)

var (
	exit = os.Exit
	stderr io.Writer = os.Stderr
)

func main() {
	ctx:= context.Background()
	logger := log.New(stderr, "", log.Lshortfile)
	m := HireManager(logger)
	m.Customers(CreateNumberOfCustomers(10))
	m.Servers(CreateNumberOfServers(5))
	m.Start(ctx)

}
