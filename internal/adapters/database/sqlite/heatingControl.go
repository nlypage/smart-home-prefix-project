package sqlite

import (
	"context"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
	"gorm.io/gorm"
)

// heatingControlStorage is a struct that contains a pointer to a gorm.DB instance.
type heatingControlStorage struct {
	db *gorm.DB
}

// NewHeatingControlStorage is a function that returns a new instance of heatingControlStorage.
func NewHeatingControlStorage(db *gorm.DB) *heatingControlStorage {
	return &heatingControlStorage{db: db}
}

// GetOne is a method that returns an error and a pointer to a HeatingControl instance.
func (s *heatingControlStorage) GetOne(id uint) (error, *entities.HeatingControl) {
	return nil, nil
}

// GetAll is a method that returns a slice of pointers to HeatingControl instances.
func (s *heatingControlStorage) GetAll(ctx context.Context) []*entities.HeatingControl {
	return nil
}

// Create is a method to create a new HeatingControl in database.
func (s *heatingControlStorage) Create(heatingControl *entities.HeatingControl) (error, *entities.HeatingControl) {
	return nil, nil
}

// Update is a method to update an existing HeatingControl in database.
func (s *heatingControlStorage) Update(heatingControl *entities.HeatingControl) (error, *entities.HeatingControl) {
	return nil, nil
}

// Delete is a method to delete an existing HeatingControl in database.
func (s *heatingControlStorage) Delete(heatingControl *entities.HeatingControl) error {
	return nil
}
