package main

import (
	"fmt"

	"github.com/fzzy/radix/redis"
)

type resque struct {
	client    *redis.Client
	Redis     string
	Namespace string
	Queue     string
}

func (res *resque) Init() (err error) {
	err = res.Dial()
	if err != nil {
		fmt.Printf("Failed to connect to redis: %s", err.Error())
		return err
	}

	reply := res.client.Cmd("SADD", res.Namespace+":queues", res.Queue)
	return reply.Err
}

func (res *resque) Dial() (err error) {
	res.client, err = redis.Dial("tcp", res.Redis)
	return err
}

func (res *resque) Enqueue(j job) (err error) {
	if res.client == nil {
		err = res.Init()
		if err != nil {
			return err
		}
	}

	queue := res.getFullQueueName()

	jobString, err := j.String()
	if err != nil {
		return err
	}

	reply := res.client.Cmd("RPUSH", queue, jobString)
	return reply.Err
}

func (res *resque) getFullQueueName() string {
	return res.Namespace + ":queue:" + res.Queue

}
