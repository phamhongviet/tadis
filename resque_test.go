package main

import (
	"os"
	"testing"
)

func TestResqueEnqueue(test *testing.T) {
	j := &jobV1{
		Class: "Test",
		Args: []string{
			"a=1",
			"b=2",
			"c=3",
		},
	}

	resque := resque{
		Redis:     redisTestServerAddress(),
		Namespace: "resque",
		Queue:     "test",
	}

	err := resque.Dial()
	if err != nil {
		test.Errorf("Failed to connect to redis: %s", err.Error())
	}
	err = resque.Enqueue(j)
	if err != nil {
		test.Errorf("Failed to enqueue: %s", err.Error())
	}
}

func redisTestServerAddress() string {
	return os.Getenv("REDIS_PORT_6379_TCP_ADDR") + ":" + os.Getenv("REDIS_PORT_6379_TCP_PORT")
}
