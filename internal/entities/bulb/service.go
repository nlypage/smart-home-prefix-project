package bulb

import (
	"context"
	"gorm.io/gorm"
)

// Service is an interface for bulb service.
type Service interface {
	Create(ctx context.Context, bulb *Bulb) error
	GetById(ctx context.Context, id uint) (*Bulb, error)
	GetAll(ctx context.Context) ([]Bulb, error)
	Update(ctx context.Context, bulb *Bulb) error
	Delete(ctx context.Context, id uint) error
}

// service implements bulb service interface.
type service struct {
	database *gorm.DB
}

// NewService creates a new bulb service.
func NewService(database *gorm.DB) Service {
	return &service{database: database}
}

// Create creates a new bulb.
func (s service) Create(ctx context.Context, bulb *Bulb) error {
	//TODO implement me
	panic("implement me")
}

// GetById returns an bulb by the id.
func (s service) GetById(ctx context.Context, id uint) (*Bulb, error) {
	//TODO implement me
	panic("implement me")
}

// GetAll returns all bulbs.
func (s service) GetAll(ctx context.Context) ([]Bulb, error) {
	//TODO implement me
	panic("implement me")
}

// Update updates an bulb.
func (s service) Update(ctx context.Context, bulb *Bulb) error {
	//TODO implement me
	panic("implement me")
}

// Delete deletes a bulb by the id.
func (s service) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}
