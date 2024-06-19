package entities

import (
	"time"
)

// Action is a type that represents a task action.
type Action string

type TaskObject interface {
	GetName() string
}

type Task struct {
	UUID      string    `json:"uuid" gorm:"primaryKey,unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	TaskObject *TaskObject `json:"task_object" gorm:"polymorphic:TaskObject;"`
	Time       string      `json:"time"`         // "HH:MM"
	DaysOfWeek string      `json:"days_of_week"` // "1,2,3,4,5,6,7"
	Action     Action      `json:"action" `
	Active     bool        `json:"active"`
}
