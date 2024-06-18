package entities

import "time"

// HeatingControl is a struct that represents a heating controller entity.
type HeatingControl struct {
	UUID      string    `json:"uuid" gorm:"primaryKey,unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	AccessToken string `json:"access_token" gorm:"unique,not null"`
	Name        string `json:"name" gorm:"unique,not null"`
	Activated   bool   `json:"activated"`
	Temperature int    `json:"temperature"`
}

// GetName is a method simply to define the TaskObject interface in task.go.
func (h HeatingControl) GetName() string {
	return h.Name
}
