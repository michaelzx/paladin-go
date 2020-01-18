package errs

import "net/http"

// BadRequestError 业务逻辑错误
// 请求不符合要求、业务不符合条件等
type BadRequestError struct {
	Message string
	Status  int
	Code    int
}

func NewBadRequestError(message string) *BadRequestError {
	return &BadRequestError{Message: message, Status: http.StatusBadRequest, Code: 10000}
}
func NewBadRequestErrorWithCode(code int, message string) *BadRequestError {
	return &BadRequestError{Message: message, Status: http.StatusBadRequest, Code: code}
}

func (e *BadRequestError) Error() string {
	return e.Message
}
func (e *BadRequestError) Suffix(msg string) *BadRequestError {
	return &BadRequestError{
		Message: e.Message + msg,
		Status:  e.Status,
		Code:    e.Code,
	}
}

func (e *BadRequestError) Prefix(msg string) *BadRequestError {
	return &BadRequestError{
		Message: msg + e.Message,
		Status:  e.Status,
		Code:    e.Code,
	}
}

// Unauthorized 用户未登录
type Unauthorized struct {
	Message string
	Status  int
}

func NewUnauthorized(message string) *Unauthorized {
	return &Unauthorized{Message: message, Status: http.StatusUnauthorized}
}

func (e *Unauthorized) Error() string {
	return e.Message
}

// Forbidden 用户未授权
type Forbidden struct {
	Message string
	Status  int
}

func NewForbidden(message string) *Forbidden {
	return &Forbidden{Message: message, Status: http.StatusForbidden}
}

func (e *Forbidden) Error() string {
	return e.Message
}

// Forbidden 用户未授权
type Unknown struct {
	Message string
	Status  int
}

func NewUnknown(message string) *Unknown {
	return &Unknown{Message: message, Status: http.StatusInternalServerError}
}

func (e *Unknown) Error() string {
	return e.Message
}
