package modelsjson

import (
	"github.com/google/uuid"
)

type Customer struct {
	Person   *Person `json:"person"`
	Products []*Item `json:"products"`
}

type Person struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Name  string    `json:"name"`
	Email string    `json:"email"`
}

type Item struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
