package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
)

type employeesService struct {
	EmployeesRepo repos.EmployeesRepo
}

func NewEmployeesService(employeesRepo repos.EmployeesRepo) *employeesService {
	return &employeesService{
		EmployeesRepo: employeesRepo,
	}
}

func (ss *employeesService) GetAll() ([]models.Employee, error) {
	return ss.EmployeesRepo.GetAll()
}

func (ss *employeesService) CreateOne(
	first_name,
	last_name,
	username,
	info string,
) (*models.Employee, error) {
	var empl *models.Employee

	uid, err := uuid.NewRandom()
	if err != nil {
		return empl, err
	}

	time_now := time.Now()

	empl, err = ss.EmployeesRepo.CreateOne(
		uid,
		first_name,
		last_name,
		username,
		info,
		time_now, time_now,
	)
	if err != nil {
		return empl, err
	}

	return empl, nil
}
