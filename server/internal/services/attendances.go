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

func NewAttendancesService(repo repos.AttendancesRepo) *attendancesService {
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
			PaymentStatusID:       item.PaymentStatusID,
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
			PaymentStatusName:     item.PaymentStatus.Name,
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

func (ss *attendancesService) CreateMany(
	attendances []models.AttendanceCreateOneRequest,
) ([]models.Attendance, error) {

	uids := make([]uuid.UUID, len(attendances))
	items := make([]models.Attendance, len(attendances))

	for ii, item := range attendances {

		time_now := time.Now()

		uid, err := uuid.NewRandom()
		if err != nil {
			return items, err
		}

		tmp, err := ss.AttendancesRepo.CreateOne(
			uid,
			item.LecturesCalendarID,
			item.StudentID,
			item.AttendanceValueID,
			item.Description,
			time_now, time_now,
		)
		if err != nil {
			ss.deleteManyFromList(uids)
			return items, err
		}

		items[ii] = *tmp
		uids[ii] = uid

	}

	return items, nil
}

func (ss *attendancesService) deleteManyFromList(uids []uuid.UUID) error {
	for _, uid := range uids {
		err := ss.AttendancesRepo.DeleteOneByID(uid)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ss *attendancesService) DeleteOneByID(uid uuid.UUID) error {

	err := ss.AttendancesRepo.DeleteOneByID(uid)
	if err != nil {
		return err
	}

	return nil
}
