package sqlite

import (
	"context"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
	"gorm.io/gorm"
)

// conditionerStorage is a struct that contains a pointer to a gorm.DB instance.
type conditionerStorage struct {
	db *gorm.DB
}

// NewConditionerStorage is a function that returns a new instance of conditionerStorage.
func NewConditionerStorage(db *gorm.DB) *conditionerStorage {
	return &conditionerStorage{db: db}
}

// GetOne is a method that returns an error and a pointer to a Conditioner instance.
func (s *conditionerStorage) GetOne(id uint) (error, *entities.Conditioner) {
	return nil, nil
}

// GetAll is a method that returns a slice of pointers to Conditioner instances.
func (s *conditionerStorage) GetAll(ctx context.Context) []*entities.Conditioner {
	return nil
}

// Create is a method to create a new Conditioner in database.
func (s *conditionerStorage) Create(conditioner *entities.Conditioner) (error, *entities.Conditioner) {
	return nil, nil
}

// Update is a method to update an existing Conditioner in database.
func (s *conditionerStorage) Update(conditioner *entities.Conditioner) (error, *entities.Conditioner) {
	return nil, nil
}

// Delete is a method to delete an existing Conditioner in database.
func (s *conditionerStorage) Delete(conditioner *entities.Conditioner) error {
	return nil
}
