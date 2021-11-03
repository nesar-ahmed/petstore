package models

// Category: defines a category of a pet
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// CategoryService: interface for Category model
type CategoryService interface {
	// Create a category
	CreateCategory() error

	// Fetch a category based on given id
	GetCategory(id int) (*Category, error)

	// Fetch all categories
	GetAllCategory() ([]*Category, error)
}
