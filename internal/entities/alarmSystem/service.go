package alarmSystem

import (
	"context"
	"gorm.io/gorm"
)

// Service is an interface for alarmSystem service.
type Service interface {
	Create(ctx context.Context, alarmSystem *AlarmSystem) error
	GetById(ctx context.Context, id uint) (*AlarmSystem, error)
	GetAll(ctx context.Context) ([]AlarmSystem, error)
	Update(ctx context.Context, alarmSystem *AlarmSystem) error
	Delete(ctx context.Context, id uint) error
}

// service implements alarmSystem service interface.
type service struct {
	database *gorm.DB
}

// NewService creates a new alarmSystem service.
func NewService(database *gorm.DB) Service {
	return &service{database: database}
}

// Create creates a new alarmSystem.
func (s service) Create(ctx context.Context, alarmSystem *AlarmSystem) error {
	//TODO implement me
	panic("implement me")
}

// GetById returns an alarmSystem by the id.
func (s service) GetById(ctx context.Context, id uint) (*AlarmSystem, error) {
	//TODO implement me
	panic("implement me")
}

// GetAll returns all alarmSystems.
func (s service) GetAll(ctx context.Context) ([]AlarmSystem, error) {
	//TODO implement me
	panic("implement me")
}

// Update updates an alarmSystem.
func (s service) Update(ctx context.Context, alarmSystem *AlarmSystem) error {
	//TODO implement me
	panic("implement me")
}

// Delete deletes an alarmSystem by the id.
func (s service) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}
