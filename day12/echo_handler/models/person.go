package models

import uuid "github.com/google/uuid"

type Person struct {
	ID    uuid.UUID
	Name  string
	Email string
}
