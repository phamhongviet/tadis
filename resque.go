package main

import (
	"github.com/fzzy/radix/redis"
)

type resque struct {
	client    redis.Client
	namespace string
	queue     string
}

func (*resque) Enqueue(j *job) error {
	return nil
}
