package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/models"
	"github.com/magic-of-gnu/crm-local-v8/server/internal/repos"
)

type invoicesService struct {
	InvoiceRepo         repos.InvoicesRepo
	LectureCalendarRepo repos.LectureCalendarRepo
	AttendanceRepo      repos.AttendancesRepo
}

func NewInvoiceService(
	repo repos.InvoicesRepo,
	lecturesCalendarRepo repos.LectureCalendarRepo,
	attendanceRepo repos.AttendancesRepo,
) *invoicesService {
	return &invoicesService{
		InvoiceRepo:         repo,
		LectureCalendarRepo: lecturesCalendarRepo,
		AttendanceRepo:      attendanceRepo,
	}
}

func (ss *invoicesService) GetAll() ([]models.Invoice, error) {

	var response []models.Invoice

	response, err := ss.InvoiceRepo.GetAll()
	if err != nil {
		return response, err
	}

	return response, nil

}

// CreateOneWithUnpaidAttendances sets new invoice_ids and payment_statuses to attendances and create new invoice
// on error it removes all created and updated items
// get valid attendances and set new invoice_id and payment_status
// find valid lectureCalendars --> find valid attendances --> updated them by ID
func (ss *invoicesService) CreateOneWithUnpaidAttendances(createItem *models.Invoice) (*models.Invoice, error) {
	var item *models.Invoice

	lecturesCalendar, err := ss.LectureCalendarRepo.GetManyFilteredByCourseIDAndDate(
		createItem.CourseID, createItem.StartDate,
	)
	if err != nil {
		return item, err
	}

	uid, err := uuid.NewRandom()
	if err != nil {
		return item, err
	}

	time_now := time.Now()

	createItem.ID = uid
	createItem.CreatedAt = time_now
	createItem.UpdatedAt = time_now

	itemInvoice, err := ss.InvoiceRepo.CreateOne(createItem)
	if err != nil {
		return itemInvoice, err
	}

	var attendancesUpdated []uuid.UUID

	for _, lecturesCalendarItem := range lecturesCalendar {
		attendanceToUpdate, err :=
			ss.AttendanceRepo.GetOneFilteredByStudentIDAndLectureCalendarID(
				lecturesCalendarItem.ID, createItem.StudentID)

		if err != nil {
			ss.InvoiceRepo.DeleteOneByID(itemInvoice.ID)
			return itemInvoice, err
		}

		attendanceToUpdate.InvoiceID = createItem.ID
		attendanceToUpdate.PaymentStatusID = createItem.PaymentStatusID
		attendanceToUpdate.UpdatedAt = time_now

		attendanceUpdated, err := ss.AttendanceRepo.UpdateOnePaymentStatuses(attendanceToUpdate)
		if err != nil {
			ss.InvoiceRepo.DeleteOneByID(itemInvoice.ID)
			// revert attendance repos
			// ss.AttendanceRepo.DeleteManyByID(attendancesUpdated)
			return itemInvoice, err
		}

		attendancesUpdated = append(attendancesUpdated, attendanceUpdated.ID)

	}

	return itemInvoice, nil
}

func (ss *invoicesService) CreateOne(createItem *models.Invoice) (*models.Invoice, error) {

	// course_id, student_id, start_date, price, payment_status_id, lectures_number, course_id, payment_status_id, start_date

	// 	with course_students as (
	// 	select id from student_courses
	// 	where course_id = '8d287eea-038b-4610-a7db-7b00f7b03fea'
	// )
	// , created_invoices as (
	// 	insert into invoices (id, course_id, student_id, start_date, price, payment_status_id, created_at, updated_at, lectures_number)
	// 	values ('f0e72348-bf37-4c46-bf03-167881020423', '8d287eea-038b-4610-a7db-7b00f7b03fea', 'd9a657d3-7da7-4fbe-88af-6bdcd5d2cb95', '2023-02-01 00:00:00.000 +0600', 40000, 'bc8c944e-bcd7-4f7c-88fc-ee41733470e3',  now(), now(), 2)
	// 	returning id
	// )
	// , lectures_calendar_ids as (
	// 	select id, date
	// 	from lectures_calendar
	// 	where course_id = '8d287eea-038b-4610-a7db-7b00f7b03fea' and lectures_calendar."date"  > '2023-02-01 00:00:00.000 +0600'
	// 	order by lectures_calendar.date asc
	// )
	// , attendance_ids_to_update as (
	// select a.id from attendances a, lectures_calendar_ids lci, course_students cs
	// where a.lectures_calendar_id = lci.id and a.student_id = 'd9a657d3-7da7-4fbe-88af-6bdcd5d2cb95'
	// order by lci.date asc
	// limit 2
	// )
	// update attendances as a
	// set invoice_id = ci.id
	// from attendance_ids_to_update as ai, created_invoices as ci
	// where a.id = ai.id

	var item *models.Invoice

	uid, err := uuid.NewRandom()
	if err != nil {
		return item, err
	}

	time_now := time.Now()

	createItem.ID = uid
	createItem.CreatedAt = time_now
	createItem.UpdatedAt = time_now

	item, err = ss.InvoiceRepo.CreateOne(createItem)
	if err != nil {
		return item, err
	}

	return item, nil
}

func (ss *invoicesService) DeleteOneByID(uid uuid.UUID) error {

	err := ss.InvoiceRepo.DeleteOneByID(uid)
	if err != nil {
		return err
	}

	return nil
}

func (ss *invoicesService) UpdateOneByID(updateItem *models.Invoice) error {

	err := ss.InvoiceRepo.UpdateOneByID(updateItem)
	if err != nil {
		return err
	}

	return nil
}
