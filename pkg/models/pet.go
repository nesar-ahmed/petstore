package models

// Pet: defines a pet model
type Pet struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Age         string   `json:"age"`
	Description string   `json:"description"`
	ImageUrl    string   `json:"imageUrl"`
	PetLocation Location `json:"petLocation"`
	PetCategory Category `json:"petCategory"`
	PetBreed    Breed    `json:"petBreed"`
}

// PetService: interface for Pet model
type PetService interface {
	// Create a pet
	CreatePet() error

	// Get a pet based on id
	GetPet(id int) (*Pet, error)

	// Get a list of pets based on category id
	GetPetByCategory(id int) ([]*Pet, error)

	// Delete a pet based on given id
	DeletePet(id int) error
}
