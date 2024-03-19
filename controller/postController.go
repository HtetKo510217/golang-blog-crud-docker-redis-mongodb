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

// helper function
func respondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	post := model.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	insertResult, err := initializers.Client.Database("blog").Collection("posts").InsertOne(ctx, post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	// Set posts in Redis with the new post
	rdb := initializers.RedisClient
	postMap := bson.M{
		"_id":     insertResult.InsertedID,
		"user_id": post.UserID,
		"title":   post.Title,
		"content": post.Content,
	}
	val, err := rdb.Get(context.Background(), "posts").Result()
	var posts []bson.M
	if err == nil {
		err = json.Unmarshal([]byte(val), &posts)
		if err != nil {
			panic(err)
		}
	}
	posts = append(posts, postMap)
	jsonData, err := json.Marshal(posts)
	if err != nil {
		panic(err)
	}
	err = rdb.Set(context.Background(), "posts", jsonData, 0).Err()
	if err != nil {
		panic(err)
	}

	respondWithJSON(w, http.StatusCreated, post)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	post := model.Post{}

	// Create a filter for the query
	filter := bson.M{"_id": id}

	// Execute the query
	err := initializers.Client.Database("blog").Collection("posts").FindOne(ctx, filter).Decode(&post)
	if err != nil {
		// Log the error for debugging
		fmt.Println("Error retrieving post:", err)

		// Return a meaningful error response
		w.WriteHeader(http.StatusNotFound) // Use 404 status code for "not found"
		w.Write([]byte(fmt.Sprintf("Post not found for id %s", params["id"])))
		return
	}

	// If the query is successful, return the post data
	respondWithJSON(w, http.StatusOK, post)

}

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	posts := []model.Post{}

	// Check if posts exist in Redis
	rdb := initializers.RedisClient
	val, err := rdb.Get(context.Background(), "posts").Result()
	if err == nil {
		// Posts exist in Redis, return them
		err = json.Unmarshal([]byte(val), &posts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error unmarshalling Redis data: %v", err)))
			return
		}
		respondWithJSON(w, http.StatusOK, posts)
		return
	}

	// Fetch posts from MongoDB
	cursor, err := initializers.Client.Database("blog").Collection("posts").Find(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error fetching posts from MongoDB: %v", err)))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post model.Post
		err := cursor.Decode(&post)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Error decoding post: %v", err)))
			return
		}
		posts = append(posts, post)
	}

	// Set posts in Redis after serializing to JSON
	jsonData, err := json.Marshal(posts)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error marshalling posts data: %v", err)))
		return
	}
	err = rdb.Set(context.Background(), "posts", jsonData, 0).Err()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Error setting posts in Redis: %v", err)))
		return
	}

	respondWithJSON(w, http.StatusOK, posts)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	post := model.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	update := bson.M{"$set": post} // Use $set operator to update specific fields
	_, err = initializers.Client.Database("blog").Collection("posts").UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	// response if updated
	respondWithJSON(w, http.StatusOK, post)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// Delete the post from MongoDB
	_, err := initializers.Client.Database("blog").Collection("posts").DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}

	// Update the posts in Redis by removing the deleted post
	rdb := initializers.RedisClient
	val, err := rdb.Get(ctx, "posts").Result()
	var posts []bson.M
	if err == nil {
		err = json.Unmarshal([]byte(val), &posts)
		if err != nil {
			panic(err)
		}
		// Find and remove the deleted post from the posts
		var updatedPosts []bson.M
		for _, post := range posts {
			postIDHex := post["_id"].(string)
			postID, _ := primitive.ObjectIDFromHex(postIDHex)
			if postID != id { // Compare the ObjectID with the ID of the post to be deleted
				updatedPosts = append(updatedPosts, post)
			}
		}
		// Update the modified posts back in Redis
		jsonData, err := json.Marshal(updatedPosts)
		if err != nil {
			panic(err)
		}
		err = rdb.Set(ctx, "posts", jsonData, 0).Err()
		if err != nil {
			panic(err)
		}
	}

	// Respond with the updated list of posts from Redis
	updatedVal, _ := rdb.Get(ctx, "posts").Result()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(updatedVal))

}

func DeleteAllPosts(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	_, err := initializers.Client.Database("blog").Collection("posts").DeleteMany(ctx, bson.M{})
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Error: %v", err)))
		return
	}
	// clear redis
	err = initializers.RedisClient.Del(ctx, "posts").Err()
	if err != nil {
		fmt.Println("Error clearing redis cache:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
