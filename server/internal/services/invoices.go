package services

import (
	"fmt"
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
		createItem.CourseID, createItem.StartDate, 1000,
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

	// Get LectureCalendars from start_date
	var itemInvoice *models.Invoice

	lectures_calendars, err := ss.LectureCalendarRepo.
		GetManyFilteredByCourseIDAndDate(createItem.CourseID, createItem.StartDate, createItem.LecturesNumber)
	if err != nil {
		return itemInvoice, err
	}

	if len(lectures_calendars) == 0 || len(lectures_calendars) > createItem.LecturesNumber {
		return itemInvoice, fmt.Errorf("error: incorrect query")
	}

	// Get respective attendances for LecturesCalendars and StudentID
	var attendances []models.Attendance

	for _, lectures_calendar_item := range lectures_calendars {

		attendance, err := ss.AttendanceRepo.
			GetOneAtendanceWithLecturesCalendarIDAndStudentID(lectures_calendar_item.ID, createItem.StudentID)

		if err != nil {
			return itemInvoice, err
		}

		attendances = append(attendances, *attendance)
	}

	// Insert Invoice
	uid, err := uuid.NewRandom()
	if err != nil {
		return itemInvoice, err
	}

	time_now := time.Now()

	createItem.ID = uid
	createItem.CreatedAt = time_now
	createItem.UpdatedAt = time_now

	itemInvoice, err = ss.InvoiceRepo.CreateOne(createItem)
	if err != nil {
		return itemInvoice, err
	}

	// Update attendances with InvoiceID
	var updatedAttendances []models.Attendance
	var updateErr error = nil

	for _, attendance := range attendances {
		err := ss.AttendanceRepo.UpdateInvoiceIDOnOneAttendance(
			attendance.ID,
			createItem.ID,
			time.Now(),
		)
		if err != nil {
			updateErr = err
			break
		}

		updatedAttendances = append(updatedAttendances, attendance)
	}

	// revert if error
	if updateErr != nil {
		ss.InvoiceRepo.DeleteOneByID(itemInvoice.ID)

		for _, attendance := range updatedAttendances {
			ss.AttendanceRepo.UpdateInvoiceIDOnOneAttendance(
				attendance.ID,
				createItem.ID,
				time.Now(),
			)
		}

	}

	return itemInvoice, nil
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
