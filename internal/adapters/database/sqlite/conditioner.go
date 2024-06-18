package sqlite

import (
	"github.com/nlypage/smart-home-prefix-project/internal/domain/common/actions"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/common/errroz"
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

// GetByUUID is a method that returns a pointer to a Conditioner and error.
func (s *conditionerStorage) GetByUUID(uuid string) (*entities.Conditioner, error) {
	var conditioner *entities.Conditioner
	err := s.db.Model(&entities.Conditioner{}).Where("uuid = ?", uuid).First(conditioner).Error
	return conditioner, err
}

// GetAll is a method that returns a slice of pointers to Conditioner instances.
func (s *conditionerStorage) GetAll(limit, offset int) ([]*entities.Conditioner, error) {
	var conditioners []*entities.Conditioner
	err := s.db.Model(&entities.Conditioner{}).Limit(limit).Offset(offset).Find(conditioners).Error
	return conditioners, err
}

// Create is a method to create a new Conditioner in database.
func (s *conditionerStorage) Create(conditioner *entities.Conditioner) (*entities.Conditioner, error) {
	err := s.db.Create(conditioner).Error
	return conditioner, err
}

// Update is a method to update an existing Conditioner in database.
func (s *conditionerStorage) Update(conditioner *entities.Conditioner) (*entities.Conditioner, error) {
	err := s.db.Model(&entities.Conditioner{}).Where("uuid = ?", conditioner.UUID).Updates(conditioner).Error
	return conditioner, err
}

// Delete is a method to delete an existing Conditioner in database.
func (s *conditionerStorage) Delete(uuid string) error {
	return s.db.Unscoped().Delete(&entities.Conditioner{}, "uuid = ?", uuid).Error
}

// DoTask is a method to perform a task on a Conditioner.
func (s *conditionerStorage) DoTask(conditioner *entities.Conditioner, task *entities.Task) error {
	if task.Active {
		switch task.Action {
		case actions.Activate:
			conditioner.Activated = true
		case actions.Disable:
			conditioner.Activated = false
		default:
			return errroz.WrongTaskAction
		}
		return s.db.Model(&entities.Conditioner{}).Where("uuid = ?", conditioner.UUID).Updates(conditioner).Error
	}
	return errroz.TaskNotActive
}
