package product

// Product represents information about product
type Product struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Products represents information about product list
type Products []Product
