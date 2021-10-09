package main

import "net/http"

func instahandler(response http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		CreateUser(response, request)
		CreatePost(response, request)
	} else if request.Method == "GET" {
		GetUserusingID(response, request)
		GetPostusingID(response, request)
	} else {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
	}
}
