package models

import (
	"time"

	"github.com/google/uuid"
)

type UserPermissions struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	PermissionID uuid.UUID
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
