package models

import "time"

type BuyTicket struct {
	EventID       uint64    `json:"event_id"`
	UserID        uint64    `json:"user_id"`
	Status        string    `json:"status" gorm:"type:enum('available','sold','cancelled');default:'available'"`
	PurchaseDate  time.Time `json:"purchase_date"`
	CancelledDate *time.Time `json:"cancelled_date"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	TotalTicket  uint64    `json:"total_ticket"`
}

type TicketResponse struct {
	UserID   uint64 `json:"user_id"`
	TicketID uint64 `json:"ticket_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	EventName   string `json:"event_name"`
	EventDesc   string `json:"event_desc"`
	EventCategory string `json:"event_category"`
}