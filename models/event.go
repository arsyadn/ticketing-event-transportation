package models

import "time"

type SubmitEvent struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Category string `json:"category"`
	Capacity int `json:"capacity"`
	Price float64 `json:"price"`
	Status      string    `json:"status" gorm:"type:enum('Active','Ongoing','Finished');default:'Active'"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedBy   uint64    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}


type EventResponse struct {
	ID			uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Capacity    int       `json:"capacity"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UpdateEvent struct {
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Capacity    int       `json:"capacity"`
	Price       float64   `json:"price"`
	Status      string    `json:"status" gorm:"type:enum('Active','Ongoing','Finished');"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	UpdatedAt   time.Time `json:"updated_at"`
}
