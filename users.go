package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Users struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:",omitempty"`
}

func CreateUser(response http.ResponseWriter, request *http.Request) {

	var user Users
	_ = json.NewDecoder(request.Body).Decode(&user)
	response.Header().Set("content-type", "application/json")
	collection := client.Database("appointy").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	result, _ := collection.InsertOne(ctx, user)
	user.ID = result.InsertedID.(primitive.ObjectID)
	json.NewEncoder(response).Encode(user)
	fmt.Println(user)
}
