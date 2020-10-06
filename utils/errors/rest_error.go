package errors

import (
	"github.com/dmazzella--/GoBasha_users-api/logger"
	"net/http"
)

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(msg string) {

}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func NewInternalServerErrorX(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}

func LogAndNewInternalServerError(dberr DBError, err error) *RestErr {
	logger.Error(dberr.GetFormattedMessage(), err)
	return &RestErr{
		Message: dberr.Id,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
