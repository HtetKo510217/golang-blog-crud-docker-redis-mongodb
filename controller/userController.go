package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/htetko/go-with-redis-mongodb/initializers"
	"github.com/htetko/go-with-redis-mongodb/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.ConnectToRedis()
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}
	_, err = initializers.Client.Database("blog").Collection("users").InsertOne(ctx, user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewDecoder(r.Body).Decode(&user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	err = initializers.Client.Database("blog").Collection("users").FindOne(ctx, id).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	users := []model.User{}

	cursor, err := initializers.Client.Database("blog").Collection("users").Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		post := model.User{}
		err := cursor.Decode(&post)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error: %v", err)))
			return
		}
		users = append(users, post)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	_, err = initializers.Client.Database("blog").Collection("users").UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": user})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	_, err := initializers.Client.Database("blog").Collection("users").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
