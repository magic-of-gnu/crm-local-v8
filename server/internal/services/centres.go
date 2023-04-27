package services

import (
	"time"

	"github.com/google/uuid"
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

func (ss *centresService) CreateOne(name, description string) (*models.Centre, error) {
	var centre *models.Centre

	uid, err := uuid.NewRandom()
	if err != nil {
		return centre, err
	}

	time_now := time.Now()

	centre, err = ss.CentresRepo.CreateOne(uid, name, description, time_now, time_now)
	if err != nil {
		return centre, err
	}

	return centre, nil
}

func (ss *centresService) DeleteOneByID(uid uuid.UUID) error {

	err := ss.CentresRepo.DeleteOneByID(uid)
	if err != nil {
		return err
	}

	return nil
}

func (ss *centresService) GetAll() ([]models.Centre, error) {
	return ss.CentresRepo.GetAll()
}

func (ss *centresService) GetOneByID(uid uuid.UUID) (*models.Centre, error) {
	return ss.CentresRepo.GetOneByID(uid)
}

func (ss *centresService) UpdateOneByID(updateItem *models.Centre) error {
	return ss.CentresRepo.UpdateOneByID(updateItem)
}
