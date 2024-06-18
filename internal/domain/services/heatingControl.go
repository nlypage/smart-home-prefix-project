package services

import (
	"context"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/dto"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
)

// HeatingControlStorage is an interface that contains methods to interact with the database.
type HeatingControlStorage interface {
	GetOne(id uint) (error, *entities.HeatingControl)
	GetAll(ctx context.Context) []*entities.HeatingControl
	Create(heatingControl *entities.HeatingControl) (error, *entities.HeatingControl)
	Update(bulb *entities.HeatingControl) (error, *entities.HeatingControl)
	Delete(bulb *entities.HeatingControl) error
}

// heatingControlService is a struct that contains a pointer to a HeatingControlStorage instance.
type heatingControlService struct {
	storage HeatingControlStorage
}

// NewHeatingControlService is a function that returns a new instance of heatingControlService.
func NewHeatingControlService(storage HeatingControlStorage) *heatingControlService {
	return &heatingControlService{storage: storage}
}

// Create is a method to create a new HeatingControl in database using a CreateHeatingControl DTO.
func (s heatingControlService) Create(ctx context.Context, createHeatingControl *dto.CreateHeatingControl) (*entities.HeatingControl, error) {
	panic("implement me")
}

// GetById is a method that returns an error and a pointer to a HeatingControl instance.
func (s heatingControlService) GetById(ctx context.Context, id uint) (*entities.HeatingControl, error) {
	//TODO implement me
	panic("implement me")
}

// GetAll is a method that returns a slice of HeatingControl instances.
func (s heatingControlService) GetAll(ctx context.Context) ([]entities.HeatingControl, error) {
	//TODO implement me
	panic("implement me")
}

// Update is a method to update an existing HeatingControl in database using an UpdateHeatingControl DTO.
func (s heatingControlService) Update(ctx context.Context, updateHeatingControl *dto.UpdateHeatingControl) error {
	//TODO implement me
	panic("implement me")
}

// Delete is a method to delete an existing HeatingControl in database.
func (s heatingControlService) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}
