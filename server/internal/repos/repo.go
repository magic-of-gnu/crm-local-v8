package repo

import "github.com/magic-of-gnu/crm-local-v8/server/internal/models"

type CentresRepo interface {
	GetAll() ([]models.Centre, error)
}
