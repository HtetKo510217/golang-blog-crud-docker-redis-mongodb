package router

import (
	"github.com/gorilla/mux"
	"github.com/htetko/go-with-redis-mongodb/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	// posts
	router.HandleFunc("/api/v1/posts", controller.GetAllPosts).Methods("GET")
	router.HandleFunc("/api/v1/post", controller.CreatePost).Methods("POST")
	router.HandleFunc("/api/v1/post/{id}", controller.GetPost).Methods("GET")
	router.HandleFunc("/api/v1/post/{id}", controller.UpdatePost).Methods("PUT")
	router.HandleFunc("/api/v1/post/{id}", controller.DeletePost).Methods("DELETE")
	router.HandleFunc("/api/v1/posts", controller.DeleteAllPosts).Methods("DELETE")
	// users
	router.HandleFunc("/api/v1/user", controller.CreateUser).Methods("POST")
	router.HandleFunc("/api/v1/user/{id}", controller.GetUser).Methods("GET")
	router.HandleFunc("/api/v1/user/{id}", controller.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/v1/user/{id}", controller.DeleteUser).Methods("DELETE")
	router.HandleFunc("/api/v1/users", controller.GetAllUsers).Methods("GET")

	// comments
	router.HandleFunc("/api/v1/post/{id}/comment", controller.CreateComment).Methods("POST")
	router.HandleFunc("/api/v1/post/{id}/comment", controller.GetCommentsByPostID).Methods("GET")

	return router
}
