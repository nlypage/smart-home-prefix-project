package dto

// CreateBulb is a struct that represents a DTO to create a new Bulb.
type CreateBulb struct {
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
}

// UpdateBulb is a struct that represents a DTO to update an existing Bulb.
type UpdateBulb struct {
	ID          uint   `json:"id"`
	AccessToken string `json:"access_token"`
	Name        string `json:"name"`
	Activated   bool   `json:"activated"`
	Brightness  int    `json:"brightness"`
}
