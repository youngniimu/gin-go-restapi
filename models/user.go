package models

import "github.com/google/uuid"

type User struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	PersonalCode uuid.UUID `json:"personalCode"`
}