package routes

import (
	"github.com/gorilla/mux"
	"github.com/piyush7833/Chat-Api/controllers"
)

func FriendRequestRoutes(router *mux.Router, protectedRouter *mux.Router) {

	protectedRouter.HandleFunc("/create", controllers.CreateFR).Methods("POST")
	protectedRouter.HandleFunc("/update", controllers.UpdateFRStatus).Methods("PATCH")
	protectedRouter.HandleFunc("/get", controllers.GetParticularFR).Methods("GET")
	protectedRouter.HandleFunc("/get-all", controllers.GetAllFR).Methods("GET")
	protectedRouter.HandleFunc("/delete", controllers.DeleteFR).Methods("DELETE")
	// router.HandleFunc()
}
