package storage

import "github.com/google/uuid"

// On error(exceptional case) it will panic - highly unlikely to happen(entropy)
func NewUUID() string {
	return uuid.Must(uuid.NewV7()).String()
}
