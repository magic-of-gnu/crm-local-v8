package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
)

type coursesService struct {
	CoursesRepo repos.CoursesRepo
}

func NewCoursesService(repo repos.CoursesRepo) *coursesService {
	return &coursesService{
		CoursesRepo: repo,
	}
}

func (ss *coursesService) GetAll() ([]models.Course, error) {
	return ss.CoursesRepo.GetAll()
}

func (ss *coursesService) CreateOne(
	name,
	description string,
) (*models.Course, error) {
	var item *models.Course

	uid, err := uuid.NewRandom()
	if err != nil {
		return item, err
	}

	time_now := time.Now()

	item, err = ss.CoursesRepo.CreateOne(
		uid,
		name,
		description,
		time_now, time_now,
	)
	if err != nil {
		return item, err
	}

	return item, nil
}
