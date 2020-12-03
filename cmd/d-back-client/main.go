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
		flagInterval int
	)
	flag.StringVar(&flagURL, "url", "", "request api url")
	flag.IntVar(&flagTimes, "times", 0, "request times")
	flag.IntVar(&flagInterval, "interval", 100, "interval between request in millisecond")
	flag.Parse()

	if len(flagURL) == 0 {
		fmt.Println("invalid url")
		return
	}
	if flagTimes == 0 {
		fmt.Println("invalid times")
		return
	}
	if flagInterval < 0 {
		fmt.Println("invalid interval")
		return
	}

	if flagTimes == -1 {
		for {
			sendRequest(flagURL)
			time.Sleep(time.Millisecond * time.Duration(flagInterval))
		}
	} else {
		for i := 0; i < flagTimes; i++ {
			sendRequest(flagURL)
			time.Sleep(time.Millisecond * time.Duration(flagInterval))
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
