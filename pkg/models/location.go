package models

// Location: defines a location of a pet
type Location struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// LocationService: interface for Location model
type LocationService interface {
	// Add a new location
	CreateLocation() error

	// Get a location
	GetLocation(id int) (*Location, error)

	// Get all locations
	GetAllLocation() ([]*Location, error)
}
