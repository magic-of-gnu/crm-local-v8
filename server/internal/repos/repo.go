package repos

import (
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type CentresRepo interface {
	GetAll() ([]models.Centre, error)
	CreateOne(id uuid.UUID, name, description string, created_at, updated_at time.Time) (*models.Centre, error)
}
