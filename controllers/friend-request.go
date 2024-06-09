package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/context"
	"github.com/piyush7833/Chat-Api/functions"
	"github.com/piyush7833/Chat-Api/helpers"
	"github.com/piyush7833/Chat-Api/types"
)

func CreateFR(w http.ResponseWriter, r *http.Request) {
	var friendRequest types.FriendRequestType
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	userId := JwtData.Id
	err := helpers.GetBodyInJson(r, &friendRequest)
	if err != nil {
		helpers.Error(w, 500, "Error in retriving body")
		return
	}
	res, error := functions.CreateFriendRequest(userId, *friendRequest.ReceiverId)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "Friend request sent successfully")
}

func DeleteFR(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	userId := JwtData.Id
	res, error := functions.DeleteFriendRequest(id, userId)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "Friend request deleted successfully")
}

func GetAllFR(w http.ResponseWriter, r *http.Request) {
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	id := JwtData.Id
	queryValues := r.URL.Query()
	reqType := queryValues.Get("type")
	page := queryValues.Get("page")
	i, _ := strconv.Atoi(page)
	var res any
	var error types.ErrorType
	if reqType == "sent" {
		res, error = functions.GetAllFriendRequestSentByUser(id, i)
	} else {
		res, error = functions.GetAllFriendRequestRecievedByUser(id, i)
	}

	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "Friend request fetched successfully")
}

func GetParticularFR(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	userId := JwtData.Id
	res, error := functions.GetParticularFriendRequest(id, userId)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "Friend request fetched successfully")
}

func UpdateFRStatus(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	id := queryValues.Get("id")
	status := queryValues.Get("status")
	JwtData := context.Get(r, "userId").(*helpers.Claims)
	userId := JwtData.Id
	res, error := functions.UpdateFriendRequestStatus(id, status, userId)
	if error.StatusCode != 0 {
		helpers.Error(w, error.StatusCode, error.Message)
		return
	}
	helpers.Success(w, 200, res, "Friend request fetched successfully")
}
