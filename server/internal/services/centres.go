package services

import (
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
)

type centresService struct {
	CentresRepo repos.CentresRepo
}

func NewCentresService(centresRepo repos.CentresRepo) *centresService {
	return &centresService{
		CentresRepo: centresRepo,
	}
}

func (ss *centresService) GetAll() ([]models.Centre, error) {
	return ss.CentresRepo.GetAll()
}
