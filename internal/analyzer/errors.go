package analyzer

import "fmt"

type LogError struct {
	ID      string
	Message string
	Err     error
}

func (e *LogError) Error() string {
	return fmt.Sprintf("[%s] %s: %v", e.ID, e.Message, e.Err)
}

func (e *LogError) Unwrap() error {
	return e.Err
}
