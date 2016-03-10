package main

import (
	"net/http"
	"fmt"
)

func handler(writer http.ResponseWriter, request *http.Request, queue enqueuer, class string) {
	data, err := parseRequest(request)
	if err != nil {
		fmt.Printf("Failed to parse request: %s", err.Error())
		return
	}

	job := newJobV1(class, data)

	err = queue.Enqueue(&job)
	if err != nil {
		fmt.Printf("Failed to enqueue job: %s", err.Error())
	}
}
