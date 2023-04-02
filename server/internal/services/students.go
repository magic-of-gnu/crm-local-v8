package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
)

type studentsService struct {
	StudentsRepo repos.StudentsRepo
}

func NewStudentService(repo repos.StudentsRepo) *studentsService {
	return &studentsService{
		StudentsRepo: repo,
	}
}

func (ss *studentsService) GetAll() ([]models.Student, error) {
	return ss.StudentsRepo.GetAll()
}

func (ss *studentsService) CreateOne(
	first_name,
	last_name,
	username string,
) (*models.Student, error) {
	var student *models.Student

	uid, err := uuid.NewRandom()
	if err != nil {
		return student, err
	}

	time_now := time.Now()

	student, err = ss.StudentsRepo.CreateOne(
		uid,
		first_name,
		last_name,
		username,
		time_now, time_now,
	)
	if err != nil {
		return student, err
	}

	return student, nil
}
