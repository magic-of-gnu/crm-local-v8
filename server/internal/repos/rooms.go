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
