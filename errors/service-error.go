package errors

// ServiceError should be used to return business error message
type ServiceError struct {
	Message string `json: "message"`
}
