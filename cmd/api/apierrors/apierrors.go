package apierrors

import (
	"fmt"
	"net/http"
)

type (
	CauseList []interface{}
	ApiError  interface {
		Message() string
		Code() string
		Status() int
		Cause() CauseList
		Error() string
	}
	apiError struct {
		ErrorMessage string    `json:"message"`
		ErrorCode    string    `json:"code"`
		ErrorStatus  int       `json:"status"`
		ErrorCause   CauseList `json:"cause"`
	}
)

func (e apiError) Message() string {
	return e.ErrorMessage
}

func (e apiError) Code() string {
	return e.ErrorCode
}

func (e apiError) Status() int {
	return e.ErrorStatus
}

func (e apiError) Cause() CauseList {
	return e.ErrorCause
}

func (e apiError) Error() string {
	return fmt.Sprintf("Message: %s; Error Code: %s;Status: %d;Cause: %v", e.ErrorMessage,
		e.ErrorCode, e.ErrorStatus, e.ErrorCause)
}

func (c CauseList) ToString() string {
	return fmt.Sprint(c)
}

func NewNotFoundApiError(message string) ApiError {
	return apiError{
		ErrorMessage: message,
		ErrorCode:    "not_found",
		ErrorStatus:  http.StatusNotFound,
		ErrorCause:   CauseList{},
	}
}
