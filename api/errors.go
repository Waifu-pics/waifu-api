package api

const (
	// ErrServer should be returned on 500 errors
	ErrServer = "there was an issue processing this request"

	// ErrInvalidJSON should be returned when body is invalid
	ErrInvalidJSON = "that was not a valid json body"

	// ErrInvalidLimit should be returned when limit is invalid
	ErrInvalidLimit = "limit query parameter is invalid, must be an integer which is greater than 1 and less than or equal to 30"
)
