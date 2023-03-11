package repos

import (
	"time"

	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type CentresRepo interface {
	GetAll() ([]models.Centre, error)
	CreateOne(id int, name, description string, created_at, updated_at time.Time) (*models.Centre, error)
}
