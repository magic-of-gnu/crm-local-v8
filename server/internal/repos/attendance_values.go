package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type attendanceValuesPostgresRepo struct {
	dbpool *pgxpool.Pool
}

func NewAttendanceValuesPostgresRepo(dbpool *pgxpool.Pool) *attendanceValuesPostgresRepo {
	return &attendanceValuesPostgresRepo{
		dbpool: dbpool,
	}
}

func (rr *attendanceValuesPostgresRepo) GetAll() ([]models.AttendanceValues, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT id, value, name, description, created_at, updated_at FROM attendance_values`

	var item []models.AttendanceValues

	err := pgxscan.Select(ctx, rr.dbpool, &item, query)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (rr *attendanceValuesPostgresRepo) CreateOne(
	uid uuid.UUID,
	value int,
	name,
	description string,
	created_at, updated_at time.Time,
) (*models.AttendanceValues, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO attendance_values (id, value, name, description, created_at, updated_at)
	VALUES (
		$1, $2, $3, $4, $5, $6
	)`

	var item models.AttendanceValues

	_, err := rr.dbpool.Exec(ctx, query,
		uid,
		value,
		name,
		description,
		created_at,
		updated_at,
	)

	if err != nil {
		fmt.Println("err: ", err)
		return &item, err
	}

	item = models.AttendanceValues{
		ID:          uid,
		Value:       value,
		Name:        name,
		Description: description,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}

	return &item, nil

}
