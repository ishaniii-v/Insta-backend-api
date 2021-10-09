package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Posts struct {
	ID              string `json:"Id,omitempty"`
	Caption         string `json:"Caption,omitempty"`
	Imageurl        string `json:"Imageurl,omitempty"`
	Postedtimestamp int    `json:"Postedtimestamp,omitempty"`
}

//CreatePost : Create a post to the database
func CreatePost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var user Users
	_ = json.NewDecoder(request.Body).Decode(&user)
	user.def()
	if user.ID = " "{
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "Error" }`))
		return
	}
	collection := client.Database("appointy").Collection("Users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, user)
	user.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(response).Encode(user)
	fmt.Println(user)

	for _, param := range params {
		switch s := param.(type) {
		case string:
			user.Posttimestamp = s
		}
	}

	return user
}
