package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/piyush7833/Chat-Api/helpers"
	"github.com/piyush7833/Chat-Api/middlewares"
	"github.com/piyush7833/Chat-Api/routes"
	"github.com/piyush7833/Chat-Api/services"
)

func init() {

}
func main() {
	if os.Getenv("ENVIRONMENT") == "" || os.Getenv("ENVIRONMENT") == "development" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal(err)
		}
	}
	services.ConnectDb()
	router := mux.NewRouter()
	protectedRouter := router.PathPrefix("/api/protected").Subrouter()
	protectedRouter.Use(middlewares.AuthMiddleware)
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		helpers.Success(w, http.StatusOK, nil, "Server is healthy and running")
	}).Methods("GET")

	routes.AuthRoutes(router.PathPrefix("/api").Subrouter())
	routes.UserRoutes(router.PathPrefix("/api/user").Subrouter(), protectedRouter.PathPrefix("/user").Subrouter())
	routes.FriendRequestRoutes(router.PathPrefix("/api/fr").Subrouter(), protectedRouter.PathPrefix("/fr").Subrouter())
	// routes.MessageRoutes(router)

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
