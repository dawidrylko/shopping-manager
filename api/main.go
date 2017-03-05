package main

import (
	"fmt"

	"./product"
)

func main() {
	fmt.Println("shopping-manager api start")

	resourceHandler()
}

// resourceHandler is unexported function
// which includes all resources
func resourceHandler() {
	product.Resource()
}
