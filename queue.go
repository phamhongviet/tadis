package main

type enqueuer interface {
	Enqueue(*job) error
}
