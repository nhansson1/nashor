package problem

import (
	"errors"
	"fmt"
)

type ErrorResponse struct {
	Status int
	Detail string
}

func (err ErrorResponse) Error() string {
	return err.Detail
}

func NewErrorResponse(s int, d string) ErrorResponse {
	return ErrorResponse{
		Status: s,
		Detail: d,
	}
}

func (e *ErrorResponse) FromErr(err error) {
	var perr RequestError
	t := InternalServerError()

	if errors.As(err, &perr) {
		t = perr.ToErrorResponse()
	} else {
		fmt.Println(err.Error())
	}

	e.Status = t.Status
	e.Detail = t.Detail
}

func InternalServerError() ErrorResponse {
	return NewErrorResponse(500, "Cannot complete your request, try again later.")
}

type RequestError struct {
	Status   int
	Endpoint string
}

func (err RequestError) Error() string {
	return fmt.Sprintf("Request to: %s returned status: %d", err.Endpoint, err.Status)
}

func (err *RequestError) ToErrorResponse() ErrorResponse {
	if err.Status == 404 {
		return NewErrorResponse(404, "Data not found.")
	}

	return InternalServerError()
}

func NewRequestError(s int, e string) RequestError {
	return RequestError{
		Status:   s,
		Endpoint: e,
	}
}
