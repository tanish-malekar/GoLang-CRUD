package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"
)

type User struct {
	Id       string
	Name     string
	Email    string
	Password string
}

type Post struct {
	Id        string
	Caption   string
	ImageURL  string
	Timestamp time.Time
	UserId    string
}

type userHandler struct {
	sync.Mutex
	client *mongo.Client
}

type postHandler struct {
	sync.Mutex
	client *mongo.Client
}

func (h *userHandler) postUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		var user User
		err = json.Unmarshal(bodyBytes, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		//encrypting the password
		user.Password = string(encrypt([]byte(user.Password), "Key"))

		h.Lock()
		coll := h.client.Database("appointyTaskDB").Collection("user")
		result, err := coll.InsertOne(context.TODO(), user)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
		w.Write([]byte("User successfully added"))
		h.Unlock()
	}
}

func (h *userHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	id := parts[2]

	h.Lock()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	coll := h.client.Database("appointyTaskDB").Collection("user")
	var user User
	if err := coll.FindOne(ctx, bson.M{"id": id}).Decode(&user); err != nil {
		log.Fatal(err)
	}
	h.Unlock()

	//decrypting the password
	user.Password = string(decrypt([]byte(user.Password), "Key"))

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.Write(jsonBytes)

}

func (h *postHandler) postPost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		bodyBytes, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		var post Post
		err = json.Unmarshal(bodyBytes, &post)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		post.Timestamp = time.Now()
		h.Lock()
		coll := h.client.Database("appointyTaskDB").Collection("post")
		result, err := coll.InsertOne(context.TODO(), post)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
		fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
		w.Write([]byte("Post successfully added"))
		h.Unlock()
	}
}

func (h *postHandler) getPostByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.String(), "/")
	id := parts[2]

	h.Lock()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	coll := h.client.Database("appointyTaskDB").Collection("post")
	var post bson.M
	if err := coll.FindOne(ctx, bson.M{"id": id}).Decode(&post); err != nil {
		log.Fatal(err)
	}
	h.Unlock()

	jsonBytes, err := json.Marshal(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.Write(jsonBytes)

}

func (h *postHandler) getPostsByUserID(w http.ResponseWriter, r *http.Request) {
	// /posts/users/<Id here>?limit=10&offset=10
	parts := strings.Split(r.URL.String(), "/")
	id := strings.Split(parts[3], "?")[0]

	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	offset, _ := strconv.Atoi(r.URL.Query().Get("offset"))

	options := options.Find()
	options.SetSkip(int64(offset))
	options.SetLimit(int64(limit))

	h.Lock()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	coll := h.client.Database("appointyTaskDB").Collection("post")
	cursor, err := coll.Find(ctx, bson.M{"userid": id}, options)
	h.Unlock()
	if err != nil {
		log.Fatal(err)
	}
	var posts []Post
	if err = cursor.All(ctx, &posts); err != nil {
		panic(err)
	}
	jsonBytes, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
	w.Header().Add("content-type", "application/json")
	w.Write(jsonBytes)

}

func newUserHandler(client *mongo.Client) *userHandler {
	return &userHandler{
		client: client,
	}
}

func newPostHandler(client *mongo.Client) *postHandler {
	return &postHandler{
		client: client,
	}
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}

func encrypt(data []byte, passphrase string) []byte {
	block, _ := aes.NewCipher([]byte(createHash(passphrase)))
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext
}

func decrypt(data []byte, passphrase string) []byte {
	key := []byte(createHash(passphrase))
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return plaintext
}

func main() {
	//establishing connection to the database
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
	postHandler := newPostHandler(client)

	http.HandleFunc("/users", userHandler.postUser)
	http.HandleFunc("/users/", userHandler.getUserByID)
	http.HandleFunc("/posts", postHandler.postPost)
	http.HandleFunc("/posts/", postHandler.getPostByID)
	http.HandleFunc("/posts/users/", postHandler.getPostsByUserID)
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
