package problem

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

func InternalServerError() ErrorResponse {
	return NewErrorResponse(500, "Cannot complete your request, try again later.")
}
