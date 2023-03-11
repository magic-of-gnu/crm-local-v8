package repos

import (
	"context"
	"time"

	"github.com/georgysavva/scany/v2/pgxscan"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type CentresPostgresRepo struct {
	dbpool *pgxpool.Pool
}

func NewCentresPostgresRepo(dbpool *pgxpool.Pool) *CentresPostgresRepo {
	return &CentresPostgresRepo{
		dbpool: dbpool,
	}
}

func (rr *CentresPostgresRepo) GetAll() ([]models.Centre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT id, name, description, created_at, updated_at FROM centres`

	var centres []models.Centre

	err := pgxscan.Select(ctx, rr.dbpool, &centres, query)
	if err != nil {
		return centres, err
	}

	return centres, nil
}

func (rr *CentresPostgresRepo) CreateOne(id int, name, description string, created_at, updated_at time.Time) (*models.Centre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO centres (id, name, description, created_at, updated_at)
	VALUES (
		$1, $2, $3, $4, $5
	) RETURNING id`

	var centre_id int
	var centre models.Centre

	res := rr.dbpool.QueryRow(ctx, query,
		id,
		name,
		description,
		created_at,
		updated_at,
	)

	err := res.Scan(&centre_id)
	if err != nil {
		return &centre, err
	}

	centre = models.Centre{
		ID:          id,
		Name:        name,
		Description: description,
		CreatedAt:   created_at,
		UpdatedAt:   updated_at,
	}

	return &centre, nil

}
