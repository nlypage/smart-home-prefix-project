package services

import (
	"github.com/nlypage/smart-home-prefix-project/internal/domain/dto"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
)

// HeatingControlStorage is an interface that contains methods to interact with the database.
type HeatingControlStorage interface {
	GetByUUID(uuid string) (*entities.HeatingControl, error)
	GetAll(limit, offset int) ([]*entities.HeatingControl, error)
	Create(heatingControl *entities.HeatingControl) (*entities.HeatingControl, error)
	Update(bulb *entities.HeatingControl) (*entities.HeatingControl, error)
	Delete(uuid string) error
	DoTask(heatingControl *entities.HeatingControl, task *entities.Task) error
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
func (s heatingControlService) Create(createHeatingControl *dto.CreateHeatingControl) (*entities.HeatingControl, error) {
	heatingControl := &entities.HeatingControl{
		AccessToken: createHeatingControl.AccessToken,
		Name:        createHeatingControl.Name,
	}
	return s.storage.Create(heatingControl)
}

// GetByUUID is a method that returns a pointer to a HeatingControl instance and error.
func (s heatingControlService) GetByUUID(uuid string) (*entities.HeatingControl, error) {
	return s.storage.GetByUUID(uuid)
}

// GetAll is a method that returns a slice of pointers to HeatingControl and error.
func (s heatingControlService) GetAll(limit, offset int) ([]*entities.HeatingControl, error) {
	return s.storage.GetAll(limit, offset)
}

// Update is a method to update an existing HeatingControl in database using an UpdateHeatingControl DTO.
func (s heatingControlService) Update(updateHeatingControl *dto.UpdateHeatingControl) (*entities.HeatingControl, error) {
	var heatingControl *entities.HeatingControl
	heatingControl.UUID = updateHeatingControl.UUID
	heatingControl.AccessToken = updateHeatingControl.AccessToken
	heatingControl.Name = updateHeatingControl.Name
	heatingControl.Activated = updateHeatingControl.Activated
	heatingControl.Temperature = updateHeatingControl.Temperature
	return s.storage.Update(heatingControl)
}

// Delete is a method to delete an existing HeatingControl in database.
func (s heatingControlService) Delete(uuid string) error {
	return s.storage.Delete(uuid)
}

// DoTask is a method to perform a task on an existing HeatingControl in database.
func (s heatingControlService) DoTask(heatingControl *entities.HeatingControl, task *entities.Task) error {
	return s.storage.DoTask(heatingControl, task)
}
