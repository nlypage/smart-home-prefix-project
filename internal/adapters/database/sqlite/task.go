package sqlite

import (
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
	"gorm.io/gorm"
)

// taskStorage is a struct that contains a pointer to a gorm.DB instance.
type taskStorage struct {
	db *gorm.DB
}

// NewTaskStorage is a function that returns a new instance of taskStorage.
func NewTaskStorage(db *gorm.DB) *taskStorage {
	return &taskStorage{db: db}
}

// GetByUUID is a method that returns a pointer to a Task instance and error.
func (s *taskStorage) GetByUUID(uuid string) (*entities.Task, error) {
	var task *entities.Task
	err := s.db.Model(&entities.Task{}).Where("uuid = ?", uuid).First(task).Error
	return task, err
}

// GetAll is a method that returns a slice of pointers to Task instances.
func (s *taskStorage) GetAll(limit, offset int) ([]*entities.Task, error) {
	var tasks []*entities.Task
	err := s.db.Model(&entities.Task{}).Limit(limit).Offset(offset).Find(tasks).Error
	return tasks, err
}

// Create is a method to create a new Task in database.
func (s *taskStorage) Create(task *entities.Task) (*entities.Task, error) {
	err := s.db.Create(task).Error
	return task, err
}

// Update is a method to update an existing Task in database.
func (s *taskStorage) Update(task *entities.Task) (*entities.Task, error) {
	err := s.db.Model(&entities.Task{}).Where("uuid = ", task.UUID).Updates(task).Error
	return task, err
}

// Delete is a method to delete an existing Task in database.
func (s *taskStorage) Delete(uuid string) error {
	err := s.db.Unscoped().Delete(&entities.Task{}, "uuid = ?", uuid).Error
	return err
}

// GetTaskObject is a method that returns a pointer to a TaskObject instance and error.
func (s *taskStorage) GetTaskObject(uuid string) (*entities.TaskObject, error) {
	var taskObject *entities.TaskObject
	err := s.db.Where("uuid = ?", uuid).First(taskObject).Error
	return taskObject, err
}
