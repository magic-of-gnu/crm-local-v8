package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type attendancesPostgresRepo struct {
	dbpool *pgxpool.Pool
}

func NewAttendancesPostgresRepo(dbpool *pgxpool.Pool) *attendancesPostgresRepo {
	return &attendancesPostgresRepo{
		dbpool: dbpool,
	}
}

func (rr *attendancesPostgresRepo) GetAll() ([]models.Attendance, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT
	a.id, a.lectures_calendar_id, a.student_id, a.attendance_value_id, a.description, a.created_at, a.updated_at,
	b.value, b."name", b.description,
	c.first_name, c.last_name, c.username,
	d.room_id, d.course_id, d.employee_id, d.date, d.duration,
	e.name,
	f.name,
	g.first_name, g.last_name, g.username
	from attendances a
	left join attendance_values b on a.attendance_value_id = b.id
	left join students c on a.student_id = c.id
	left join lectures_calendar d on a.lectures_calendar_id = d.id
	left join rooms e on d.room_id = e.id
	left join courses f on d.course_id = f.id
	left join employees g on d.employee_id = g.id;`

	var items []models.Attendance

	rows, err := rr.dbpool.Query(ctx, query)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Attendance
		err = rows.Scan(
			&item.ID,
			&item.LecturesCalendarID,
			&item.StudentID,
			&item.AttendanceValueID,
			&item.Description,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.AttendanceValues.Value,
			&item.AttendanceValues.Name,
			&item.AttendanceValues.Description,
			&item.Student.FirstName,
			&item.Student.LastName,
			&item.Student.Username,
			&item.LectureCalendar.RoomID,
			&item.LectureCalendar.CourseID,
			&item.LectureCalendar.EmployeeID,
			&item.LectureCalendar.Date,
			&item.LectureCalendar.Duration,
			&item.LectureCalendar.Room.Name,
			&item.LectureCalendar.Course.Name,
			&item.LectureCalendar.Employee.FirstName,
			&item.LectureCalendar.Employee.LastName,
			&item.LectureCalendar.Employee.Username,
		)

		if err != nil {
			return items, err
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}

func (rr *attendancesPostgresRepo) CreateOne(
	uid,
	lecture_calendar_id,
	student_id,
	attendance_value_id uuid.UUID,
	description string,
	created_at, updated_at time.Time,
) (*models.Attendance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO attendances (
		id, lectures_calendar_id, student_id, attendance_value_id, description,
	created_at, updated_at
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7
	) ON conflict (lectures_calendar_id, student_id) do update set attendance_value_id = excluded.attendance_value_id, description = excluded.description, updated_at = excluded.updated_at;`

	var item models.Attendance

	res, err := rr.dbpool.Exec(ctx, query,
		uid,
		lecture_calendar_id,
		student_id,
		attendance_value_id,
		description,
		created_at, updated_at,
	)

	if err != nil {
		fmt.Println("err: ", err)
		return &item, err
	}

	if res.RowsAffected() == 0 {
		t := "error: db write was not completed"
		fmt.Println("err: ", t)
		return &item, fmt.Errorf("err: %s", t)
	}

	item = models.Attendance{
		ID:                 uid,
		LecturesCalendarID: lecture_calendar_id,
		StudentID:          student_id,
		AttendanceValueID:  attendance_value_id,
		Description:        description,
		CreatedAt:          created_at,
		UpdatedAt:          updated_at,
		LectureCalendar:    models.LectureCalendar{},
		Student:            models.Student{},
		AttendanceValues:   models.AttendanceValues{},
	}

	return &item, nil

}

func (rr *attendancesPostgresRepo) DeleteOneByID(uid uuid.UUID) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := "DELETE FROM attendances WHERE id = $1"

	row, err := rr.dbpool.Exec(ctx, query, uid)

	if err != nil {
		fmt.Println("err: ", err)
		return err
	}

	if row.RowsAffected() == 0 {
		fmt.Println("err: data was not deleted")
		return fmt.Errorf("data was not deleted")
	}

	return nil

}
