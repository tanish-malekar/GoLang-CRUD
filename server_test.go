package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

//some values for testing
const (
	Id       string = "id6"
	Name     string = "Test Name"
	Email    string = "test@gmail.com"
	Password string = "pass123"
	Caption  string = "Test Caption"
	ImageURL string = "sampleURL.com"
	UserId   string = "id1"
)

func TestPostUser(t *testing.T) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tanish:password%23123@cluster0.ouv6u.mongodb.net/AppointyTaskDB?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	userHandler := newUserHandler(client)

	// Create new HTTP request
	var jsonStr = []byte(`{"Id":"` + Id + `","Name":"` + Name + `","Email":"` + Email + `","Password":"` + Password + `"}`)
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonStr))

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(userHandler.postUser)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Assert HTTP response
	expected := `User successfully added`
	if strings.TrimSpace(response.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}

func TestGetUserByID(t *testing.T) {
	//conecting to mongo database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tanish:password%23123@cluster0.ouv6u.mongodb.net/AppointyTaskDB?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	userHandler := newUserHandler(client)

	// Create new HTTP request
	req, _ := http.NewRequest("GET", "/users/"+Id, nil)

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(userHandler.getUserByID)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var user User
	json.NewDecoder(io.Reader(response.Body)).Decode(&user)

	// Assert HTTP response
	var expected User = User{
		Id:       Id,
		Name:     Name,
		Email:    Email,
		Password: Password,
	}
	if user != expected {
		t.Errorf("Incorrect response recived: %v", user)
	}

}

func TestPostPost(t *testing.T) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tanish:password%23123@cluster0.ouv6u.mongodb.net/AppointyTaskDB?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	postHandler := newPostHandler(client)

	// Create new HTTP request
	var jsonStr = []byte(`{"Id":"` + Id + `","Caption":"` + Caption + `","ImageURL":"` + ImageURL + `","UserID":"` + UserId + `"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonStr))

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(postHandler.postPost)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Assert HTTP response
	expected := `Post successfully added`
	if strings.TrimSpace(response.Body.String()) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			response.Body.String(), expected)
	}
}

func TestGetPostByID(t *testing.T) {
	//conecting to mongo database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tanish:password%23123@cluster0.ouv6u.mongodb.net/AppointyTaskDB?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	postHandler := newPostHandler(client)

	// Create new HTTP request
	req, _ := http.NewRequest("GET", "/posts/"+Id, nil)

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(postHandler.getPostByID)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Decode HTTP response
	var post Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// Assert HTTP response
	var expected Post = Post{
		Id:       Id,
		Caption:  Caption,
		ImageURL: ImageURL,
		UserId:   UserId,
	}

	if post.Id != expected.Id || post.Caption != expected.Caption || post.ImageURL != expected.ImageURL || post.UserId != expected.UserId {
		t.Errorf("Incorrect response recived: %v", post)
	}

}

func TestGetPostsByUserID(t *testing.T) {
	//conecting to mongo database
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://tanish:password%23123@cluster0.ouv6u.mongodb.net/AppointyTaskDB?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	postHandler := newPostHandler(client)

	// Create new HTTP request
	req, _ := http.NewRequest("GET", "/posts/users/"+Id, nil)

	// Assing HTTP Request handler Function (controller function)
	handler := http.HandlerFunc(postHandler.getPostsByUserID)
	// Record the HTTP Response
	response := httptest.NewRecorder()
	// Dispatch the HTTP Request
	handler.ServeHTTP(response, req)

	// Assert HTTP status
	status := response.Code
	if status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

}
