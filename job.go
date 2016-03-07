package main

import (
	"encoding/json"
)

type job interface {
	String() string
}

type jobV1 struct {
	Class string   `json:"class"`
	Args  []string `json:"args"`
}

func (j jobV1) String() (string, error) {
	data, err := json.Marshal(j)
	return string(data), err
}

type jobV2 struct {
	Class string
	Vars  map[string]string
}
