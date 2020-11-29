package main

import (
	"fmt"
	"net/http"

	"github.com/amobe/d-back/pkg/infra/inmem"

	"github.com/amobe/d-back/pkg/handler"
	"github.com/amobe/d-back/pkg/infra/webserver"
	"github.com/amobe/d-back/pkg/service/iplimiter"
)

func main() {
	const listenAddr = "0.0.0.0:10080"

	server := webserver.NewGinServer(listenAddr)

	ipLimiterRepository := inmem.NewIPLimiterRepository()
	ipLimiterService := iplimiter.NewIPLimiterService(ipLimiterRepository)

	hello := handler.NewHelloHandler(ipLimiterService)

	err := server.AddRouter(http.MethodPost, "/hello", hello.Handle)
	if err != nil {
		fmt.Printf("fail to add hello route: %s\n", err)
		return
	}

	server.Start()
}
