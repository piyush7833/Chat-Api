package routes

import (
	"github.com/gorilla/mux"
	"github.com/piyush7833/Chat-Api/controllers"
)

func AuthRoutes(router *mux.Router) {
	router.HandleFunc("/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/signin", controllers.SignIn).Methods("POST")
	router.HandleFunc("/verify", controllers.VerifyUser).Methods("POST")
	router.HandleFunc("/recover", controllers.RecoverPassword).Methods("POST")
	router.HandleFunc("/change", controllers.ChangePassword).Methods("POST")
}
