package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type lectureCalendarPostgresRepo struct {
	dbpool *pgxpool.Pool
}

func NewLectureCalendarPostgresRepo(dbpool *pgxpool.Pool) *lectureCalendarPostgresRepo {
	return &lectureCalendarPostgresRepo{
		dbpool: dbpool,
	}
}

func (rr *lectureCalendarPostgresRepo) GetAll() ([]models.LectureCalendar, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `select a.id, a.room_id, a.course_id, a.employee_id, a.date, a.duration, a.created_at, a.updated_at, 
	   b.centre_id, b.name, b.num_seats, b.info, b.created_at, b.updated_at, 
	   c.name, c.description, c.created_at, c.updated_at,
	   d.first_name, d.last_name, d.username, d.info, d.created_at, d.updated_at
	   from lectures_calendar a
left join rooms b on a.room_id  = b.id
left join courses c on a.course_id = c.id
left join employees d on a.employee_id = d.id `

	var items []models.LectureCalendar

	rows, err := rr.dbpool.Query(ctx, query)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.LectureCalendar
		err = rows.Scan(
			&item.ID,
			&item.RoomID,
			&item.CourseID,
			&item.EmployeeID,
			&item.Date,
			&item.Duration,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.Room.CentreID,
			&item.Room.Name,
			&item.Room.NumSeats,
			&item.Room.Info,
			&item.Room.CreatedAt,
			&item.Room.UpdatedAt,
			&item.Course.Name,
			&item.Course.Description,
			&item.Course.CreatedAt,
			&item.Course.UpdatedAt,
			&item.Employee.FirstName,
			&item.Employee.LastName,
			&item.Employee.Username,
			&item.Employee.Info,
			&item.Employee.CreatedAt,
			&item.Employee.UpdatedAt,
		)

		if err != nil {
			return items, err
		}

		item.Room.ID = item.RoomID
		item.Course.ID = item.CourseID
		item.Employee.ID = item.EmployeeID

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}

func (rr *lectureCalendarPostgresRepo) CreateOne(
	uid,
	roomd_uid,
	course_uid,
	employee_uid uuid.UUID,
	date time.Time,
	duration time.Duration,
	created_at, updated_at time.Time,
) (*models.LectureCalendar, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO lectures_calendar (id, room_id, course_id, employee_id, date, duration, created_at, updated_at)
	VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8
	)`

	var item models.LectureCalendar

	_, err := rr.dbpool.Exec(ctx, query,
		uid,
		roomd_uid,
		course_uid,
		employee_uid,
		date,
		duration,
		created_at, updated_at,
	)

	if err != nil {
		return &item, err
	}

	item = models.LectureCalendar{
		ID:         uid,
		RoomID:     roomd_uid,
		CourseID:   course_uid,
		EmployeeID: employee_uid,
		Date:       date,
		Duration:   duration,
		CreatedAt:  created_at,
		UpdatedAt:  updated_at,
	}

	return &item, nil

}

func (rr *lectureCalendarPostgresRepo) DeleteOneByID(uid uuid.UUID) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `DELETE FROM lectures_calendar WHERE id = $1`

	row, err := rr.dbpool.Exec(ctx, query, uid)

	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return fmt.Errorf("data was not deleted")
	}

	return nil

}

func (rr *lectureCalendarPostgresRepo) GetManyByCourseID(course_id uuid.UUID) ([]models.LectureCalendar, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var items []models.LectureCalendar

	query := `select a.id, a.room_id, a.course_id, a.employee_id, a.date, a.duration, a.created_at, a.updated_at, 
	   b.centre_id, b.name, b.num_seats, b.info, b.created_at, b.updated_at, 
	   c.name, c.description, c.created_at, c.updated_at,
	   d.first_name, d.last_name, d.username, d.info, d.created_at, d.updated_at
	   from lectures_calendar a
