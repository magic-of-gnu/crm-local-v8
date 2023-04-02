package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
)

type attendanceValuesService struct {
	AttendanceValuesRepo repos.AttendanceValues
}

func NewAttendanceValuesService(repo repos.AttendanceValues) *attendanceValuesService {
	return &attendanceValuesService{
		AttendanceValuesRepo: repo,
	}
}

func (ss *attendanceValuesService) GetAll() ([]models.AttendanceValues, error) {
	return ss.AttendanceValuesRepo.GetAll()
}

func (ss *attendanceValuesService) CreateOne(
	value int,
	name,
	description string,
) (*models.AttendanceValues, error) {
	var item *models.AttendanceValues

	uid, err := uuid.NewRandom()
	if err != nil {
		return item, err
	}

	time_now := time.Now()

	item, err = ss.AttendanceValuesRepo.CreateOne(
		uid,
		value,
		name,
		description,
		time_now, time_now,
	)
	if err != nil {
		return item, err
	}

	return item, nil
}
