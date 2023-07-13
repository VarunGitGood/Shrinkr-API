package types

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

type CustomError struct {
	Message string
	Status  int
}

func (e *CustomError) Error() string {
	return e.Message
}

func (e *CustomError) StatusCode() int {
	return e.Status
}
