package services

import (
	"ticketing-go/models"
	"ticketing-go/repositories"

	"fmt"
)

type EventService interface {
	CreateEvent(event *models.SubmitEvent) error
	GetEvents() ([]models.EventResponse, error)
	DeleteEvent(id string) error
	UpdateEvent(event *models.UpdateEvent) error
}

type eventService struct {
	repo repositories.EventRepository
}

func NewEventService(repo repositories.EventRepository) EventService {
	return &eventService{repo}
}

func (s *eventService) CreateEvent(event *models.SubmitEvent) error {
	if event.Capacity < 0 {
		return fmt.Errorf("The capacity must be more than 0")
	}
	// validation check name
	existingEvents, err := s.repo.GetEvents()
	if err != nil {
		return err
	}
	for _, e := range existingEvents {
		if e.Name == event.Name {
			return fmt.Errorf("Event with the same name already exists")
		}
	}

	return s.repo.Create(event)
}

func (s *eventService) GetEvents() ([]models.EventResponse, error) {
	return s.repo.GetEvents()
}

func (s *eventService) DeleteEvent(id string) error {
	// Check if event exists
	exists, err := s.repo.CheckerExist(id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Event with id %s not found", id)
	}
	// Check if event is already deleted
	exists, err = s.repo.CheckerAlreadyDeleted(id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("Event already deleted")
	}

	return s.repo.DeleteEvent(id)
}

func (s *eventService) UpdateEvent(event *models.UpdateEvent) error {

	if event.Status != "Active" && event.Status != "Ongoing" && event.Status != "Finished" {
		return fmt.Errorf("Status must be either 'Active', 'Ongoing', or 'Finished'")
	}
	// Check if event exists
	exists, err := s.repo.CheckerExist(fmt.Sprintf("%d", event.ID))
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("Event with id %d not found", event.ID)
	}

	// checker id is deleted
	exists, err = s.repo.CheckerAlreadyDeleted(fmt.Sprintf("%d", event.ID))
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("Event with id %d is already deleted", event.ID)
	}

	return s.repo.UpdateEvent(event)
}