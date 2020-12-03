package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/amobe/d-back/pkg/handler"
	"github.com/amobe/d-back/pkg/infra/inmem"
	"github.com/amobe/d-back/pkg/infra/webserver"
	"github.com/amobe/d-back/pkg/service/iplimiter"
)

func main() {
	envPort := os.Getenv("PORT")
	envLimitNumber := os.Getenv("NUMBER")
	envLimitDuration := os.Getenv("DURATION")

	if _, err := strconv.ParseInt(envPort, 10, 16); err != nil {
		fmt.Printf("invalid port: %v", err)
		return
	}
	limitNumber64, err := strconv.ParseUint(envLimitNumber, 10, 32)
	if err != nil {
		fmt.Printf("invalid limit number: %v", err)
		return
	}
	limitNumber := uint32(limitNumber64)
	limitDuration, err := time.ParseDuration(envLimitDuration)
	if err != nil {
		fmt.Printf("invalid limit duration: %v", err)
		return
	}

	fmt.Printf("[Info] Start d-back limit %d requests per %s\n", limitNumber, limitDuration)

	server := webserver.NewGinServer(fmt.Sprintf("0.0.0.0:%s", envPort))

	ipLimiterRepository := inmem.NewIPLimiterRepository()
	ipLimiterService := iplimiter.NewIPLimiterService(ipLimiterRepository, limitNumber, limitDuration)

	hello := handler.NewHelloHandler(ipLimiterService)

	if err := server.AddRouter(http.MethodPost, "/hello", hello.Handle); err != nil {
		fmt.Printf("fail to add hello route: %s\n", err)
		return
	}

	server.Start()
}
