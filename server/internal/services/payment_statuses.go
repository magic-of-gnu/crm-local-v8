package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
)

type paymentStatusesService struct {
	PaymentStatusesRepo repos.PaymentStatusRepo
}

func NewPaymentStatusesService(
	repo repos.PaymentStatusRepo,
) *paymentStatusesService {
	return &paymentStatusesService{
		PaymentStatusesRepo: repo,
	}
}

func (ss *paymentStatusesService) GetAll() ([]models.PaymentStatus, error) {

	var response []models.PaymentStatus

	response, err := ss.PaymentStatusesRepo.GetAll()
	if err != nil {
		return response, err
	}

	return response, nil

}

func (ss *paymentStatusesService) CreateOne(createItem *models.PaymentStatus) (*models.PaymentStatus, error) {

	var item *models.PaymentStatus

	uid, err := uuid.NewRandom()
	if err != nil {
		return item, err
	}

	time_now := time.Now()

	createItem.ID = uid
	createItem.CreatedAt = time_now
	createItem.UpdatedAt = time_now

	item, err = ss.PaymentStatusesRepo.CreateOne(createItem)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (ss *paymentStatusesService) GetOneByID(uid uuid.UUID) (*models.PaymentStatus, error) {

	item, err := ss.PaymentStatusesRepo.GetOneByID(uid)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (ss *paymentStatusesService) DeleteOneByID(uid uuid.UUID) error {

	err := ss.PaymentStatusesRepo.DeleteOneByID(uid)
	if err != nil {
		return err
	}

	return nil
}

func (ss *paymentStatusesService) UpdateOneByID(updateItem *models.PaymentStatus) error {

	// getItem, err := ss.PaymentStatusesRepo.GetOneByID(updateItem.ID)

	err := ss.PaymentStatusesRepo.UpdateOneByID(updateItem)
	if err != nil {
		return err
	}

	return nil
}
