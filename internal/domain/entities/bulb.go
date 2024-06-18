package entities

import "time"

// Bulb is a struct that represents a bulb entity.
type Bulb struct {
	UUID      string    `json:"uuid" gorm:"primaryKey,unique"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	AccessToken string `json:"access_token" gorm:"unique,not null"`
	Name        string `json:"name" gorm:"unique,not null"`
	Activated   bool   `json:"activated"`
	Brightness  int    `json:"brightness"`
}

// GetName is a method simply to define the TaskObject interface in task.go.
func (b Bulb) GetName() string {
	return b.Name
}
