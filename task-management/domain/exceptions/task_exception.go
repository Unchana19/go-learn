package exceptions

import "errors"

var (
	ErrInvalidPriority = errors.New("invalid priority")
	ErrTaskNotFound    = errors.New("task not found")
	ErrInvalidStatus   = errors.New("invalid status")
)
