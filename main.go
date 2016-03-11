package main

import (
	"fmt"
	"net/http"
)

const (
	defaultRedisAddress   = "127.0.0.1:6379"
	defaultRedisNameSpace = "resque"
	defaultQueue          = "tadis"
	defaultClass          = "tadis"
	defaultURL            = "/api"
	defaultPort           = "7415"
)

func main() {
	resque := &resque{
		Redis:     defaultRedisAddress,
		Namespace: defaultRedisNameSpace,
		Queue:     defaultQueue,
	}
	resque.Init()

	http.HandleFunc(defaultURL, func(writer http.ResponseWriter, request *http.Request) {
		handler(writer, request, resque, defaultClass)
	})

	err := http.ListenAndServe(":"+defaultPort, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
