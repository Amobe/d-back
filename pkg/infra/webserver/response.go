package webserver

// Response represents the general response of the webserver.
type Response struct {
	code int
	body interface{}
}

// NewResponse creates the instance of the response includes json body.
func NewResponse(code int, body interface{}) Response {
	return Response{
		code: code,
		body: body,
	}
}

// NewErrResponse creates the instance of err response.
func NewErrResponse(code int, err error) Response {
	return Response{
		code: code,
		body: ErrResponse{Message: err.Error()},
	}
}

// ErrResponse represent the response with err message.
type ErrResponse struct {
	Message string `json:"message"`
}
