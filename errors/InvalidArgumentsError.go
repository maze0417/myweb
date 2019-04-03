package errors

import "fmt"

type ValidationError struct {
	Err int
	Errdesc string
}
func (e ValidationError) Error() string {
	return fmt.Sprintf("%v: %v",e.Err, e.Errdesc)
}