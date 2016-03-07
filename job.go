package main

type job interface {
	String() string
}

type jobV1 struct {
	Class string
	Args  []string
}

type jobV2 struct {
	Class string
	Vars  map[string]string
}
