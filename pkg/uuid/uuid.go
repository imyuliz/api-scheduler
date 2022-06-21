package uuid

import "github.com/google/uuid"

// NewUUID return uuid string
func NewUUID() string {
	return uuid.NewString()
}
