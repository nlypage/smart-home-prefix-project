package sqlite

import (
	"context"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
	"gorm.io/gorm"
)

// alarmSystemStorage is a struct that contains a pointer to a gorm.DB instance.
type alarmSystemStorage struct {
	db *gorm.DB
}

// NewAlarmSystemStorage is a function that returns a new instance of alarmSystemStorage.
func NewAlarmSystemStorage(db *gorm.DB) *alarmSystemStorage {
	return &alarmSystemStorage{db: db}
}

// GetOne is a method that returns an error and a pointer to an AlarmSystem instance.
func (s *alarmSystemStorage) GetOne(id uint) (error, *entities.AlarmSystem) {
	return nil, nil
}

// GetAll is a method that returns a slice of pointers to AlarmSystem instances.
func (s *alarmSystemStorage) GetAll(ctx context.Context) []*entities.AlarmSystem {
	return nil
}

// Create is a method to create a new AlarmSystem in database.
func (s *alarmSystemStorage) Create(alarmSystem *entities.AlarmSystem) (error, *entities.AlarmSystem) {
	return nil, nil
}

// Update is a method to update an existing AlarmSystem in database.
func (s *alarmSystemStorage) Update(alarmSystem *entities.AlarmSystem) (error, *entities.AlarmSystem) {
	return nil, nil
}

// Delete is a method to delete an existing AlarmSystem in database.
func (s *alarmSystemStorage) Delete(alarmSystem *entities.AlarmSystem) error {
	return nil
}
