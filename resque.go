package main

import (
	"github.com/fzzy/radix/redis"
)

type resque struct {
	client    *redis.Client
	Redis     string
	Namespace string
	Queue     string
}

func (res *resque) Dial() (err error) {
	res.client, err = redis.Dial("tcp", res.Redis)
	return err
}

func (res *resque) Enqueue(j job) error {
	queue := res.Namespace + ":queue:" + res.Queue

	jobString, err := j.String()
	if err != nil {
		return err
	}

	reply := res.client.Cmd("RPUSH", queue, jobString)
	return reply.Err
}
