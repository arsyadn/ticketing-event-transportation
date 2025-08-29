package services

import (
	"ticketing-go/models"
	"ticketing-go/repositories"

	"fmt"
	"time"

	"gorm.io/gorm"
)

type TicketService interface {
	CreateTicket(ticket *models.BuyTicket) error
	GetAllTickets(userId uint64, userRole string) ([]models.TicketResponse, error)
	GetTicketByID(ticketID uint64, userID uint64, userRole string) (*models.TicketResponse, error)
}

type ticketService struct {
	repo     repositories.TicketRepository
	eventRepo repositories.EventRepository
	userRepo  repositories.UserRepository
}

func NewTicketService(repo repositories.TicketRepository, eventRepo repositories.EventRepository, userRepo repositories.UserRepository) TicketService {
	return &ticketService{
		repo:     repo,
		eventRepo: eventRepo,
		userRepo:  userRepo,
	}
}

func (s *ticketService) CreateTicket(ticket *models.BuyTicket) error {
	// validate event_id exists
	eventIDStr := fmt.Sprintf("%d", ticket.EventID)
	eventExists, err := s.eventRepo.CheckerExist(eventIDStr)
	if err != nil || !eventExists {
		return fmt.Errorf("event_id %d does not exist", ticket.EventID)
	}

	// validate user_id exists
	userIdExist, err := s.userRepo.FindByID(ticket.UserID)
	if err != nil || userIdExist == nil {
		return fmt.Errorf("user_id %d does not exist", ticket.UserID)
	}

	// validate total_payment matches event price * total_tickets
	event, err := s.eventRepo.FindByID(ticket.EventID)
	if err != nil || event == nil {
		return fmt.Errorf("event_id %d not found for validation", ticket.EventID)
	}

	// validate event capacity is sufficient
	if event.Capacity < int(ticket.TotalTicket) {
		return fmt.Errorf("not enough capacity: available %d, requested %d", event.Capacity, ticket.TotalTicket)
	}

	// reduce event capacity
	event.Capacity -= int(ticket.TotalTicket)

	now := time.Now()
	ticket.PurchaseDate = now
	ticket.CreatedAt = now
	ticket.UpdatedAt = now

	// Begin transaction (gorm.DB)
	tx, err := s.repo.BeginTx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %v", err)
	}
	gormTx := tx.(*gorm.DB)

	// Create ticket
	if err := s.repo.CreateTx(gormTx, ticket); err != nil {
		s.repo.RollbackTx(gormTx)
		return fmt.Errorf("failed to create ticket: %v", err)
	}

	// Update event capacity
	if err := s.eventRepo.UpdateCapacity(gormTx, event.ID, event.Capacity); err != nil {
		s.repo.RollbackTx(gormTx)
		return fmt.Errorf("failed to update event capacity: %v", err)
	}

	// If event capacity is 0 after update, set status to 'sold'
	if event.Capacity == 0 {
		if err := s.repo.UpdateStatus(gormTx, event.ID, "sold"); err != nil {
			s.repo.RollbackTx(gormTx)
			return fmt.Errorf("failed to update event status to sold: %v", err)
		}
	}

	// Commit transaction
	if err := s.repo.CommitTx(gormTx); err != nil {
		return fmt.Errorf("failed to commit transaction: %v", err)
	}

	return nil
}

func (s *ticketService) GetAllTickets(userId uint64, userRole string) ([]models.TicketResponse, error) {
	if userRole == "Admin" {
		return s.repo.GetAllTicketEvents()
	}
	return s.repo.GetAllTicketEventsUserOnly(userId)
}

func (s *ticketService) GetTicketByID(ticketID uint64, userId uint64, userRole string) (*models.TicketResponse, error) {
	ticket, err := s.repo.GetTicketByID(ticketID)
	if err != nil {
		return nil, err
	}

	// Check if the user has permission to view the ticket
	if userRole != "Admin" && ticket.UserID != userId {
		return nil, fmt.Errorf("access denied")
	}

	return ticket, nil
}