package task

import (
	"github.com/nlypage/smart-home-prefix-project/internal/domain/common/errroz"
	"github.com/nlypage/smart-home-prefix-project/internal/domain/entities"
	"log"
	"strings"
	"time"
)

// Service is an interface that contains methods to interact with tasks.
type Service interface {
	GetAll(limit, offset int) ([]*entities.Task, error)
	GetTaskObject(uuid string) (*entities.TaskObject, error)
}

// AlarmSystemService is an interface that contains a method to execute an AlarmSystem task.
type AlarmSystemService interface {
	DoTask(alarmSystem *entities.AlarmSystem, task *entities.Task) error
}

// BulbService is an interface that contains a method to execute a Bulb task.
type BulbService interface {
	DoTask(bulb *entities.Bulb, task *entities.Task) error
}

// ConditionerService is an interface that contains a method to execute a Conditioner task.
type ConditionerService interface {
	DoTask(conditioner *entities.Conditioner, task *entities.Task) error
}

// HeatingControlService is an interface that contains a method to execute a HeatingControl task.
type HeatingControlService interface {
	DoTask(heatingController *entities.HeatingControl, task *entities.Task) error
}

// taskUseCase is a struct that contains instances of services.
type taskUseCase struct {
	taskService           Service
	alarmSystemService    AlarmSystemService
	bulbService           BulbService
	conditionerService    ConditionerService
	heatingControlService HeatingControlService
}

// NewTaskUseCase is a function that returns a new instance of taskUseCase.
func NewTaskUseCase(
	taskService Service,
	alarmSystemService AlarmSystemService,
	bulbService BulbService,
	conditionerService ConditionerService,
	heatingControlService HeatingControlService,
) *taskUseCase {
	return &taskUseCase{
		taskService:           taskService,
		alarmSystemService:    alarmSystemService,
		bulbService:           bulbService,
		conditionerService:    conditionerService,
		heatingControlService: heatingControlService,
	}
}

// StartTasksExecution is a method that starts the execution of tasks.
func (u *taskUseCase) StartTasksExecution() error {
	for {
		tasks, err := u.taskService.GetAll(0, 0)
		if err != nil {
			return err
		}

		for _, task := range tasks {
			if task.Active {
				// TODO: Add timezone
				if task.Time == time.Now().Format("15:04") && strings.Contains(task.DaysOfWeek, time.Now().Weekday().String()) {
					taskObject := *task.TaskObject
					var errDoTask error
					switch v := taskObject.(type) {
					case *entities.AlarmSystem:
						errDoTask = u.alarmSystemService.DoTask(v, task)
					case *entities.Bulb:
						errDoTask = u.bulbService.DoTask(v, task)
					case *entities.Conditioner:
						errDoTask = u.conditionerService.DoTask(v, task)
					case *entities.HeatingControl:
						errDoTask = u.heatingControlService.DoTask(v, task)
					default:
						errDoTask = errroz.InvalidTaskObject
					}

					if errDoTask != nil {
						log.Println(errDoTask)
					}

				}
			}
		}
		time.Sleep(1*time.Minute + 1*time.Second)
	}
}
