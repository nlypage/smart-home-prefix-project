package errroz

import "errors"

// Declaring constants for error messages

var (
	WrongTaskAction = errors.New("wrong task action")
	TaskNotActive   = errors.New("task is not active")
)
