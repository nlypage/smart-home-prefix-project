package entities

import (
	"gorm.io/gorm"
)

// AlarmSystem is a struct that represents an alarm system entity.
type AlarmSystem struct {
	gorm.Model
	AccessToken string `json:"access_token,omitempty"`
	Name        string `json:"name,omitempty"`
	Activated   bool   `json:"activated,omitempty"`
	Triggered   bool   `json:"triggered,omitempty"`
}
