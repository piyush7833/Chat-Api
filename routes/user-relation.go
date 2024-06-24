package routes

import (
	"github.com/gorilla/mux"
	"github.com/piyush7833/Chat-Api/controllers"
)

func UserRelationRoutes(router *mux.Router, protectedRouter *mux.Router) {

	protectedRouter.HandleFunc("/create", controllers.CreateUserRelations).Methods("POST")
	protectedRouter.HandleFunc("/update", controllers.UpdateUserRelations).Methods("PATCH")
	protectedRouter.HandleFunc("/get", controllers.GetParticularUserRelations).Methods("GET")
	protectedRouter.HandleFunc("/get-all", controllers.GetAllUserRelations).Methods("GET")
	protectedRouter.HandleFunc("/delete", controllers.DeleteUserRelations).Methods("DELETE")
	// router.HandleFunc()
}
