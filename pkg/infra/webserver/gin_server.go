package webserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GinServer represents a web service implement with gin.
type GinServer struct {
	addr   string
	router *gin.Engine
}

// NewGinServer creates a GinServer instance.
// Given addr to select the listening addr and port of web service.
func NewGinServer(addr string) *GinServer {
	return &GinServer{
		addr:   addr,
		router: gin.Default(),
	}
}

// Start runs the web service process.
func (s *GinServer) Start() {
	s.router.Run(s.addr)
}

// AddRouter registers a request handler with given path and method.
func (s *GinServer) AddRouter(method string, path string, handler func(*http.Request) Response) error {
	switch method {
	case http.MethodGet:
		s.router.GET(path, ginHandlerWrapper(handler))
		return nil
	case http.MethodPost:
		s.router.POST(path, ginHandlerWrapper(handler))
		return nil
	}
	return fmt.Errorf("invalid method")
}

// ginHandlerWrapper wraps the handler becomes a gin handler.
func ginHandlerWrapper(handler func(*http.Request) Response) func(*gin.Context) {
	return func(ctx *gin.Context) {
		resp := handler(ctx.Request)
		ctx.JSON(resp.code, resp.body)
	}
}
