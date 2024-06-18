package entities

import (
	"time"
)

// AlarmSystem is a struct that represents an alarm system entity.
type AlarmSystem struct {
	UUID      string    `json:"uuid" gorm:"primaryKey,unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	AccessToken string `json:"access_token" gorm:"unique,not null"`
	Name        string `json:"name" gorm:"unique,not null"`
	Activated   bool   `json:"activated"`
	Triggered   bool   `json:"triggered"`
}

// GetName is a method simply to define the TaskObject interface in task.go.
func (a AlarmSystem) GetName() string {
	return a.Name
}
