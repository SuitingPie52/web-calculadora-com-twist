package main

import(
	"net/http"
	
)

var ResultController = map[string]map[string]func(http.ResponseWriter, *http.Request) {

	"GET": GetFunctions,

}

var GetFunctions = map[string]func(http.ResponseWriter, *http.Request) {
	
	"/result": GetResult,
	
}

func GetResult (res http.ResponseWriter, req *http.Request) {

	res.Write([]byte("achou"))
	return

}

