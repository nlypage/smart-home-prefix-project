package dto

import (
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
)

// CreateTask is a struct that represents a DTO to create a new Task.
type CreateTask struct {
	TaskObjectUUID string          `json:"task_object_uuid"`
	Time           string          `json:"time"`         // "HH:MM"
	DaysOfWeek     string          `json:"days_of_week"` // "1,2,3,4,5,6,7"
	Action         entities.Action `json:"action"`
}

// UpdateTask is a struct that represents a DTO to update an existing Task
type UpdateTask struct {
	UUID           string          `json:"uuid"`
	TaskObjectUUID string          `json:"task_object_uuid"`
	Time           string          `json:"time"`         // "HH:MM"
	DaysOfWeek     string          `json:"days_of_week"` // "1,2,3,4,5,6,7"
	Action         entities.Action `json:"action"`
	Active         bool            `json:"active"`
}
