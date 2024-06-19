package services

import (
	"github.com/nlypage/smart-home-prefix-project/internal/domain/dto"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
)

// TaskStorage is an interface that contains methods to interact with the database.
type TaskStorage interface {
	GetByUUID(uuid string) (*entities.Task, error)
	GetAll(limit, offset int) ([]*entities.Task, error)
	Create(task *entities.Task) (*entities.Task, error)
	Update(bulb *entities.Task) (*entities.Task, error)
	Delete(uuid string) error
	GetTaskObject(uuid string) (*entities.TaskObject, error)
}

// taskService is a struct that contains a pointer to a TaskStorage instance.
type taskService struct {
	storage TaskStorage
}

// NewTaskService is a function that returns a new instance of taskService.
func NewTaskService(storage TaskStorage) *taskService {
	return &taskService{storage: storage}
}

// Create is a method to create a new Task in database using a CreateTask DTO.
func (s taskService) Create(createTask *dto.CreateTask) (*entities.Task, error) {
	taskObject, err := s.storage.GetTaskObject(createTask.TaskObjectUUID)
	if err != nil {
		return nil, err
	}

	task := &entities.Task{
		TaskObject: taskObject,
		Time:       createTask.Time,
		DaysOfWeek: createTask.DaysOfWeek,
		Action:     createTask.Action,
	}

	return s.storage.Create(task)
}

// GetByUUID is a method that returns a pointer to a Task instance and error.
func (s taskService) GetByUUID(uuid string) (*entities.Task, error) {
	return s.storage.GetByUUID(uuid)
}

// GetAll is a method that returns a slice of pointers to a Task and error.
func (s taskService) GetAll(limit, offset int) ([]*entities.Task, error) {
	return s.storage.GetAll(limit, offset)
}

// Update is a method to update an existing Task in database using an UpdateTask DTO.
func (s taskService) Update(updateTask *dto.UpdateTask) (*entities.Task, error) {
	var task *entities.Task

	taskObject, err := s.storage.GetTaskObject(updateTask.TaskObjectUUID)
	if err != nil {
		return nil, err
	}

	task.UUID = updateTask.UUID
	task.TaskObject = taskObject
	task.Time = updateTask.Time
	task.DaysOfWeek = updateTask.DaysOfWeek
	task.Action = updateTask.Action
	task.Active = updateTask.Active
	return s.storage.Update(task)
}

// Delete is a method to delete an existing Task in database.
func (s taskService) Delete(uuid string) error {
	return s.storage.Delete(uuid)
}

//func (s taskService) StartTasksExecution() error {
//	tasks, err := s.storage.GetAll(0, 0)
//	if err != nil {
//		return err
//	}
//
//	for _, task := range tasks {
//		if task.Active {
//			if task.Time == time.Now().Format("15:04") && strings.Contains(task.DaysOfWeek, time.Now().Weekday().String()) {
//				taskObject := *task.TaskObject
//				switch v := taskObject.(type) {
//				case *entities.AlarmSystem:
//
//				case *entities.Bulb:
//					// Do something
//				case *entities.Conditioner:
//					// Do something
//				case *entities.HeatingControl:
//					// Do something
//				}
//			}
//		}
//	}
//}
