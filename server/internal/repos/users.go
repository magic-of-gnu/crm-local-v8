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

type usersPostgresRepo struct {
	dbpool *pgxpool.Pool
}

func NewUsersPostgresRepo(dbpool *pgxpool.Pool) *usersPostgresRepo {
	return &usersPostgresRepo{
		dbpool: dbpool,
	}
}

func (rr *usersPostgresRepo) GetAll() ([]models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT id, first_name, last_name, username, password, last_login, is_admin, user_type, created_at, updated_at FROM users`

	var items []models.User

	err := pgxscan.Select(ctx, rr.dbpool, &items, query)
	if err != nil {
		return items, err
	}

	return items, nil
}

func (rr *usersPostgresRepo) CreateOne(
	uid uuid.UUID,
	first_name,
	last_name,
	username,
	password string,
	last_login time.Time,
	is_admin bool,
	user_type int,
	created_at, updated_at time.Time,
) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO users (
		id,
		first_name,
		last_name,
		username,
		password,
		last_login,
		is_admin,
		user_type,
		created_at, updated_at
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10
	)`

	var item models.User

	res, err := rr.dbpool.Exec(ctx, query,
		uid,
		first_name,
		last_name,
		username,
		password,
		last_login,
		is_admin,
		user_type,
		created_at,
		updated_at,
	)

	if err != nil {
		fmt.Println("err: ", err)
		return &item, err
	}

	if res.RowsAffected() != 1 {
		return &item, fmt.Errorf("error: error occured during write to db")
	}

	item = models.User{
		ID:        uid,
		FirstName: first_name,
		LastName:  last_name,
		Username:  username,
		Password:  password,
		LastLogin: last_login,
		IsAdmin:   is_admin,
		UserType:  user_type,
		UpdatedAt: updated_at,
		CreatedAt: created_at,
	}

	return &item, nil

}

func (rr *usersPostgresRepo) DeleteOneByID(uid uuid.UUID) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := "DELETE FROM users WHERE id = $1"

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
