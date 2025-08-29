package repositories

import (
	"fmt"
	"ticketing-go/models"

	"gorm.io/gorm"
)

type TicketRepository interface {
	Create(ticket *models.BuyTicket) error
	BeginTx() (any, error)
	CreateTx(tx any, ticket *models.BuyTicket) error
	CommitTx(tx any) error
	RollbackTx(tx any) error
	UpdateStatus(tx any, id uint64, status string) error
	GetAllTicketEvents() ([]models.TicketResponse, error)
	GetAllTicketEventsUserOnly(userId uint64) ([]models.TicketResponse, error)
	GetTicketByID(ticketID uint64) (*models.TicketResponse, error)
}


type ticketRepository struct {
	db *gorm.DB
}

func NewTicketRepository(db *gorm.DB) TicketRepository {
	return &ticketRepository{
		db: db,
	}
}
func (r *ticketRepository) Create(ticket *models.BuyTicket) error {
	return r.db.Table("tickets").Create(ticket).Error
}

func (r *ticketRepository) BeginTx() (any, error) {
	return r.db.Begin(), nil
}

func (r *ticketRepository) CreateTx(tx any, ticket *models.BuyTicket) error {
	return tx.(*gorm.DB).Table("tickets").Create(ticket).Error
}

func (r *ticketRepository) CommitTx(tx any) error {
	return tx.(*gorm.DB).Commit().Error
}

func (r *ticketRepository) RollbackTx(tx any) error {
	return tx.(*gorm.DB).Rollback().Error
}

func (r *ticketRepository) UpdateStatus(tx any, id uint64, status string) error {
	return tx.(*gorm.DB).Table("tickets").Where("id = ?", id).Update("status", status).Error
}

func (r *ticketRepository) GetAllTicketEvents() ([]models.TicketResponse, error) {
	var tickets []models.TicketResponse
	err := r.db.Raw(`
		select 
			t.id as ticket_id,
			u.name,
			u.email,
			e.name as event_name,
			e.description as event_desc,
			e.category as event_category
		from tickets t
		join events e on e.id = t.event_id
		join users u on u.id = t.user_id
	`).Scan(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *ticketRepository) GetAllTicketEventsUserOnly(userId uint64) ([]models.TicketResponse, error) {
	fmt.Println("Fetching tickets for user ID:", userId) 
	var tickets []models.TicketResponse
	err := r.db.Raw(`
		select 
			t.id as ticket_id,
			u.id as user_id,
			u.name,
			u.email,
			e.name as event_name,
			e.description as event_desc,
			e.category as event_category
		from tickets t
		join events e on e.id = t.event_id
		join users u on u.id = t.user_id
		where t.user_id = ?
	`, userId).Scan(&tickets).Error
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

func (r *ticketRepository) GetTicketByID(ticketID uint64) (*models.TicketResponse, error) {
	var ticket models.TicketResponse
	err := r.db.Raw(
		`select 
			t.id as ticket_id,
			u.id as user_id,
			u.name,
			u.email,
			e.name as event_name,
			e.description as event_desc,
			e.category as event_category
		from tickets t
		join events e on e.id = t.event_id
		join users u on u.id = t.user_id
		where t.id = ?`, ticketID).Scan(&ticket).Error
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}