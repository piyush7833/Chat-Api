package controllers

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/piyush7833/Chat-Api/config"
	"github.com/piyush7833/Chat-Api/functions"
	"github.com/piyush7833/Chat-Api/helpers"
	"github.com/piyush7833/Chat-Api/types"
)

func CreateUserRelations(w http.ResponseWriter, r *http.Request) {
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	var relation types.CreateUserRelationType
	err := helpers.GetBodyInJson(r, &relation)
	if err != nil {
		helpers.Error(w, 500, err.Error())
		return
	}
	var senderId = JwtData.Id
	var status = "pending"
	if relation.RelatedUserId == nil {
		helpers.Error(w, 500, "ReceiverId is required")
		return
	}
	relation.UserId = &senderId
	relation.Status = &status
	res, error := functions.CreateUserRelation(relation)
	if error.StatusCode != 0 {
		// fmt.Println(error.StatusCode, error.Message, "error from fucntion")
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 201, res, "User relation created sent successfully")
}

func UpdateUserRelations(w http.ResponseWriter, r *http.Request) {
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	queryValues := r.URL.Query()
	id := queryValues.Get("id")
	var relation types.UpdateUserRelationType
	err := helpers.GetBodyInJson(r, &relation)
	if err != nil {
		helpers.Error(w, 500, err.Error())
		return
	}
	res, error := functions.UpdateUserRelation(relation, id, JwtData.Id)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "Relation updated successfully")
}

func GetAllUserRelations(w http.ResponseWriter, r *http.Request) {
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	queryValues := r.URL.Query()

	status := helpers.ProcessQuerryParams(queryValues["status"], []string{"all"}, "string").(string)
	relation_type := helpers.ProcessQuerryParams(queryValues["type"], []string{"all"}, "string").(string)
	page := helpers.ProcessQuerryParams(queryValues["page"], []string{"0"}, "number").(int)
	selectedColumns := helpers.ProcessQuerryParams(queryValues["columns"], config.GetDefaultUserRelationColumns(), "array").([]string)
	orderBy := helpers.ProcessQuerryParams(queryValues["orderBy"], []string{"createdAt"}, "array").([]string)
	isDesc := helpers.ProcessQuerryParams(queryValues["isDesc"], []string{"true"}, "bool").(bool)

	res, error := functions.GetAllUserRelations(status, page, selectedColumns, JwtData.Id, relation_type, orderBy, isDesc)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "Relations fetched successfully")
}

func GetParticularUserRelations(w http.ResponseWriter, r *http.Request) {
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	queryValues := r.URL.Query()

	id := helpers.ProcessQuerryParams(queryValues["id"], []string{""}, "string").(string)
	var selectedColumns = helpers.ProcessQuerryParams(queryValues["columns"], config.GetDefaultUserRelationColumns(), "array").([]string)

	res, error := functions.GetParticularUserRelation(id, selectedColumns, JwtData.Id)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "Relation fetched successfully")
}

func DeleteUserRelations(w http.ResponseWriter, r *http.Request) { //unrelate
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	queryValues := r.URL.Query()
	id := queryValues.Get("id")
	res, error := functions.DeleteUserRelations(id, JwtData.Id)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "Relation deleted successfully")
}
