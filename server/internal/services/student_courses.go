package services

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
)

var validCustomColumn map[string]bool = map[string]bool{
	"student_id": true,
	"course_id":  true,
	"id":         true,
}

type studentCoursesService struct {
	StudentCoursesRepo repos.StudentCoursesRepo
}

func NewStudentCoursesService(repo repos.StudentCoursesRepo) *studentCoursesService {
	return &studentCoursesService{
		StudentCoursesRepo: repo,
	}
}

func (ss *studentCoursesService) GetAll() ([]models.StudentCoursesListResponse, error) {

	student_courses, err := ss.StudentCoursesRepo.GetAll()
	if err != nil {
		return make([]models.StudentCoursesListResponse, 0), err
	}

	student_courses_response := make([]models.StudentCoursesListResponse, len(student_courses))
	for idx, item := range student_courses {
		student_courses_response[idx] = models.StudentCoursesListResponse{
			ID:               item.ID,
			StudentID:        item.StudentID,
			CourseID:         item.CourseID,
			PaymentAmount:    item.PaymentAmount,
			Description:      item.Description,
			CreatedAt:        item.CreatedAt,
			UpdatedAt:        item.UpdatedAt,
			StudentFirstName: item.Student.FirstName,
			StudentLastName:  item.Student.LastName,
			CourseName:       item.Course.Name,
		}
	}

	return student_courses_response, nil
}

func (ss *studentCoursesService) CreateOne(
	student_uid,
	course_uid uuid.UUID,
	payment_amount int,
	description string,
) (*models.StudentCourses, error) {
	var item *models.StudentCourses

	uid, err := uuid.NewRandom()
	if err != nil {
		return item, err
	}

	time_now := time.Now()

	item, err = ss.StudentCoursesRepo.CreateOne(
		uid,
		student_uid,
		course_uid,
		payment_amount,
		description,
		time_now, time_now,
	)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (ss *studentCoursesService) DeleteOneByID(uid uuid.UUID) error {
	return ss.StudentCoursesRepo.DeleteOneByID(uid)
}

func (ss *studentCoursesService) GetOneByID(uid uuid.UUID) (*models.StudentCourses, error) {
	return ss.StudentCoursesRepo.GetOneByID(uid)
}

func (ss *studentCoursesService) GetManyByCustomID(uid uuid.UUID, column_name string) ([]models.StudentCourses, error) {

	var items []models.StudentCourses

	_, ok := validCustomColumn[column_name]
	if !ok {
		return items, fmt.Errorf("invalid column name")
	}

	items, err := ss.StudentCoursesRepo.GetManyByCustomID(uid, column_name)
	if err != nil {
		return items, err
	}

	return items, nil
}
