package services

import (
	"context"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/dto"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
)

// AlarmSystemStorage is an interface that contains methods to interact with the database.
type AlarmSystemStorage interface {
	GetOne(id uint) (error, *entities.AlarmSystem)
	GetAll(ctx context.Context) []*entities.AlarmSystem
	Create(alarmSystem *entities.AlarmSystem) (error, *entities.AlarmSystem)
	Update(alarmSystem *entities.AlarmSystem) (error, *entities.AlarmSystem)
	Delete(alarmSystem *entities.AlarmSystem) error
}

// alarmSystemService is a struct that contains a pointer to an AlarmSystemStorage instance.
type alarmSystemService struct {
	storage AlarmSystemStorage
}

// NewAlarmSystemService is a function that returns a new instance of alarmSystemService.
func NewAlarmSystemService(storage AlarmSystemStorage) *alarmSystemService {
	return &alarmSystemService{storage: storage}
}

// Create is a method to create a new AlarmSystem in database using CreateAlarmSystem DTO.
func (s alarmSystemService) Create(ctx context.Context, createAlarmSystem *dto.CreateAlarmSystem) (*entities.AlarmSystem, error) {
	panic("implement me")
}

// GetById is a method that returns an error and a pointer to an AlarmSystem instance.
func (s alarmSystemService) GetById(ctx context.Context, id uint) (*entities.AlarmSystem, error) {
	//TODO implement me
	panic("implement me")
}

// GetAll is a method that returns a slice of AlarmSystem instances.
func (s alarmSystemService) GetAll(ctx context.Context) ([]entities.AlarmSystem, error) {
	//TODO implement me
	panic("implement me")
}

// Update is a method to update an existing AlarmSystem in database using UpdateAlarmSystem DTO.
func (s alarmSystemService) Update(ctx context.Context, updateAlarmSystem *dto.UpdateAlarmSystem) error {
	//TODO implement me
	panic("implement me")
}

// Delete is a method to delete an existing AlarmSystem in database.
func (s alarmSystemService) Delete(ctx context.Context, id uint) error {
	//TODO implement me
	panic("implement me")
}
