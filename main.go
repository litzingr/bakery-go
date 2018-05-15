package main

import (
	"io"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"context"
	"fmt"
	"time"
	"log"
)

var (
	exit = os.Exit
	stderr io.Writer = os.Stderr
)

func main() {
	ctx, calcel := context.WithCancel(context.Background())
	logger := log.New(stderr, "", log.Lshortfile)
	m := HireManager()
	m.Customers()
	m.Servers()
	m.Start()

}
