package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(test *testing.T) {
	recorder := httptest.NewRecorder()

	body := strings.NewReader("a=1&b=2&c=3")
	request, err := http.NewRequest("POST", "http://handler.test/api", body)
	if err != nil {
		fmt.Println(err.Error())
	}

	class := "test"

	handler(recorder, request, class)
}
