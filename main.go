package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("Instagram Running")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	Users, _ = mongo.Connect(ctx, clientOptions)

	http.HandleFunc("/users", instahandler)
	http.HandleFunc("/posts/", GetPosts)
	http.HandleFunc("/user/", GetUserusingID)
	http.HandleFunc("/posts/", GetPostusingID)
	http.ListenAndServe(":12345", nil)
}
