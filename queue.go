package main

type enqueuer interface {
	Enqueue(*job)
}
