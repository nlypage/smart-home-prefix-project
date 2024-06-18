package entities

import "gorm.io/gorm"

// Bulb is a struct that represents a bulb entity.
type Bulb struct {
	gorm.Model
	AccessToken string `json:"access_token,omitempty"`
	Name        string `json:"name,omitempty" json:"name,omitempty"`
	Activated   bool   `json:"activated,omitempty" json:"activated,omitempty"`
	Brightness  int    `json:"brightness,omitempty" json:"brightness,omitempty"`
}
