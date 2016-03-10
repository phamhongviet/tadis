package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(test *testing.T) {
	recorder := httptest.NewRecorder()

	body := strings.NewReader("a=1&b=2&c=3")
	request, err := http.NewRequest("POST", "http://handler.test/api", body)
	if err != nil {
		fmt.Println(err.Error())
	}

	class := "test"

	resque := &resque{
		Redis:     redisTestServerAddress(),
		Namespace: "resque",
		Queue:     "test",
	}

	handler(recorder, request, resque, class)

	err = resque.Dial()
	if err != nil {
		test.Errorf("Failed to connect to redis: %s", err.Error())
	}

	reply := resque.client.Cmd("RPOP", resque.getFullQueueName())
	if reply.Err != nil {
		test.Errorf("Failed to get data from redis: %s", err.Error())
	}
	if reply.String() != `{"class":"test","args":["a=1","b=2","c=3"]}` {
		test.Errorf("Enqueued job and data in redis mismatched")
	}
}
