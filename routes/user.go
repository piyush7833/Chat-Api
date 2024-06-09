package routes

import (
	"github.com/gorilla/mux"
	"github.com/piyush7833/Chat-Api/controllers"
)

func UserRoutes(router *mux.Router, protectedRouter *mux.Router) {

	protectedRouter.HandleFunc("/update", controllers.UpdateUser).Methods("PATCH")
	protectedRouter.HandleFunc("/get", controllers.GetParticularUser).Methods("GET")
	protectedRouter.HandleFunc("/get-all", controllers.GetAllUser).Methods("GET")
	protectedRouter.HandleFunc("/delete", controllers.DeleteUser).Methods("DELETE")
	// router.HandleFunc()
}
