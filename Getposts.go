package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var lock sync.Mutex
var Defaultskip = int64(0)
var Defaultlimit = int64(10)
var skip = Defaultskip
var limit = Defaultlimit

//CheckPosts : Returns a list of all posts of a user
func CheckPosts(email string) []Users {
	collection := client.Database("appointy").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	opts := options.Find()
	opts.SetSort{bson.D{
		{Key: "user", Value: bson.M{"$gt": Checkuser}},
	}
	opts.Skip = &skip
	cursor, _ := collection.Find(ctx, bson.D{
		{Key: "users.email", Value: email},
	}, opts)
	var usersreturn []Users
	var post Posts
	for cursor.Next(ctx) {
		cursor.Decode(&post)
		usersreturn = append(usersreturn, post)
	}
	return usersreturn
}

// GetsPots : Gets a list of all posts of a user
func GetPosts(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
		return
	}
	response.Header().Set("content-type", "application/json")
	fmt.Println((request.URL.Query()["post"][0]))
	if len(request.URL.Query()["limit"]) != 0 {
		limit, _ = strconv.ParseInt(request.URL.Query()["limit"][0], 0, 64)
	}
	if len(request.URL.Query()["ofset"]) != 0 {
		skip, _ = strconv.ParseInt(request.URL.Query()["offset"][0], 0, 64)
	}
	email := request.URL.Query()["post"][0]
	userposts := CheckPosts(email)
	if len(userposts) == 0 {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "Post not present" }`))
		return
	}
	json.NewEncoder(response).Encode(userposts)
	skip = Defaultskip
	limit = Defaultlimit
}
