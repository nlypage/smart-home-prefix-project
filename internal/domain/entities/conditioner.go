package entities

import "gorm.io/gorm"

// Conditioner is a struct that represents a conditioner entity.
type Conditioner struct {
	gorm.Model
	AccessToken string `json:"access_token,omitempty"`
	Name        string `json:"name,omitempty"`
	Activated   bool   `json:"activated,omitempty"`
	Temperature int    `json:"temperature,omitempty"`
}
