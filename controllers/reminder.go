package controllers

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/piyush7833/Chat-Api/config"
	"github.com/piyush7833/Chat-Api/functions"
	"github.com/piyush7833/Chat-Api/helpers"
	"github.com/piyush7833/Chat-Api/types"
)

func CreateReminder(w http.ResponseWriter, r *http.Request) {
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	var reminder types.CreateReminderType
	err := helpers.GetBodyInJson(r, &reminder)
	if err != nil {
		helpers.Error(w, 500, err.Error())
		return
	}
	// fmt.Println("Reminder:", reminder)
	reminder.SenderId = &JwtData.Id
	res, error := functions.CreateReminder(reminder)

	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 201, res, "Reminder created successfully")
}

func GetParticularReminder(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	JwtData := context.Get(r, "userId").(*helpers.Claims)
	id := helpers.ProcessQuerryParams(queryValues["id"], []string{""}, "string").(string)
	// fmt.Println(id)
	page := helpers.ProcessQuerryParams(queryValues["page"], []string{"0"}, "number").(int)
	selectedColumns := helpers.ProcessQuerryParams(queryValues["columns"], config.GetDefaultReminderColumns(), "array").([]string)
	orderByColumns := helpers.ProcessQuerryParams(queryValues["orderBy"], []string{"createdAt"}, "array").([]string)
	isDesc := helpers.ProcessQuerryParams(queryValues["isDesc"], []string{"true"}, "bool").(bool)

	res, err := functions.GetParticularReminder(JwtData.Id, id, page, selectedColumns, orderByColumns, isDesc)
	if err.StatusCode != 0 {
		helpers.Error(w, err.StatusCode, err.Message)
		return
	}

	helpers.Success(w, 200, res, "Reminder fetched successfully")

}
func GetAllReminder(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()

	JwtData := context.Get(r, "userId").(*helpers.Claims)
	page := helpers.ProcessQuerryParams(queryValues["page"], []string{"0"}, "number").(int)
	reminder_type := helpers.ProcessQuerryParams(queryValues["type"], []string{"all"}, "string").(string)
	selectedColumns := helpers.ProcessQuerryParams(queryValues["columns"], config.GetDefaultReminderColumns(), "array").([]string)
	orderByColumns := helpers.ProcessQuerryParams(queryValues["orderBy"], []string{"createdAt"}, "array").([]string)
	isDesc := helpers.ProcessQuerryParams(queryValues["isDesc"], []string{"true"}, "bool").(bool)

	res, err := functions.GetAllReminder(JwtData.Id, page, reminder_type, selectedColumns, orderByColumns, isDesc)
	if err.StatusCode != 0 {
		helpers.Error(w, err.StatusCode, err.Message)
		return
	}

	helpers.Success(w, 200, res, "Reminder fetched successfully")
}
func UpdateReminder(w http.ResponseWriter, r *http.Request) {
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	queryValues := r.URL.Query()

	var reminder types.UpdateReminderType
	id := helpers.ProcessQuerryParams(queryValues["id"], []string{""}, "string").(string)

	err := helpers.GetBodyInJson(r, &reminder)
	if err != nil {
		helpers.Error(w, 500, err.Error())
		return
	}

	res, error := functions.UpdateReminder(reminder, id, JwtData.Id)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "Reminder updated successfully")
}

func DeleteReminder(w http.ResponseWriter, r *http.Request) {
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	queryValues := r.URL.Query()
	id := helpers.ProcessQuerryParams(queryValues["id"], []string{""}, "string").(string)

	res, error := functions.DeleteReminder(id, JwtData.Id)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "Reminder deleted successfully")
}
