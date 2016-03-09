package main

import (
	"net/url"
	"testing"
)

func TestJobV1String(test *testing.T) {
	job := jobV1{
		Class: "test",
		Args: []string{
			"a=1",
			"b=2",
			"c=3",
		},
	}

	expected := `{"class":"test","args":["a=1","b=2","c=3"]}`
	result, err := job.String()
	if err != nil {
		test.Errorf("Failed to convert jobV1 to string: %s", err.Error())
	}
	if result != expected {
		test.Errorf("Failed to convert jobV1 to string\nExpected: %s\nResult: %s", expected, result)
	}
}

func TestNewJobV1(test *testing.T) {
	class := "Test"
	data := url.Values{}
	data.Set("a", "1")
	data.Set("b", "2")
	data.Set("c", "3")

	job := newJobV1(class, data)

	if job.Class != "Test" {
		test.Errorf("Failed to set job class")
	}

	if len(job.Args) != 3 {
		test.Errorf("Failed to set job args")
	}
	for _, v := range job.Args {
		if (v != "a=1") && (v != "b=2") && (v != "c=3") {
			test.Errorf("Failed to set job args")
		}
	}
}
