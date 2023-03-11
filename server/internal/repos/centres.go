package repo

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
