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

	err := resque.Enqueue(j)
	if err != nil {
		test.Errorf("Failed to enqueue: %s", err.Error())
	}

	queue := resque.Namespace + ":queue:" + resque.Queue
	reply := resque.client.Cmd("LPOP", queue)
	if reply.Err != nil {
		test.Errorf("Failed to get enqueued job: %s", err.Error())
	}

	jobString, _ := j.String()
	if reply.String() != jobString {
		test.Errorf("Jobs' strings in resque are mismatched")
	}
}

func redisTestServerAddress() string {
	return os.Getenv("REDIS_PORT_6379_TCP_ADDR") + ":" + os.Getenv("REDIS_PORT_6379_TCP_PORT")
}
