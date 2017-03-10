package main

import (
	"fmt"

	"./resources/product"
	"./server"
)

func main() {
	fmt.Println("shopping-manager api start")

	loadResourceHandler()

	server.Start()
}

// loadResourceHandler is unexported function
// which includes all resources
func loadResourceHandler() {
	product.Handler()
}
