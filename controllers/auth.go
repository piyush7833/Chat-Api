package controllers

import (
	"net/http"

	"github.com/piyush7833/Chat-Api/functions"
	"github.com/piyush7833/Chat-Api/helpers"
	"github.com/piyush7833/Chat-Api/types"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	var user types.SignupType
	err := helpers.GetBodyInJson(r, &user)
	if helpers.CheckNilErr(err, "Error in retriving body", w) {
		return
	}
	res, error := functions.SignUp(user)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 201, res, "User created successfully")
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var user types.SignInType
	err := helpers.GetBodyInJson(r, &user)
	if helpers.CheckNilErr(err, "Error in retriving body", w) {
		return
	}
	res, error := functions.Signin(user, w)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "User logged in successfully")
}

func VerifyUser(w http.ResponseWriter, r *http.Request) {

}
func RecoverPassword(w http.ResponseWriter, r *http.Request) {

}
func ChangePassword(w http.ResponseWriter, r *http.Request) {

}
