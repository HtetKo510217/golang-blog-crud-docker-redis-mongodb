package controller

import (
	"context"
	"encoding/json"
	"log"
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
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	var comment model.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Fatal(err)
	}

	_, err = initializers.Client.Database("blog").Collection("comments").InsertOne(ctx, comment)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode(comment)
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Comment created successfully"))
}

func GetCommentsByPostID(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var comments []model.Comment
	cursor, err := initializers.Client.Database("blog").Collection("comments").Find(ctx, bson.M{"post_id": id})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var comment model.Comment
		err := cursor.Decode(&comment)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, comment)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(comments)
	w.WriteHeader(http.StatusOK)
}
