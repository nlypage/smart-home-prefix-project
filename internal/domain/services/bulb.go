package services

import (
	"github.com/google/uuid"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/dto"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
)

// BulbStorage is an interface that contains methods to interact with the database.
type BulbStorage interface {
	GetByUUID(uuid string) (*entities.Bulb, error)
	GetAll(limit, offset int) ([]*entities.Bulb, error)
	Create(bulb *entities.Bulb) (*entities.Bulb, error)
	Update(bulb *entities.Bulb) (*entities.Bulb, error)
	Delete(uuid string) error
	DoTask(bulb *entities.Bulb, task *entities.Task) error
}

// bulbService is a struct that contains a pointer to a BulbStorage instance.
type bulbService struct {
	storage BulbStorage
}

// NewBulbService is a function that returns a new instance of bulbService.
func NewBulbService(storage BulbStorage) *bulbService {
	return &bulbService{storage: storage}
}

// Create is a method to create a new Bulb in database using a CreateBulb DTO.
func (s bulbService) Create(createBulb *dto.CreateBulb) (*entities.Bulb, error) {
	bulb := &entities.Bulb{
		UUID:        uuid.NewString(),
		AccessToken: createBulb.AccessToken,
		Name:        createBulb.Name,
	}
	return s.storage.Create(bulb)
}

// GetByUUID is a method that returns a pointer to a Bulb instance and error.
func (s bulbService) GetByUUID(uuid string) (*entities.Bulb, error) {
	return s.storage.GetByUUID(uuid)
}

// GetAll is a method that returns a slice of pointers to a Bulb and error.
func (s bulbService) GetAll(limit, offset int) ([]*entities.Bulb, error) {
	return s.storage.GetAll(limit, offset)
}

// Update is a method to update an existing Bulb in database using an UpdateBulb DTO.
func (s bulbService) Update(updateBulb *dto.UpdateBulb) (*entities.Bulb, error) {
	var bulb *entities.Bulb
	bulb.UUID = updateBulb.UUID
	bulb.AccessToken = updateBulb.AccessToken
	bulb.Name = updateBulb.Name
	bulb.Activated = updateBulb.Activated
	bulb.Brightness = updateBulb.Brightness
	return s.storage.Update(bulb)
}

// Delete is a method to delete an existing Bulb in database.
func (s bulbService) Delete(uuid string) error {
	return s.storage.Delete(uuid)
}

// DoTask is a method to perform a task on a Bulb.
func (s bulbService) DoTask(bulb *entities.Bulb, task *entities.Task) error {
	return s.storage.DoTask(bulb, task)
}
