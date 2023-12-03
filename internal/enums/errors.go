package enums

type MessageStatus string

const (
	ErrBadRequest   MessageStatus = "Bad Request"
	ErrNotFound     MessageStatus = "Not Found"
	ErrForbidden    MessageStatus = "Forbidden"
	ErrConflict     MessageStatus = "Conflict"
	ErrInternal     MessageStatus = "Internal Server Error"
	ErrUnauthorized MessageStatus = "Unauthorized"
	StatusOK        MessageStatus = "OK"
	StatusCreated   MessageStatus = "Created"
)

type ErrorMessage string

const (
	ErrInvalidEmailOrPassword ErrorMessage = "Invalid email or password"
	ErrInvalidRefreshToken    ErrorMessage = "Invalid refresh token"
	ErrInvalidAccessToken     ErrorMessage = "Invalid access token"
	ErrInvalidToken           ErrorMessage = "Invalid token"
	ErrExpiredToken           ErrorMessage = "Expired token"
	ErrInvalidRequest         ErrorMessage = "Invalid request"
	ErrInvalidID              ErrorMessage = "Invalid ID"
)
