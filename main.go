package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	"github.com/htetko/go-with-redis-mongodb/initializers"
	"github.com/htetko/go-with-redis-mongodb/router"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}
func main() {
	router := router.Router()
	corsHandler := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	http.Handle("/", corsHandler(router))

	fmt.Println("starting the applition ..")
	log.Fatal(http.ListenAndServe(":3000", nil))
	fmt.Println("server is running on port 3000")

}
