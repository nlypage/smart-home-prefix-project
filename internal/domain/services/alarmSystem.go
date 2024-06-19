package services

import (
	"github.com/google/uuid"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/dto"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
)

// AlarmSystemStorage is an interface that contains methods to interact with the database.
type AlarmSystemStorage interface {
	Create(alarmSystem *entities.AlarmSystem) (*entities.AlarmSystem, error)
	GetByUUID(uuid string) (*entities.AlarmSystem, error)
	GetAll(limit, offset int) ([]*entities.AlarmSystem, error)
	Update(alarmSystem *entities.AlarmSystem) (*entities.AlarmSystem, error)
	Delete(uuid string) error
	DoTask(alarmSystem *entities.AlarmSystem, task *entities.Task) error
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
func (s alarmSystemService) Create(createAlarmSystem *dto.CreateAlarmSystem) (*entities.AlarmSystem, error) {
	alarmSystem := &entities.AlarmSystem{
		UUID:        uuid.NewString(),
		AccessToken: createAlarmSystem.AccessToken,
		Name:        createAlarmSystem.Name,
	}
	return s.storage.Create(alarmSystem)
}

// GetByUUID is a method that returns a pointer to an AlarmSystem instance and error.
func (s alarmSystemService) GetByUUID(uuid string) (*entities.AlarmSystem, error) {
	return s.storage.GetByUUID(uuid)
}

// GetAll is a method that returns a slice of pointers to AlarmSystem and error.
func (s alarmSystemService) GetAll(limit, offset int) ([]*entities.AlarmSystem, error) {
	return s.storage.GetAll(limit, offset)
}

// Update is a method to update an existing AlarmSystem in database using UpdateAlarmSystem DTO.
func (s alarmSystemService) Update(updateAlarmSystem *dto.UpdateAlarmSystem) (*entities.AlarmSystem, error) {
	var alarmSystem *entities.AlarmSystem
	alarmSystem.UUID = updateAlarmSystem.UUID
	alarmSystem.AccessToken = updateAlarmSystem.AccessToken
	alarmSystem.Name = updateAlarmSystem.Name
	alarmSystem.Activated = updateAlarmSystem.Activated
	alarmSystem.Triggered = updateAlarmSystem.Triggered
	return s.storage.Update(alarmSystem)
}

// Delete is a method to delete an existing AlarmSystem in database.
func (s alarmSystemService) Delete(uuid string) error {
	return s.storage.Delete(uuid)
}

func (s alarmSystemService) DoTask(alarmSystem *entities.AlarmSystem, task *entities.Task) error {
	return s.storage.DoTask(alarmSystem, task)
}
