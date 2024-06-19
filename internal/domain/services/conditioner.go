package services

import (
	"github.com/google/uuid"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/dto"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
)

// ConditionerStorage is an interface that contains methods to interact with the database.
type ConditionerStorage interface {
	GetByUUID(uuid string) (*entities.Conditioner, error)
	GetAll(limit, offset int) ([]*entities.Conditioner, error)
	Create(conditioner *entities.Conditioner) (*entities.Conditioner, error)
	Update(conditioner *entities.Conditioner) (*entities.Conditioner, error)
	Delete(uuid string) error
	DoTask(conditioner *entities.Conditioner, task *entities.Task) error
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
func (s conditionerService) Create(createConditioner *dto.CreateConditioner) (*entities.Conditioner, error) {
	conditioner := &entities.Conditioner{
		UUID:        uuid.NewString(),
		AccessToken: createConditioner.AccessToken,
		Name:        createConditioner.Name,
	}
	return s.storage.Create(conditioner)
}

// GetByUUID is a method that returns a pointer to a Conditioner instance and error.
func (s conditionerService) GetByUUID(uuid string) (*entities.Conditioner, error) {
	return s.storage.GetByUUID(uuid)
}

// GetAll is a method that returns a slice of pointers to Conditioner.
func (s conditionerService) GetAll(limit, offset int) ([]*entities.Conditioner, error) {
	return s.storage.GetAll(limit, offset)
}

// Update is a method to update an existing Conditioner in database using an UpdateConditioner DTO.
func (s conditionerService) Update(updateConditioner *dto.UpdateConditioner) (*entities.Conditioner, error) {
	var conditioner *entities.Conditioner
	conditioner.UUID = updateConditioner.UUID
	conditioner.AccessToken = updateConditioner.AccessToken
	conditioner.Name = updateConditioner.Name
	conditioner.Activated = updateConditioner.Activated
	conditioner.Temperature = updateConditioner.Temperature
	return s.storage.Update(conditioner)
}

// Delete is a method to delete an existing Conditioner in database.
func (s conditionerService) Delete(uuid string) error {
	return s.storage.Delete(uuid)
}

// DoTask is a method to execute a task on a Conditioner.
func (s conditionerService) DoTask(conditioner *entities.Conditioner, task *entities.Task) error {
	return s.storage.DoTask(conditioner, task)
}
