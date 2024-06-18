package sqlite

import (
	"context"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
	"gorm.io/gorm"
)

// bulbStorage is a struct that contains a pointer to a gorm.DB instance.
type bulbStorage struct {
	db *gorm.DB
}

// NewBulbStorage is a function that returns a new instance of bulbStorage.
func NewBulbStorage(db *gorm.DB) *bulbStorage {
	return &bulbStorage{db: db}
}

// GetOne is a method that returns an error and a pointer to a Bulb instance.
func (s *bulbStorage) GetOne(id uint) (error, *entities.Bulb) {
	return nil, nil
}

// GetAll is a method that returns a slice of pointers to Bulb instances.
func (s *bulbStorage) GetAll(ctx context.Context) []*entities.Bulb {
	return nil
}

// Create is a method to create a new Bulb in database.
func (s *bulbStorage) Create(bulb *entities.Bulb) (error, *entities.Bulb) {
	return nil, nil
}

// Update is a method to update an existing Bulb in database.
func (s *bulbStorage) Update(bulb *entities.Bulb) (error, *entities.Bulb) {
	return nil, nil
}

// Delete is a method to delete an existing Bulb in database.
func (s *bulbStorage) Delete(bulb *entities.Bulb) error {
	return nil
}
