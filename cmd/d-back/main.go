package main

import (
	"fmt"
	"net/http"

	"github.com/amobe/d-back/pkg/handler"
	"github.com/amobe/d-back/pkg/infra/webserver"
)

func main() {
	const listenAddr = "0.0.0.0:10080"

	server := webserver.NewGinServer(listenAddr)

	hello := handler.NewHelloHandler()

	err := server.AddRouter(http.MethodPost, "/request", hello.Handle)
	if err != nil {
		fmt.Printf("fail to add hello route: %s\n", err)
		return
	}

	server.Start()
}
