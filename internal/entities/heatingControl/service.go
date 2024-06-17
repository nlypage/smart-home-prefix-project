package heatingControl

import (
	"context"
	"gorm.io/gorm"
)

// Service is an interface for heatingControl service.
type Service interface {
	Create(ctx context.Context, heatingControl *HeatingControl) error
	GetById(ctx context.Context, id uint) (*HeatingControl, error)
	GetAll(ctx context.Context) ([]HeatingControl, error)
	Update(ctx context.Context, heatingControl *HeatingControl) error
	Delete(ctx context.Context, id uint) error
}

// service implements heatingControl service interface.
type service struct {
	database *gorm.DB
}

// NewService creates a new heatingControl service.
func NewService(database *gorm.DB) Service {
	return &service{database: database}
}

// Create creates a new heatingControl.
func (s service) Create(ctx context.Context, heatingControl *HeatingControl) error {
	//TODO implement me
	panic("implement me")
}

// GetById returns an heatingControl by the id.
func (s service) GetById(ctx context.Context, id uint) (*HeatingControl, error) {
	//TODO implement me
	panic("implement me")
}

// GetAll returns all heatingControls.
func (s service) GetAll(ctx context.Context) ([]HeatingControl, error) {
	//TODO implement me
	panic("implement me")
}

// Update updates an heatingControl.
func (s service) Update(ctx context.Context, heatingControl *HeatingControl) error {
	//TODO implement me
	panic("implement me")
}

// Delete deletes an heatingControl by the id.
func (s service) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}
