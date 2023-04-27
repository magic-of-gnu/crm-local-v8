package services

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
)

type lectureCalendarService struct {
	LectureCalendarRepo repos.LectureCalendarRepo
}

func NewLectureCalendarService(repo repos.LectureCalendarRepo) *lectureCalendarService {
	return &lectureCalendarService{
		LectureCalendarRepo: repo,
	}
}

func (ss *lectureCalendarService) GetAll() ([]models.LectureCalendarResponse, error) {

	repo_response, err := ss.LectureCalendarRepo.GetAll()
	if err != nil {
		return make([]models.LectureCalendarResponse, 0), err
	}

	response := make([]models.LectureCalendarResponse, len(repo_response))
	for idx, item := range repo_response {
		response[idx] = models.LectureCalendarResponse{
			ID:                item.ID,
			RoomID:            item.RoomID,
			RoomName:          item.Room.Name,
			CourseID:          item.CourseID,
			CourseName:        item.Course.Name,
			EmployeeID:        item.EmployeeID,
			EmployeeFirstName: item.Employee.FirstName,
			EmployeeLastName:  item.Employee.LastName,
			Date:              item.Date.Unix(),
			Duration:          int(item.Duration.Nanoseconds()) / 10e9,
			CreatedAt:         item.CreatedAt,
			UpdatedAt:         item.UpdatedAt,
		}
	}

	return response, nil
}

func (ss *lectureCalendarService) CreateOne(
	room_uid,
	course_uid,
	employee_uid uuid.UUID,
	date time.Time,
	duration time.Duration,
) (*models.LectureCalendar, error) {
	var item *models.LectureCalendar

	uid, err := uuid.NewRandom()
	if err != nil {
		return item, err
	}

	time_now := time.Now()

	item, err = ss.LectureCalendarRepo.CreateOne(
		uid,
		room_uid,
		course_uid,
		employee_uid,
		date,
		duration,
		time_now, time_now,
	)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (ss *lectureCalendarService) CreateMany(
	room_uid,
	course_uid,
	employee_uid uuid.UUID,
	start_date time.Time,
	end_date time.Time,
	days_and_times []models.DayAndTimeRequest,
) ([]models.LectureCalendar, error) {

	week_days := make(map[string]models.DayAndTimeRequest)
	for _, item := range days_and_times {
		week_days[WeekDays[item.Day]] = models.DayAndTimeRequest{
			StartTimeRequest: item.StartTimeRequest,
			Duration:         item.Duration,
		}
	}

	var items []models.LectureCalendar

	for d := start_date; d.After(end_date) == false; d = d.AddDate(0, 0, 1) {
		val, ok := week_days[d.Weekday().String()]

		if !ok {
			continue
		}

		duration_minutes := time.Duration(val.Duration) * time.Minute
		start_time_request := val.StartTimeRequest

		hours_string := start_time_request[:2]
		hours, err := strconv.Atoi(hours_string)
		if err != nil {
			// revert changes
			ss.DeleteManyByID(items)
			return items, err
		}
		minutes, err := strconv.Atoi(start_time_request[2:])
		if err != nil {
			// revert changes
			ss.DeleteManyByID(items)
			return items, err
		}

		sd := d
		sd = sd.Add(time.Duration(hours) * time.Hour)
		sd = sd.Add(time.Duration(minutes) * time.Minute)

		_item, err := ss.CreateOne(
			room_uid,
			course_uid,
			employee_uid,
			sd,
			duration_minutes,
		)

		if err != nil {
			// revert changes
			ss.DeleteManyByID(items)
			return items, err
		}

		items = append(items, *_item)

	}

	return items, nil

}

func (ss *lectureCalendarService) DeleteOneByID(uid uuid.UUID) error {

	err := ss.LectureCalendarRepo.DeleteOneByID(uid)
	if err != nil {
		return err
	}

	return nil
}

func (ss *lectureCalendarService) DeleteManyByID(lectures_calendar []models.LectureCalendar) error {

	deleted_count := 0
	for _, item := range lectures_calendar {
		err := ss.LectureCalendarRepo.DeleteOneByID(item.ID)
		if err == nil {
			deleted_count += 1

		}
	}

	return nil

}

func (ss *lectureCalendarService) GetManyByCourseID(course_id uuid.UUID) ([]models.LectureCalendar, error) {

	response, err := ss.LectureCalendarRepo.GetManyByCourseID(course_id)
	if err != nil {
		return make([]models.LectureCalendar, 0), err
	}

	return response, nil
}

func (ss *lectureCalendarService) GetOneByID(uid uuid.UUID) (*models.LectureCalendar, error) {

	item, err := ss.LectureCalendarRepo.GetOneByID(uid)

	if err == nil {
		return item, nil
	}

	return item, nil

}
