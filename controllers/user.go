package controllers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/context"
	"github.com/piyush7833/Chat-Api/functions"
	"github.com/piyush7833/Chat-Api/helpers"
	"github.com/piyush7833/Chat-Api/types"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	var user types.UpdateUserType
	err := helpers.GetBodyInJson(r, &user)
	if err != nil {
		helpers.Error(w, 500, "Error in retriving body")
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
	page := queryValues.Get("page")
	i, _ := strconv.Atoi(page)
	columns := queryValues["columns"]
	var selectedColumns = []string{}

	if len(columns) <= 0 {
		selectedColumns = append(selectedColumns, "id", "name", "username", "email", "phone", "image", "createdAt")
	} else {
		selectedColumns = append(selectedColumns, strings.Split(columns[0], ",")...)
	}
	res, error := functions.GetAllUser(i, selectedColumns)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "User fetched successfully")
}

func GetParticularUser(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")
	userName := queryValues.Get("userName")
	if id == "" && userName == "" {
		JwtData := context.Get(r, "userId").(*helpers.Claims)
		id = JwtData.Id
	}

	columns := queryValues["columns"]
	var selectedColumns = []string{}
	if len(columns) <= 0 {
		selectedColumns = append(selectedColumns, "id", "name", "username", "email", "phone", "image", "createdAt")
	} else {
		selectedColumns = append(selectedColumns, strings.Split(columns[0], ",")...)
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
