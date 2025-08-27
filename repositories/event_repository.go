package repositories

import (
	"ticketing-go/models"

	"gorm.io/gorm"
)

type EventRepository interface {
	Create(event *models.SubmitEvent) error
	GetEvents() ([]models.EventResponse, error)
	DeleteEvent(id string) error
	CheckerAlreadyDeleted(id string) (bool, error)
	CheckerExist(id string) (bool, error)
	UpdateEvent(event *models.UpdateEvent) error
	FindByID(id uint64) (*models.EventResponse, error)
	UpdateCapacity(tx any, id uint64, capacity int) error
	UpdateStatus(tx any, id uint64, status string) error
}

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepository{db}
}

func (r *eventRepository) Create(event *models.SubmitEvent) error {
	return r.db.Table("events").Create(event).Error
}

func (r *eventRepository) GetEvents() ([]models.EventResponse, error) {
	var events []models.EventResponse
	err := r.db.Raw("SELECT * FROM events WHERE deleted_at IS NULL").Scan(&events).Error
	return events, err
}

func (r *eventRepository) DeleteEvent(id string) error {
	return r.db.Exec("UPDATE events SET deleted_at = NOW() WHERE id = ?", id).Error
}

func (r *eventRepository) CheckerAlreadyDeleted(id string) (bool, error) {
	var count int64
	err := r.db.Raw("SELECT COUNT(*) FROM events WHERE id = ? AND deleted_at IS NOT NULL", id).Scan(&count).Error
	return count > 0, err
}

func (r *eventRepository) CheckerExist(id string) (bool, error) {
	var count int64
	err := r.db.Raw("SELECT COUNT(*) FROM events WHERE id = ?", id).Scan(&count).Error
	return count > 0, err
}

func (r *eventRepository) UpdateEvent(event *models.UpdateEvent) error {
	return r.db.Table("events").Where("id = ?", event.ID).Updates(event).Error
}

func (r *eventRepository) FindByID(id uint64) (*models.EventResponse, error) {
	var event models.EventResponse
	err := r.db.Table("events").Where("id = ?", id).First(&event).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *eventRepository) UpdateCapacity(tx any, id uint64, capacity int) error {
	return tx.(*gorm.DB).Table("events").Where("id = ?", id).Update("capacity", capacity).Error
}

func (r *eventRepository) UpdateStatus(tx any, id uint64, status string) error {
	return tx.(*gorm.DB).Table("events").Where("id = ?", id).Update("status", status).Error
}