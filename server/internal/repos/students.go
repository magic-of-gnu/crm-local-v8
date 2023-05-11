package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type StudentsPostgresRepo struct {
	dbpool *pgxpool.Pool
}

func NewStudentssPostgresRepo(dbpool *pgxpool.Pool) *StudentsPostgresRepo {
	return &StudentsPostgresRepo{
		dbpool: dbpool,
	}
}

func (rr *StudentsPostgresRepo) GetAll() ([]models.Student, error) {
	fmt.Println("here")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT id, first_name, last_name, username, created_at, updated_at FROM students`

	var student []models.Student

	err := pgxscan.Select(ctx, rr.dbpool, &student, query)
	if err != nil {
		fmt.Println("error here2")
		return student, err
	}

	return student, nil
}

func (rr *StudentsPostgresRepo) CreateOne(
	uid uuid.UUID,
	first_name,
	last_name,
	username string,
	created_at, updated_at time.Time,
) (*models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO students (id, first_name, last_name, username, created_at, updated_at)
	VALUES (
		$1, $2, $3, $4, $5, $6
	) RETURNING id`

	var student models.Student
	var student_id uuid.NullUUID

	res := rr.dbpool.QueryRow(ctx, query,
		uid,
		first_name,
		last_name,
		username,
		created_at,
		updated_at,
	)

	err := res.Scan(&student_id)
	if err != nil {
		fmt.Println("err: ", err)
		return &student, err
	}

	student = models.Student{
		ID:        uid,
		FirstName: first_name,
		LastName:  last_name,
		Username:  username,
		UpdatedAt: updated_at,
		CreatedAt: created_at,
	}

	return &student, nil

}

func (rr *StudentsPostgresRepo) DeleteOneByID(uid uuid.UUID) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := "DELETE FROM students WHERE id = $1"

	row, err := rr.dbpool.Exec(ctx, query, uid)

	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return fmt.Errorf("data was not deleted")
	}

	return nil

}

func (rr *StudentsPostgresRepo) GetOneByID(uid uuid.UUID) (*models.Student, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT id, first_name, last_name, username, created_at, updated_at FROM students
	where id = $1`

	var result []models.Student
	err := pgxscan.Select(ctx, rr.dbpool, &result, query, uid)

	if err == pgx.ErrNoRows {
		return &models.Student{}, nil
	}

	if err != nil {
		return &result[0], err
	}

	return &result[0], nil
}

func (rr *StudentsPostgresRepo) UpdateOneByID(createItem *models.Student) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `UPDATE students SET
	first_name = $2,
	last_name = $3,
	username = $4,
	updated_at = $5
	WHERE id = $1
	`

	row, err := rr.dbpool.Exec(ctx, query,
		&createItem.ID,
		&createItem.FirstName,
		&createItem.LastName,
		&createItem.Username,
		&createItem.UpdatedAt,
	)

	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return fmt.Errorf("data was not updated")
	}

	return nil

}
