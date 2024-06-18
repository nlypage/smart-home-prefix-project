package services

import (
	"context"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/dto"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
)

// BulbStorage is an interface that contains methods to interact with the database.
type BulbStorage interface {
	GetOne(id uint) (error, *entities.Bulb)
	GetAll(ctx context.Context) []*entities.Bulb
	Create(bulb *entities.Bulb) (error, *entities.Bulb)
	Update(bulb *entities.Bulb) (error, *entities.Bulb)
	Delete(bulb *entities.Bulb) error
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
func (s bulbService) Create(ctx context.Context, createBulb *dto.CreateBulb) (*entities.Bulb, error) {
	panic("implement me")
}

// GetById is a method that returns an error and a pointer to a Bulb instance.
func (s bulbService) GetById(ctx context.Context, id uint) (*entities.Bulb, error) {
	//TODO implement me
	panic("implement me")
}

// GetAll is a method that returns a slice of Bulb instances.
func (s bulbService) GetAll(ctx context.Context) ([]entities.Bulb, error) {
	//TODO implement me
	panic("implement me")
}

// Update is a method to update an existing Bulb in database using an UpdateBulb DTO.
func (s bulbService) Update(ctx context.Context, updateBulb *dto.UpdateBulb) error {
	//TODO implement me
	panic("implement me")
}

// Delete is a method to delete an existing Bulb in database.
func (s bulbService) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}
