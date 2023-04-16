package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type studentCoursesPostgresRepo struct {
	dbpool *pgxpool.Pool
}

func NewStudentCoursesPostgresRepo(dbpool *pgxpool.Pool) *studentCoursesPostgresRepo {
	return &studentCoursesPostgresRepo{
		dbpool: dbpool,
	}
}

func (rr *studentCoursesPostgresRepo) GetAll() ([]models.StudentCourses, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `select a.id, a.student_id, a.course_id, a.payment_amount, a.description, a.created_at, a.updated_at, b.first_name, b.last_name, b.username, b.created_at, b.updated_at, c.name, c.description, c.created_at, c.updated_at from student_courses a
	left join students b on a.student_id  = b.id
	left join courses c on a.course_id = c.id`

	var items []models.StudentCourses

	rows, err := rr.dbpool.Query(ctx, query)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.StudentCourses = models.StudentCourses{}
		err = rows.Scan(
			&item.ID,
			&item.StudentID,
			&item.CourseID,
			&item.PaymentAmount,
			&item.Description,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.Student.FirstName,
			&item.Student.LastName,
			&item.Student.Username,
			&item.Student.UpdatedAt,
			&item.Student.CreatedAt,
			&item.Course.Name,
			&item.Course.Description,
			&item.Course.CreatedAt,
			&item.Course.UpdatedAt,
		)
		if err != nil {
			return items, err
		}

		item.Student.ID = item.StudentID
		item.Course.ID = item.CourseID

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}

func (rr *studentCoursesPostgresRepo) CreateOne(
	uid,
	student_uid,
	course_uid uuid.UUID,
	payment_amount int,
	description string,
	created_at, updated_at time.Time,
) (*models.StudentCourses, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO student_courses (id, student_id, course_id, payment_amount, description, created_at, updated_at)
	VALUES (
		$1, $2, $3, $4, $5, $6, $7
	)`

	var item models.StudentCourses

	_, err := rr.dbpool.Exec(ctx, query,
		uid,
		student_uid,
		course_uid,
		payment_amount,
		description,
		created_at, updated_at,
	)

	if err != nil {
		fmt.Println("err: ", err)
		return &item, err
	}

	item = models.StudentCourses{
		ID:            uid,
		StudentID:     student_uid,
		CourseID:      course_uid,
		PaymentAmount: payment_amount,
		Description:   description,
		CreatedAt:     created_at,
		UpdatedAt:     updated_at,
	}

	return &item, nil

}

func (rr *studentCoursesPostgresRepo) DeleteOneByID(uid uuid.UUID) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := "DELETE FROM student_courses WHERE id = $1"

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

func (rr *studentCoursesPostgresRepo) GetManyByCustomID(custom_id uuid.UUID, column_name string) ([]models.StudentCourses, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := fmt.Sprintf(`SELECT a.id, a.student_id, a.course_id, a.payment_amount, a.description, a.created_at, a.updated_at,
	b.first_name, b.last_name, b.username, b.created_at, b.updated_at,
	c.name, c.description, c.created_at, c.updated_at
	FROM student_courses a
	LEFT JOIN students b on a.student_id  = b.id
	LEFT JOIN courses c on a.course_id = c.id
	WHERE %s = $1`, column_name)

	var items []models.StudentCourses

	rows, err := rr.dbpool.Query(ctx, query,
		custom_id,
	)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.StudentCourses = models.StudentCourses{}
		err = rows.Scan(
			&item.ID,
			&item.StudentID,
			&item.CourseID,
			&item.PaymentAmount,
			&item.Description,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.Student.FirstName,
			&item.Student.LastName,
			&item.Student.Username,
			&item.Student.UpdatedAt,
			&item.Student.CreatedAt,
			&item.Course.Name,
			&item.Course.Description,
			&item.Course.CreatedAt,
			&item.Course.UpdatedAt,
		)
		if err != nil {
			return items, err
		}

		item.Student.ID = item.StudentID
		item.Course.ID = item.CourseID

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil
}
