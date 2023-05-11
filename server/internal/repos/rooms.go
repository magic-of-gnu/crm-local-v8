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

type RoomsPostgresRepo struct {
	dbpool *pgxpool.Pool
}

func NewRoomsPostgresRepo(dbpool *pgxpool.Pool) *RoomsPostgresRepo {
	return &RoomsPostgresRepo{
		dbpool: dbpool,
	}
}

func (rr *RoomsPostgresRepo) GetAll() ([]models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT id, centre_id, name, num_seats, info, created_at, updated_at FROM rooms`

	var rooms []models.Room

	err := pgxscan.Select(ctx, rr.dbpool, &rooms, query)
	if err != nil {
		fmt.Println("error here2")
		return rooms, err
	}

	return rooms, nil
}

func (rr *RoomsPostgresRepo) CreateOne(
	uid,
	centre_uid uuid.UUID,
	name string,
	num_seats int,
	info string,
	created_at, updated_at time.Time,
) (*models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO rooms (id, centre_id, name, num_seats, info, created_at, updated_at)
	VALUES (
		$1, $2, $3, $4, $5, $6, $7
	) RETURNING id`

	var room_uid uuid.UUID
	var room models.Room

	res := rr.dbpool.QueryRow(ctx, query,
		uid,
		centre_uid,
		name,
		num_seats,
		info,
		created_at,
		updated_at,
	)

	err := res.Scan(&room_uid)
	if err != nil {
		fmt.Println("err: ", err)
		return &room, err
	}

	room = models.Room{
		ID:        uid,
		CentreID:  centre_uid,
		Name:      name,
		NumSeats:  num_seats,
		Info:      info,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}

	return &room, nil

}

func (rr *RoomsPostgresRepo) DeleteOneByID(uid uuid.UUID) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := "DELETE FROM rooms WHERE id = $1"

	row, err := rr.dbpool.Exec(ctx, query, uid)

	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return fmt.Errorf("data was not deleted")
	}

	return nil

}

func (rr *RoomsPostgresRepo) GetOneByID(uid uuid.UUID) (*models.Room, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `SELECT id, centre_id, name, num_seats, info, created_at, updated_at FROM rooms
	where id = $1`

	var result []models.Room
	err := pgxscan.Select(ctx, rr.dbpool, &result, query, uid)

	if err == pgx.ErrNoRows {
		return &models.Room{}, nil
	}

	if err != nil {
		return &result[0], err
	}

	return &result[0], nil
}

func (rr *RoomsPostgresRepo) UpdateOneByID(createItem *models.Room) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `UPDATE employees SET
	centre_id = $2,
	name = $3,
	num_seats = $4,
	info = $5,
	updated_at = $6
	WHERE id = $1
	`

	row, err := rr.dbpool.Exec(ctx, query,
		&createItem.ID,
		&createItem.CentreID,
		&createItem.Name,
		&createItem.NumSeats,
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
