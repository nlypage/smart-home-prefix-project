package entities

import (
	"time"
)

// Action is a type that represents a task action.
type Action string

type TaskObject interface {
	GetName() string
}

func GetTaskObjectStruct(taskObject TaskObject) interface{} {
	switch v := taskObject.(type) {
	case *AlarmSystem:
		return v
	case *Bulb:
		return v
	case *Conditioner:
		return v
	case *HeatingControl:
		return v
	default:
		return nil
	}
}

type Task struct {
	UUID      string    `json:"uuid" gorm:"primaryKey,unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	TaskObject TaskObject `json:"task_object" gorm:"polymorphic:TaskObject;"`
	Time       time.Time  `json:"time"`
	DaysOfWeek []int      `json:"days_of_week"`
	Action     Action     `json:"action" `
	Active     bool       `json:"active"`
}
