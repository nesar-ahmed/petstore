package models

// Breed: defines a breed of a pet
type Breed struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	CategoryID int    `json:"categoryId"`
}

// BreedService: interface for Breed model
type BreedService interface {
	// Create a breed
	CreateBreed() error

	// Get a breed based on category id
	GetBreedByCategory(id int) ([]*Breed, error)
}
