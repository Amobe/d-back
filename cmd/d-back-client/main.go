package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var (
		flagURL      string
		flagTimes    int
		flagInterval string
	)
	flag.StringVar(&flagURL, "url", "", "request api url")
	flag.IntVar(&flagTimes, "times", 0, "request times")
	flag.StringVar(&flagInterval, "interval", "1s", "interval between request")
	flag.Parse()

	if len(flagURL) == 0 {
		fmt.Println("invalid url")
		return
	}
	if flagTimes == 0 {
		fmt.Println("invalid times")
		return
	}
	duration, err := time.ParseDuration(flagInterval)
	if err != nil {
		fmt.Printf("invalid interval: %s\n", err)
		return
	}

	if flagTimes == -1 {
		for {
			sendRequest(flagURL)
			time.Sleep(duration)
		}
	} else {
		for i := 0; i < flagTimes; i++ {
			sendRequest(flagURL)
			time.Sleep(duration)
		}
	}
}

func sendRequest(url string) {
	resp, err := http.Post(url, "application/json", nil)
	if err != nil {
		fmt.Printf("[Err] post: %s\n", err)
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("[Err] body read: %s\n", err)
		return
	}
	fmt.Printf("status: %s, body: %s\n", resp.Status, data)
}
