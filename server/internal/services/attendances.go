package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
)

type attendancesService struct {
	AttendancesRepo repos.AttendancesRepo
}

func NewAttendancesServiceService(repo repos.AttendancesRepo) *attendancesService {
	return &attendancesService{
		AttendancesRepo: repo,
	}
}

func (ss *attendancesService) GetAll() ([]models.AttendanceResponse, error) {

	var response []models.AttendanceResponse

	repo_response, err := ss.AttendancesRepo.GetAll()
	if err != nil {
		return response, err
	}

	for _, item := range repo_response {
		response = append(response, models.AttendanceResponse{
			ID:                    item.ID,
			LecturesCalendarID:    item.LecturesCalendarID,
			StudentID:             item.StudentID,
			AttendanceValueID:     item.AttendanceValueID,
			Description:           item.Description,
			CreatedAt:             item.CreatedAt,
			UpdatedAt:             item.UpdatedAt,
			AttendanceValue:       item.AttendanceValues.Value,
			AttendanceName:        item.AttendanceValues.Name,
			AttendanceDescription: item.AttendanceValues.Description,
			StudentFirstName:      item.Student.FirstName,
			StudentLastName:       item.Student.LastName,
			StudentUsername:       item.Student.Username,
			RoomID:                item.LectureCalendar.RoomID,
			CourseID:              item.LectureCalendar.CourseID,
			EmployeeID:            item.LectureCalendar.EmployeeID,
			LectureDate:           item.LectureCalendar.Date.Unix(),
			LectureDuration:       int(item.LectureCalendar.Duration) / 10e9,
			RoomName:              item.LectureCalendar.Room.Name,
			CourseName:            item.LectureCalendar.Course.Name,
			EmployeeFirstName:     item.LectureCalendar.Employee.FirstName,
			EmployeeLastName:      item.LectureCalendar.Employee.LastName,
			EmployeeUsername:      item.LectureCalendar.Employee.Username,
		})
	}

	return response, nil

}

func (ss *attendancesService) CreateOne(
	lecture_calendar_id,
	student_id,
	attendance_value_id uuid.UUID,
	description string,
) (*models.Attendance, error) {
	var item *models.Attendance

	uid, err := uuid.NewRandom()
	if err != nil {
		return item, err
	}

	time_now := time.Now()

	item, err = ss.AttendancesRepo.CreateOne(
		uid,
		lecture_calendar_id,
		student_id,
		attendance_value_id,
		description,
		time_now, time_now,
	)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (ss *attendancesService) DeleteOneByID(uid uuid.UUID) error {

	err := ss.AttendancesRepo.DeleteOneByID(uid)
	if err != nil {
		return err
	}

	return nil
}
