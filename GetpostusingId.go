package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"path"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CheckPostusingID : Checks the post with the provided id
func CheckPostusingID(id primitive.ObjectID) (Posts, error) {
	var post Posts
	collection := client.Database("appointy").Collection("Post")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, Posts{ID: id}).Decode(&post)
	if post.ID != id {
		err = errors.New("Error 400: ID not present")
	}
	return post, err
}

//GetPostusingID : Gives the post with the provided id
func GetPostusingID(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte(`{ "message": "Incorrect Method" }`))
		return
	}
	response.Header().Set("content-type", "application/json")
	fmt.Println(path.Base(request.URL.Path))
	id, _ := primitive.ObjectIDFromHex(path.Base(request.URL.Path))
	postusingID, err := CheckPostusingID(id)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(postusingID)

}
