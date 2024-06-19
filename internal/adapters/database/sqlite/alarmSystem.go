package sqlite

import (
	"github.com/nlypage/smart-home-prefix-project/internal/domain/common/actions"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/common/errroz"
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

// Create is a method to create a new AlarmSystem in database.
func (s *alarmSystemStorage) Create(alarmSystem *entities.AlarmSystem) (*entities.AlarmSystem, error) {
	err := s.db.Create(&alarmSystem).Error
	return alarmSystem, err
}

// GetByUUID is a method that returns an error and a pointer to an AlarmSystem instance.
func (s *alarmSystemStorage) GetByUUID(uuid string) (*entities.AlarmSystem, error) {
	var alarmSystem *entities.AlarmSystem
	err := s.db.Model(&entities.AlarmSystem{}).Where("uuid = ?", uuid).First(&alarmSystem).Error
	return alarmSystem, err
}

// GetAll is a method that returns a slice of pointers to AlarmSystem instances.
func (s *alarmSystemStorage) GetAll(limit, offset int) ([]*entities.AlarmSystem, error) {
	var alarmSystems []*entities.AlarmSystem
	err := s.db.Model(&entities.AlarmSystem{}).Limit(limit).Offset(offset).Find(&alarmSystems).Error
	return alarmSystems, err
}

// Update is a method to update an existing AlarmSystem in database.
func (s *alarmSystemStorage) Update(alarmSystem *entities.AlarmSystem) (*entities.AlarmSystem, error) {
	err := s.db.Model(&entities.AlarmSystem{}).Where("uuid = ?", alarmSystem.UUID).Updates(&alarmSystem).Error
	return alarmSystem, err
}

// Delete is a method to delete an existing AlarmSystem in database.
func (s *alarmSystemStorage) Delete(uuid string) error {
	return s.db.Unscoped().Delete(&entities.AlarmSystem{}, "uuid = ?", uuid).Error
}

// DoTask is a method to perform a task on an AlarmSystem.
func (s *alarmSystemStorage) DoTask(alarmSystem *entities.AlarmSystem, task *entities.Task) error {
	if task.Active {
		switch task.Action {
		case actions.Activate:
			alarmSystem.Activated = true
		case actions.Disable:
			alarmSystem.Activated = false
		default:
			return errroz.WrongTaskAction
		}
		return s.db.Model(&entities.AlarmSystem{}).Where("uuid = ?", alarmSystem.UUID).Updates(&alarmSystem).Error
	}
	return errroz.TaskNotActive
}
