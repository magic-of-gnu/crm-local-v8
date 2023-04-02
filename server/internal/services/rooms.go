package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
)

type roomsService struct {
	RoomsRepo repos.RoomsRepo
}

func NewRoomsService(roomsRepo repos.RoomsRepo) *roomsService {
	return &roomsService{
		RoomsRepo: roomsRepo,
	}
}

func (ss *roomsService) GetAll() ([]models.Room, error) {
	return ss.RoomsRepo.GetAll()
}

func (ss *roomsService) CreateOne(
	centre_uid uuid.UUID,
	name string,
	num_seats int,
	info string,
) (*models.Room, error) {
	var room *models.Room

	uid, err := uuid.NewRandom()
	if err != nil {
		return room, err
	}

	time_now := time.Now()

	room, err = ss.RoomsRepo.CreateOne(
		uid,
		centre_uid,
		name,
		num_seats,
		info,
		time_now, time_now,
	)
	if err != nil {
		return room, err
	}

	return room, nil
}
