package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Hello represents a handler
type Hello struct {
}

// NewHelloHandler creates a hello handler instance.
func NewHelloHandler() Hello {
	return Hello{}
}

// Handle handles the hello request.
// The remote address of the client is recorded.
func (h *Hello) Handle(ctx *gin.Context) {
	fmt.Println(ctx.Request.RemoteAddr)
}
