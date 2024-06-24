package controllers

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/piyush7833/Chat-Api/config"
	"github.com/piyush7833/Chat-Api/functions"
	"github.com/piyush7833/Chat-Api/helpers"
	"github.com/piyush7833/Chat-Api/types"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	var user types.UpdateUserType
	err := helpers.GetBodyInJson(r, &user)
	if err != nil {
		helpers.Error(w, 500, err.Error())
		return
	}
	res, error := functions.UpdateUser(user, JwtData.Id)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "User updated successfully")
}

func GetAllUser(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	page := helpers.ProcessQuerryParams(queryValues["page"], []string{"1"}, "number").(int)
	selectedColumns := helpers.ProcessQuerryParams(queryValues["columns"], config.GetUserDefaultColumns(), "array").([]string)
	orderByColumns := helpers.ProcessQuerryParams(queryValues["orderBy"], []string{"createdAt"}, "array").([]string)
	isDesc := helpers.ProcessQuerryParams(queryValues["isDesc"], []string{"true"}, "bool").(bool)

	res, error := functions.GetAllUser(page, selectedColumns, orderByColumns, isDesc)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "User fetched successfully")
}

func GetParticularUser(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	id := helpers.ProcessQuerryParams(queryValues["id"], []string{""}, "string").(string)
	userName := helpers.ProcessQuerryParams(queryValues["username"], []string{""}, "string").(string)
	selectedColumns := helpers.ProcessQuerryParams(queryValues["columns"], config.GetUserDefaultColumns(), "array").([]string)

	if id == "" && userName == "" {
		JwtData := context.Get(r, "userId").(*helpers.Claims)
		id = JwtData.Id
	}

	var res any
	var error types.ErrorType

	if userName != "" {
		res, error = functions.GetUsersByUsername(userName, selectedColumns)
	} else {
		res, error = functions.GetParticularUserById(id, selectedColumns)
	}

	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "User fetched successfully")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	id := JwtData.Id
	res, error := functions.DeleteUser(id)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "User deleted successfully")
}
