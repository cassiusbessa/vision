package entities

import "github.com/google/uuid"

type Author struct {
	ID    uuid.UUID
	Name  string
	Image string
}
