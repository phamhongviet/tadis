package main

import (
	"encoding/json"
	"net/url"
)

type job interface {
	String() (string, error)
}

type jobV1 struct {
	Class string   `json:"class"`
	Args  []string `json:"args"`
}

func (j jobV1) String() (string, error) {
	data, err := json.Marshal(j)
	return string(data), err
}

func newJobV1(class string, data url.Values) jobV1 {
	var args []string
	for k, v := range data {
		for _, vv := range v {
			args = append(args, k+"="+vv)
		}
	}

	return jobV1{
		Class: class,
		Args:  args,
	}
}

type jobV2 struct {
	Class string
	Vars  map[string]string
}
