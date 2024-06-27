package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/piyush7833/Chat-Api/docs"
	"github.com/piyush7833/Chat-Api/helpers"
	"github.com/piyush7833/Chat-Api/middlewares"
	"github.com/piyush7833/Chat-Api/routes"
	"github.com/piyush7833/Chat-Api/services"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Chat-API
// @version 1.0
// @description This is a sample swagger documentation for chat-api
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @host localhost:8080
// @BasePath /
func main() {
	if os.Getenv("ENVIRONMENT") == "" || os.Getenv("ENVIRONMENT") == "development" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatal(err)
		}
	}

	services.ConnectDb()
	// scripts.CreateTables()
	router := InitRoutes()

	log.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	protectedRouter := router.PathPrefix("/api/protected").Subrouter()
	protectedRouter.Use(middlewares.AuthMiddleware)

	// Routes
	routes.AuthRoutes(router.PathPrefix("/api").Subrouter())
	routes.UserRoutes(router.PathPrefix("/api/user").Subrouter(), protectedRouter.PathPrefix("/user").Subrouter())
	routes.UserRelationRoutes(router.PathPrefix("/api/ur").Subrouter(), protectedRouter.PathPrefix("/ur").Subrouter())
	routes.ReminderRoutes(router.PathPrefix("/api/reminder").Subrouter(), protectedRouter.PathPrefix("/reminder").Subrouter())

	//graphql
	// srv := InitGraphql()
	// router.Handle("/query", srv)
	// router.Handle("/graphql", playground.Handler("GraphQL playground", "/query"))

	// Health check
	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		helpers.Success(w, http.StatusOK, nil, "Server is healthy and running")
	}).Methods("GET")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helpers.Success(w, http.StatusOK, nil, "Welcome to chat-api")
	}).Methods("GET")

	// Swagger documentation
	router.PathPrefix("/documentation/").Handler(httpSwagger.WrapHandler)

	return router
}

// func InitGraphql() *handler.Server {
// 	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
// 	return srv
// }
