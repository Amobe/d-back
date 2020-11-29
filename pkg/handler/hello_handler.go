package handler

import (
	"fmt"
	"net/http"

	"github.com/amobe/d-back/pkg/entity"
	"github.com/amobe/d-back/pkg/service/iplimiter"

	"github.com/gin-gonic/gin"
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
func (h *Hello) Handle(ctx *gin.Context) {
	ipAddress, err := entity.NewIPAddress(ctx.Request.RemoteAddr)
	if err != nil {
		fmt.Printf("[Err] HelloHandle | new ip address: %s\n", err)
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	token, err := h.ils.AcceptRequest(ipAddress)
	if err != nil {
		fmt.Printf("[Err] HelloHandle | accept request: %s\n", err)
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}
	fmt.Println(token)
}
