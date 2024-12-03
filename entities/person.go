// Package entities holds all the entities that are shared across subdomains.
package entities

import "github.com/google/uuid"

// Person is an entity that represents a person in all domains.
type Person struct{
	ID uuid.UUID
	Name string
	Age int
}