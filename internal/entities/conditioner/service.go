package conditioner

import (
	"context"
	"gorm.io/gorm"
)

// Service is an interface for conditioner service.
type Service interface {
	Create(ctx context.Context, conditioner *Conditioner) error
	GetById(ctx context.Context, id uint) (*Conditioner, error)
	GetAll(ctx context.Context) ([]Conditioner, error)
	Update(ctx context.Context, conditioner *Conditioner) error
	Delete(ctx context.Context, id uint) error
}

// service implements conditioner service interface.
type service struct {
	database *gorm.DB
}

// NewService creates a new conditioner service.
func NewService(database *gorm.DB) Service {
	return &service{database: database}
}

// Create creates a new conditioner.
func (s service) Create(ctx context.Context, conditioner *Conditioner) error {
	//TODO implement me
	panic("implement me")
}

// GetById returns a conditioner by the id.
func (s service) GetById(ctx context.Context, id uint) (*Conditioner, error) {
	//TODO implement me
	panic("implement me")
}

// GetAll returns all conditioners.
func (s service) GetAll(ctx context.Context) ([]Conditioner, error) {
	//TODO implement me
	panic("implement me")
}

// Update updates a conditioner.
func (s service) Update(ctx context.Context, conditioner *Conditioner) error {
	//TODO implement me
	panic("implement me")
}

// Delete deletes a conditioner by the id.
func (s service) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}
