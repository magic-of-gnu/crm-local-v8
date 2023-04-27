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

type PaymentStatusesPostgresRepo struct {
	dbpool *pgxpool.Pool
}

func NewPaymentStatusesPostgresRepo(dbpool *pgxpool.Pool) *PaymentStatusesPostgresRepo {
	return &PaymentStatusesPostgresRepo{
		dbpool: dbpool,
	}
}

func (rr *PaymentStatusesPostgresRepo) GetAll() ([]models.PaymentStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `select id, name, description, created_at, updated_at from payment_statuses`

	var results []models.PaymentStatus

	err := pgxscan.Select(ctx, rr.dbpool, &results, query)
	if err != nil {
		return results, err
	}

	return results, nil
}

func (rr *PaymentStatusesPostgresRepo) GetOneByID(uid uuid.UUID) (*models.PaymentStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `select id, name, description, created_at, updated_at from payment_statuses
	where id = $1`

	var result models.PaymentStatus

	row := rr.dbpool.QueryRow(ctx, query, uid)
	err := row.Scan(
		&result.ID,
		&result.Name,
		&result.Description,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		return &result, err
	}

	return &result, nil
}

func (rr *PaymentStatusesPostgresRepo) CreateOne(item *models.PaymentStatus) (*models.PaymentStatus, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO payment_statuses (id, name, description, created_at, updated_at)
	VALUES (
		$1, $2, $3, $4, $5
	) RETURNING id`

	res, err := rr.dbpool.Exec(ctx, query,
		&item.ID,
		&item.Name,
		&item.Description,
		&item.CreatedAt,
		&item.UpdatedAt,
	)

	if err != nil {
		return item, err
	}

	if res.RowsAffected() != 1 {
		return item, fmt.Errorf("error during writing to db")
	}

	return item, nil
}

func (rr *PaymentStatusesPostgresRepo) DeleteOneByID(uid uuid.UUID) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := "DELETE FROM payment_statuses WHERE id = $1"

	row, err := rr.dbpool.Exec(ctx, query, uid)

	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return fmt.Errorf("data was not deleted")
	}

	return nil

}

func (rr *PaymentStatusesPostgresRepo) UpdateOneByID(createItem *models.PaymentStatus) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	fmt.Println("name: ", createItem.Name)

	query := `UPDATE payment_statuses SET
	name = coalesce($2, name),
	description = coalesce($3, description),
	updated_at = $4
	WHERE id = $1
	`

	// AND
	//   ($2 IS NOT NULL AND $2 IS DISTINCT FROM name OR
	//   $3 IS NOT NULL AND $3 IS DISTINCT FROM description);

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
