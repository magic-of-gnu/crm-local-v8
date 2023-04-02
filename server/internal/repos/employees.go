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

type EmployeesPostgresRepo struct {
	dbpool *pgxpool.Pool
}

func NewEmployeesPostgresRepo(dbpool *pgxpool.Pool) *EmployeesPostgresRepo {
	return &EmployeesPostgresRepo{
		dbpool: dbpool,
	}
}

func (rr *EmployeesPostgresRepo) GetAll() ([]models.Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT id, first_name, last_name, username, info, created_at, updated_at FROM employees`

	var employees []models.Employee

	err := pgxscan.Select(ctx, rr.dbpool, &employees, query)
	if err != nil {
		fmt.Println("error here2")
		return employees, err
	}

	return employees, nil
}

func (rr *EmployeesPostgresRepo) CreateOne(
	uid uuid.UUID,
	first_name,
	last_name,
	username,
	info string,
	created_at, updated_at time.Time,
) (*models.Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO employees (id, first_name, last_name, username, info, created_at, updated_at)
	VALUES (
		$1, $2, $3, $4, $5, $6, $7
	) RETURNING id`

	var employee_uid uuid.UUID
	var employee models.Employee

	res := rr.dbpool.QueryRow(ctx, query,
		uid,
		first_name,
		last_name,
		username,
		info,
		created_at,
		updated_at,
	)

	err := res.Scan(&employee_uid)
	if err != nil {
		fmt.Println("err: ", err)
		return &employee, err
	}

	employee = models.Employee{
		ID:        uid,
		FirstName: first_name,
		LastName:  last_name,
		Username:  username,
		Info:      info,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}

	return &employee, nil

}
