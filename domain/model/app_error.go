package model

// AppError holds metadata about global app.
type AppError struct {
	Err     error
	Message string
	Code    int
	Stack   []byte
}
