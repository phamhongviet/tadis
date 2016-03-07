package main

import (
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
