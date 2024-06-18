package sqlite

import (
	"github.com/nlypage/smart-home-prefix-project/internal/domain/common/actions"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/common/errroz"
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

// GetByUUID is a method that returns an error and a pointer to a HeatingControl instance.
func (s *heatingControlStorage) GetByUUID(uuid string) (*entities.HeatingControl, error) {
	var heatingControl *entities.HeatingControl
	err := s.db.Model(&entities.HeatingControl{}).Where("uuid = ?", uuid).First(heatingControl).Error
	return heatingControl, err
}

// GetAll is a method that returns a slice of pointers to HeatingControl instances.
func (s *heatingControlStorage) GetAll(limit, offset int) ([]*entities.HeatingControl, error) {
	var heatingControls []*entities.HeatingControl
	err := s.db.Model(&entities.HeatingControl{}).Limit(limit).Offset(offset).Find(heatingControls).Error
	return heatingControls, err
}

// Create is a method to create a new HeatingControl in database.
func (s *heatingControlStorage) Create(heatingControl *entities.HeatingControl) (*entities.HeatingControl, error) {
	err := s.db.Create(heatingControl).Error
	return heatingControl, err
}

// Update is a method to update an existing HeatingControl in database.
func (s *heatingControlStorage) Update(heatingControl *entities.HeatingControl) (*entities.HeatingControl, error) {
	err := s.db.Model(&entities.HeatingControl{}).Where("uuid = ", heatingControl.UUID).Updates(heatingControl).Error
	return heatingControl, err
}

// Delete is a method to delete an existing HeatingControl in database.
func (s *heatingControlStorage) Delete(uuid string) error {
	err := s.db.Unscoped().Delete(&entities.HeatingControl{}, "uuid = ?", uuid).Error
	return err
}

// DoTask is a method to perform a task on a HeatingControl.
func (s *heatingControlStorage) DoTask(heatingControl *entities.HeatingControl, task *entities.Task) error {
	if task.Active {
		switch task.Action {
		case actions.Activate:
			heatingControl.Activated = true
		case actions.Disable:
			heatingControl.Activated = false
		default:
			return errroz.WrongTaskAction
		}
		return s.db.Model(&entities.HeatingControl{}).Where("uuid = ?", heatingControl.UUID).Updates(heatingControl).Error
	}
	return errroz.TaskNotActive
}