left join rooms b on a.room_id  = b.id
left join courses c on a.course_id = c.id
left join employees d on a.employee_id = d.id
WHERE course_id = $1`

	rows, err := rr.dbpool.Query(ctx, query, course_id)
	defer rows.Close()

	if err != nil {
		return items, err
	}

	for rows.Next() {
		var item models.LectureCalendar
		err = rows.Scan(
			&item.ID,
			&item.RoomID,
			&item.CourseID,
			&item.EmployeeID,
			&item.Date,
			&item.Duration,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.Room.CentreID,
			&item.Room.Name,
			&item.Room.NumSeats,
			&item.Room.Info,
			&item.Room.CreatedAt,
			&item.Room.UpdatedAt,
			&item.Course.Name,
			&item.Course.Description,
			&item.Course.CreatedAt,
			&item.Course.UpdatedAt,
			&item.Employee.FirstName,
			&item.Employee.LastName,
			&item.Employee.Username,
			&item.Employee.Info,
			&item.Employee.CreatedAt,
			&item.Employee.UpdatedAt,
		)

		if err != nil {
			return items, err
		}

		item.Room.ID = item.RoomID
		item.Course.ID = item.CourseID
		item.Employee.ID = item.EmployeeID

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil

}

func (rr *lectureCalendarPostgresRepo) GetOneByID(uid uuid.UUID) (*models.LectureCalendar, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var item models.LectureCalendar

	query := `select a.id, a.room_id, a.course_id, a.employee_id, a.date, a.duration, a.created_at, a.updated_at, 
	   b.centre_id, b.name, b.num_seats, b.info, b.created_at, b.updated_at, 
	   c.name, c.description, c.created_at, c.updated_at,
	   d.first_name, d.last_name, d.username, d.info, d.created_at, d.updated_at
	   from lectures_calendar a
left join rooms b on a.room_id  = b.id
left join courses c on a.course_id = c.id
left join employees d on a.employee_id = d.id
WHERE a.id = $1`

	row := rr.dbpool.QueryRow(ctx, query, uid)

	err := row.Scan(
		&item.ID,
		&item.RoomID,
		&item.CourseID,
		&item.EmployeeID,
		&item.Date,
		&item.Duration,
		&item.CreatedAt,
		&item.UpdatedAt,
		&item.Room.CentreID,
		&item.Room.Name,
		&item.Room.NumSeats,
		&item.Room.Info,
		&item.Room.CreatedAt,
		&item.Room.UpdatedAt,
		&item.Course.Name,
		&item.Course.Description,
		&item.Course.CreatedAt,
		&item.Course.UpdatedAt,
		&item.Employee.FirstName,
		&item.Employee.LastName,
		&item.Employee.Username,
		&item.Employee.Info,
		&item.Employee.CreatedAt,
		&item.Employee.UpdatedAt,
	)

	if err != nil {
		return &item, err
	}

	item.Room.ID = item.RoomID
	item.Course.ID = item.CourseID
	item.Employee.ID = item.EmployeeID

	return &item, nil

}

func (rr *lectureCalendarPostgresRepo) GetManyFilteredByCourseIDAndDate(course_id uuid.UUID, start_date time.Time) ([]models.LectureCalendar, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var items []models.LectureCalendar

	query := `SELECT a.id, a.room_id, a.course_id, a.employee_id, a.date, a.duration, a.created_at, a.updated_at, 
	   b.centre_id, b.name, b.num_seats, b.info, b.created_at, b.updated_at, 
	   c.name, c.description, c.created_at, c.updated_at,
	   d.first_name, d.last_name, d.username, d.info, d.created_at, d.updated_at
	   FROM lectures_calendar a
LEFT JOIN rooms b on a.room_id  = b.id
LEFT JOIN courses c on a.course_id = c.id
LEFT JOIN employees d on a.employee_id = d.id
WHERE a.course_id = $1 AND a."date" > $2
ORDER BY a."date" ASC`

	rows, err := rr.dbpool.Query(ctx, query, course_id, start_date)
	if err != nil {
		return items, err
	}

	defer rows.Close()

	for rows.Next() {
		var item models.LectureCalendar
		err = rows.Scan(
			&item.ID,
			&item.RoomID,
			&item.CourseID,
			&item.EmployeeID,
			&item.Date,
			&item.Duration,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.Room.CentreID,
			&item.Room.Name,
			&item.Room.NumSeats,
			&item.Room.Info,
			&item.Room.CreatedAt,
			&item.Room.UpdatedAt,
			&item.Course.Name,
			&item.Course.Description,
			&item.Course.CreatedAt,
			&item.Course.UpdatedAt,
			&item.Employee.FirstName,
			&item.Employee.LastName,
			&item.Employee.Username,
			&item.Employee.Info,
			&item.Employee.CreatedAt,
			&item.Employee.UpdatedAt,
		)

		if err != nil {
			return items, err
		}

		item.Room.ID = item.RoomID
		item.Course.ID = item.CourseID
		item.Employee.ID = item.EmployeeID

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil

}
