package main

import (
	"fmt"
	"time"

	"./resources/product"
	"./server"
)

func main() {
	fmt.Println(time.Now(), "Shopping Manager api start")

	loadResourceHandler()

	server.Start()
}

// loadResourceHandler is unexported function
// which includes all resources
func loadResourceHandler() {
	fmt.Println(time.Now(), "Load resource handler")

	product.Handler()
}
