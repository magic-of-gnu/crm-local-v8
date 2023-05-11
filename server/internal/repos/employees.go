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

func (rr *EmployeesPostgresRepo) DeleteOneByID(uid uuid.UUID) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := "DELETE FROM employees WHERE id = $1"

	row, err := rr.dbpool.Exec(ctx, query, uid)

	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return fmt.Errorf("data was not deleted")
	}

	return nil

}

func (rr *EmployeesPostgresRepo) UpdateOneByID(createItem *models.Employee) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `UPDATE employees SET
	first_name = $2,
	last_name = $3,
	username = $4,
	info = $5,
	updated_at = $6
	WHERE id = $1
	`

	row, err := rr.dbpool.Exec(ctx, query,
		&createItem.ID,
		&createItem.FirstName,
		&createItem.LastName,
		&createItem.Username,
		&createItem.Info,
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

func (rr *EmployeesPostgresRepo) GetOneByID(uid uuid.UUID) (*models.Employee, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT id, first_name, last_name, username, info, created_at, updated_at FROM employees
	where id = $1`

	var result models.Employee

	row := rr.dbpool.QueryRow(ctx, query, uid)
	err := row.Scan(
		&result.ID,
		&result.FirstName,
		&result.LastName,
		&result.Username,
		&result.Info,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err == pgx.ErrNoRows {
		return &result, nil
	}

	if err != nil {
		return &result, err
	}

	return &result, nil
}
