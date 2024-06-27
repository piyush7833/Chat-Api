package routes

import (
	"github.com/gorilla/mux"
	"github.com/piyush7833/Chat-Api/controllers"
)

func ReminderRoutes(router *mux.Router, protectedRouter *mux.Router) {
	// fmt.Println("Reminder routes")
	protectedRouter.HandleFunc("/create", controllers.CreateReminder).Methods("POST")
	protectedRouter.HandleFunc("/update", controllers.UpdateReminder).Methods("PATCH")
	protectedRouter.HandleFunc("/get", controllers.GetParticularReminder).Methods("GET")
	protectedRouter.HandleFunc("/get-all", controllers.GetAllReminder).Methods("GET")
	protectedRouter.HandleFunc("/delete", controllers.DeleteReminder).Methods("DELETE")
	// router.HandleFunc()
}
