package services

import (
	"context"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/dto"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
)

// ConditionerStorage is an interface that contains methods to interact with the database.
type ConditionerStorage interface {
	GetOne(id uint) (error, *entities.Conditioner)
	GetAll(ctx context.Context) []*entities.Conditioner
	Create(conditioner *entities.Conditioner) (error, *entities.Conditioner)
	Update(conditioner *entities.Conditioner) (error, *entities.Conditioner)
	Delete(conditioner *entities.Conditioner) error
}

// conditionerService is a struct that contains a pointer to a ConditionerStorage instance.
type conditionerService struct {
	storage ConditionerStorage
}

// NewConditionerService is a function that returns a new instance of conditionerService.
func NewConditionerService(storage ConditionerStorage) *conditionerService {
	return &conditionerService{storage: storage}
}

// Create is a method to create a new Conditioner in database using a CreateConditioner DTO.
func (s conditionerService) Create(ctx context.Context, createConditioner *dto.CreateConditioner) (*entities.Bulb, error) {
	panic("implement me")
}

// GetById is a method that returns an error and a pointer to a Conditioner instance.
func (s conditionerService) GetById(ctx context.Context, id uint) (*entities.Conditioner, error) {
	//TODO implement me
	panic("implement me")
}

// GetAll is a method that returns a slice of Conditioner instances.
func (s conditionerService) GetAll(ctx context.Context) ([]entities.Conditioner, error) {
	//TODO implement me
	panic("implement me")
}

// Update is a method to update an existing Conditioner in database using an UpdateConditioner DTO.
func (s conditionerService) Update(ctx context.Context, updateConditioner *dto.UpdateConditioner) error {
	//TODO implement me
	panic("implement me")
}

// Delete is a method to delete an existing Conditioner in database.
func (s conditionerService) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}
