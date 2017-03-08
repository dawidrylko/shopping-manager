package main

import (
	"fmt"

	"./resources/product"
	"./server"
)

func main() {
	fmt.Println("shopping-manager api start")

	resourceHandler()
	server.Start()
}

// resourceHandler is unexported function
// which includes all resources
func resourceHandler() {
	product.Resource()
}
