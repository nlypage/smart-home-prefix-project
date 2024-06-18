package entities

import "gorm.io/gorm"

// HeatingControl is a struct that represents a heating controller entity.
type HeatingControl struct {
	gorm.Model
	AccessToken string `json:"access_token"`
	Name        string `json:"name,omitempty"`
	Activated   bool   `json:"activated,omitempty"`
	Temperature int    `json:"temperature,omitempty"`
}
