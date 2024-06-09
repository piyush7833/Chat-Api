package test

import (
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/piyush7833/Chat-Api/middlewares"
	"github.com/piyush7833/Chat-Api/routes"
	"github.com/piyush7833/Chat-Api/scripts"
	"github.com/piyush7833/Chat-Api/services"
)

func Init() error {
	if os.Getenv("ENVIRONMENT") == "" || os.Getenv("ENVIRONMENT") == "test" {
		if err := godotenv.Load(".env.test"); err != nil {
			return err
		}
	}
	services.ConnectDb()
	scripts.CreateTables()
	return nil
}

func InitRoutes() (*mux.Router, *mux.Router) {
	router := mux.NewRouter()
	protectedRouter := router.PathPrefix("/api/protected").Subrouter()
	protectedRouter.Use(middlewares.AuthMiddleware)
	routes.AuthRoutes(router.PathPrefix("/api").Subrouter())
	routes.UserRoutes(router.PathPrefix("/api/user").Subrouter(), protectedRouter.PathPrefix("/user").Subrouter())
	routes.FriendRequestRoutes(router.PathPrefix("/api/fr").Subrouter(), protectedRouter.PathPrefix("/fr").Subrouter())
	return router, protectedRouter
}
