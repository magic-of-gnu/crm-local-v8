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

type coursesPostgresRepo struct {
	dbpool *pgxpool.Pool
}

func NewCoursesPostgresRepo(dbpool *pgxpool.Pool) *coursesPostgresRepo {
	return &coursesPostgresRepo{
		dbpool: dbpool,
	}
}

func (rr *coursesPostgresRepo) GetAll() ([]models.Course, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT id, name, description, created_at, updated_at FROM courses`

	var item []models.Course

	err := pgxscan.Select(ctx, rr.dbpool, &item, query)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (rr *coursesPostgresRepo) CreateOne(
	uid uuid.UUID,
	name,
	description string,
	created_at, updated_at time.Time,
) (*models.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO courses (id, name, description, created_at, updated_at)
	VALUES (
		$1, $2, $3, $4, $5
	)`

	var item models.Course

	_, err := rr.dbpool.Exec(ctx, query,
		uid,
		name,
		description,
		created_at,
		updated_at,
	)

	if err != nil {
		fmt.Println("err: ", err)
		return &item, err
	}

	item = models.Course{
		ID:          uid,
		Name:        name,
		Description: description,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}

	return &item, nil

}

func (rr *coursesPostgresRepo) DeleteOneByID(uid uuid.UUID) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := "DELETE FROM courses WHERE id = $1"

	row, err := rr.dbpool.Exec(ctx, query, uid)

	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return fmt.Errorf("data was not deleted")
	}

	return nil

}

func (rr *coursesPostgresRepo) GetOneByID(uid uuid.UUID) (*models.Course, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT id, name, description, created_at, updated_at FROM courses
	where id = $1`

	var result []models.Course
	err := pgxscan.Select(ctx, rr.dbpool, &result, query, uid)

	if err == pgx.ErrNoRows {
		return &models.Course{}, nil
	}

	if err != nil {
		return &result[0], err
	}

	return &result[0], nil
}

func (rr *coursesPostgresRepo) UpdateOneByID(createItem *models.Course) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `UPDATE courses SET
	name = $2,
	description = $3,
	updated_at = $4
	WHERE id = $1
	`

	row, err := rr.dbpool.Exec(ctx, query,
		&createItem.ID,
		&createItem.Name,
		&createItem.Description,
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
