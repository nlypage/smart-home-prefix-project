package sqlite

import (
	"github.com/nlypage/smart-home-prefix-project/internal/domain/common/actions"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/common/errroz"
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

// GetByUUID is a method that returns an error and a pointer to a Bulb instance.
func (s *bulbStorage) GetByUUID(uuid string) (*entities.Bulb, error) {
	var bulb *entities.Bulb
	err := s.db.Model(&entities.Bulb{}).Where("uuid = ?", uuid).First(&bulb).Error
	return bulb, err
}

// GetAll is a method that returns a slice of pointers to Bulb instances.
func (s *bulbStorage) GetAll(limit, offset int) ([]*entities.Bulb, error) {
	var bulbs []*entities.Bulb
	err := s.db.Model(&entities.Bulb{}).Limit(limit).Offset(offset).Find(&bulbs).Error
	return bulbs, err
}

// Create is a method to create a new Bulb in database.
func (s *bulbStorage) Create(bulb *entities.Bulb) (*entities.Bulb, error) {
	err := s.db.Create(&bulb).Error
	return bulb, err
}

// Update is a method to update an existing Bulb in database.
func (s *bulbStorage) Update(bulb *entities.Bulb) (*entities.Bulb, error) {
	err := s.db.Model(&entities.Bulb{}).Where("uuid = ?", bulb.UUID).Updates(&bulb).Error
	return bulb, err
}

// Delete is a method to delete an existing Bulb in database.
func (s *bulbStorage) Delete(uuid string) error {
	return s.db.Unscoped().Delete(&entities.Bulb{}, "uuid = ?", uuid).Error
}

// DoTask is a method to perform a task on a Bulb.
func (s *bulbStorage) DoTask(buLb *entities.Bulb, task *entities.Task) error {
	if task.Active {
		switch task.Action {
		case actions.Activate:
			buLb.Activated = true
		case actions.Disable:
			buLb.Activated = false
		default:
			return errroz.WrongTaskAction
		}
		return s.db.Model(&entities.AlarmSystem{}).Where("uuid = ", buLb.UUID).Updates(&buLb).Error
	}
	return errroz.TaskNotActive
}
