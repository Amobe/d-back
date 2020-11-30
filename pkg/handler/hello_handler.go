package handler

import (
	"fmt"
	"net/http"

	"github.com/amobe/d-back/pkg/entity"
	"github.com/amobe/d-back/pkg/infra/webserver"
	"github.com/amobe/d-back/pkg/service/iplimiter"
)

// Hello represents a handler
type Hello struct {
	ils iplimiter.Service
}

// NewHelloHandler creates a hello handler instance.
func NewHelloHandler(ils iplimiter.Service) Hello {
	return Hello{
		ils: ils,
	}
}

// Handle handles the hello request.
// The remote address of the client is recorded.
func (h *Hello) Handle(req *http.Request) webserver.Response {
	ipAddress, err := entity.NewIPAddress(req.RemoteAddr)
	if err != nil {
		fmt.Printf("[Err] HelloHandle | new ip address: %s\n", err)
		return webserver.NewErrResponse(http.StatusBadRequest, err)
	}

	token, err := h.ils.AcceptRequest(ipAddress)
	if err != nil {
		fmt.Printf("[Err] HelloHandle | accept request: %s\n", err)
		return webserver.NewErrResponse(http.StatusServiceUnavailable, err)
	}

	resp := HelloResponse{
		Index: token.Index(),
	}
	return webserver.NewResponse(http.StatusOK, resp)
}
