package rest_err

import (
	"net/http"
)

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (r *RestErr) Error() string {
	return r.Message
}

//Cria um construtor

func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

//Criação dos metodos para cada tipo de erro que estamos trabalhando

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "bad request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "internal server error",
		Code:    http.StatusInternalServerError,
	}
}

func NewRequestForbiddenError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewRequestUnauthorizedError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Unauthorized error",
		Code:    http.StatusUnauthorized,
	}
}

func NewRequestNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     "Error Not Found",
		Code:    http.StatusNotFound,
	}
}
