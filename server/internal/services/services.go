package services

import (
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type CentresService interface {
	GetAll() ([]models.Centre, error)
	CreateOne(name, description string) (*models.Centre, error)
}
