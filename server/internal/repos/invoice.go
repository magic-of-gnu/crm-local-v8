package repos

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
)

type InvoicePostgresRepo struct {
	dbpool *pgxpool.Pool
}

func NewInvoicesPostgresRepo(dbpool *pgxpool.Pool) *InvoicePostgresRepo {
	return &InvoicePostgresRepo{
		dbpool: dbpool,
	}
}

func (rr *InvoicePostgresRepo) GetAll() ([]models.Invoice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `
SELECT a.id, a.course_id, a.student_id, a.start_date, a.price, a.payment_status_id,
	a.lectures_number, a.created_at, a.updated_at,
	b.name,
	c.first_name, c.last_name, c.username,
	d.name, d.description
	FROM invoices a
	LEFT JOIN courses b ON a.course_id = b.id
	LEFT JOIN students c ON a.student_id = c.id
	LEFT JOIN payment_statuses d ON a.payment_status_id = d.id`

	var items []models.Invoice

	rows, err := rr.dbpool.Query(ctx, query)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.Invoice
		err = rows.Scan(
			&item.ID,
			&item.CourseID,
			&item.StudentID,
			&item.StartDate,
			&item.Price,
			&item.PaymentStatusID,
			&item.LectureNumber,
			&item.CreatedAt,
			&item.UpdatedAt,
			&item.Course.Name,
			&item.Student.FirstName,
			&item.Student.LastName,
			&item.Student.Username,
			&item.PaymentStatus.Name,
			&item.PaymentStatus.Description,
		)

		if err != nil {
			return items, err
		}

		item.Course.ID = item.CourseID
		item.Student.ID = item.StudentID
		item.PaymentStatus.ID = item.PaymentStatusID

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return items, err
	}

	return items, nil

}

func (rr *InvoicePostgresRepo) CreateOne(createItem *models.Invoice) (*models.Invoice, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `INSERT INTO invoices (id, course_id, student_id, start_date, price, payment_status_id,
	lectures_number, created_at, updated_at,
	VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9
	) RETURNING id`

	var item models.Invoice

	res := rr.dbpool.QueryRow(ctx, query,
		&createItem.ID,
		&createItem.CourseID,
		&createItem.StudentID,
		&createItem.StartDate,
		&createItem.Price,
		&createItem.PaymentStatusID,
		&createItem.LectureNumber,
		&createItem.CreatedAt, &createItem.UpdatedAt,
	)

	err := res.Scan(&item.ID)
	if err != nil {
		return &item, err
	}

	return createItem, nil

}

func (rr *InvoicePostgresRepo) DeleteOneByID(uid uuid.UUID) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `UPDATE attendances
	SET invoice_id = null
	WHERE invoice_id = $1;
	
	DELETE FROM invoices WHERE id = $1`

	row, err := rr.dbpool.Exec(ctx, query, uid)

	if err != nil {
		return err
	}

	if row.RowsAffected() == 0 {
		return fmt.Errorf("data was not deleted")
	}

	return nil

}

func (rr *InvoicePostgresRepo) UpdateOneByID(createItem *models.Invoice) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := `UPDATE invoices
	SET course_id = $2, student_id = $3, start_date = $4, price = $5, payment_status_id = $6,
	lectures_number = $7, updated_at = $8
	WHERE id = $1;`

	row, err := rr.dbpool.Exec(ctx, query,
		&createItem.ID,
		&createItem.CourseID,
		&createItem.StudentID,
		&createItem.StartDate,
		&createItem.Price,
		&createItem.PaymentStatusID,
		&createItem.LectureNumber,
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
