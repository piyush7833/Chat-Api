package test

import (
	"fmt"
	"testing"

	"github.com/piyush7833/Chat-Api/config"
	"github.com/piyush7833/Chat-Api/services"
)

// create user relation tests
func TestCreateUserRelation(t *testing.T) {
	var whereClause = "username = " + "'related_user'"
	res, err := services.GetRows("users", 0, []string{"id"}, config.GetValidUserColumns(), &whereClause, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	userRelation := `{
				"relatedUserId":  "` + fmt.Sprintf("%v", res[0]["id"]) + `"
			}`
	// fmt.Println(userRelation)
	TestPostRequest(t, userRelation, "/api/protected/ur/create", 201, "User relation created successfully", true)
}
func TestCreateUserRelationErrorMissingField(t *testing.T) {
	userRelation := `{
			}`
	// fmt.Println(userRelation)
	TestPostRequest(t, userRelation, "/api/protected/ur/create", 500, "Missing required field", true)
}
func TestCreateUserRelationErrorDuplicateRelation(t *testing.T) {
	var whereClause = "username = 'related_user'"
	res, err := services.GetRows("users", 0, []string{"id"}, config.GetValidUserColumns(), &whereClause, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	userRelation := `{
				"relatedUserId":  "` + fmt.Sprintf("%v", res[0]["id"]) + `"
			}`
	// fmt.Println(userRelation)
	TestPostRequest(t, userRelation, "/api/protected/ur/create", 400, "Relation already exists", true)
}
func TestCreateUserRelationErrorRelationAlreadyExist(t *testing.T) { //current user log out and new user login and try to create relation with same user //configure this test in last
	var whereClause = "username = 'related_user'"
	res, err := services.GetRows("users", 0, []string{"id"}, config.GetValidUserColumns(), &whereClause, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	userRelation := `{
				"relatedUserId":  "` + fmt.Sprintf("%v", res[0]["id"]) + `"
			}`
	// fmt.Println(userRelation)
	TestPostRequest(t, userRelation, "/api/protected/ur/create", 400, "Relation already exists", true)
}
func TestCreateUserRelationErrorUnauthorized(t *testing.T) {
	var whereClause = "username = " + "'related_user'"
	res, err := services.GetRows("users", 0, []string{"id"}, config.GetValidUserColumns(), &whereClause, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	userRelation := `{
				"relatedUserId":  "` + fmt.Sprintf("%v", res[0]["id"]) + `"
			}`
	// fmt.Println(userRelation)
	TestPostRequest(t, userRelation, "/api/protected/ur/create", 401, "Login Required", false)
}
func TestCreateUserRelationErrorInvalidRelatedUserId(t *testing.T) {
	userRelation := `{
				"relatedUserId":  "randomId"
			}`
	// fmt.Println(userRelation)
	TestPostRequest(t, userRelation, "/api/protected/ur/create", 404, "Related user does not exists", true)
}

// get all user relation tests
func TestGetAllUserRelation(t *testing.T) {
	TestGetRequest(t, "/api/protected/ur/get-all", 200, "User relations fetched successfully", true)
}
func TestGetAllUserRelationParticularStatus(t *testing.T) {
	TestGetRequest(t, "/api/protected/ur/get-all?status=pending", 200, "User relations fetched successfully", true)
}
func TestGetAllUserRelationParticularType(t *testing.T) {
	TestGetRequest(t, "/api/protected/ur/get-all?type=sent", 200, "User relations fetched successfully", true)
}
func TestGetAllUserRelationParticularTypeParticularStatus(t *testing.T) {
	TestGetRequest(t, "/api/protected/ur/get-all?type=sent&status=pending", 200, "User relations fetched successfully", true)
}
func TestGetAllUserRelationNoneFound(t *testing.T) {
	TestGetRequest(t, "/api/protected/ur/get-all", 200, "no userReltion found", true)
}
func TestGetAllUserRelationErrorUnauthorized(t *testing.T) {
	TestGetRequest(t, "/api/protected/ur/get-all", 401, "login required", false)
}
func TestGetAllUserRelationErrorInvalidColumns(t *testing.T) {
	TestGetRequest(t, "/api/protected/ur/get-all?columns=invalid", 400, "no valid columns selected", true)
}
func TestGetAllUserRelationErrorInvalidStatus(t *testing.T) {
	TestGetRequest(t, "/api/protected/ur/get-all?status=invalid", 400, "invalid status", true)
}
func TestGetAllUserRelationErrorInvalidType(t *testing.T) {
	TestGetRequest(t, "/api/protected/ur/get-all?type=invalid", 400, "invalid type", true)
}

// get particular user relation tests
func TestGetParticularUserRelation(t *testing.T) {
	res, err := services.GetRows("userRelation", 0, []string{"id"}, config.GetValidUserColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	url := fmt.Sprintf("/api/protected/ur/get?id=%v", res[0]["id"])
	TestGetRequest(t, url, 200, "User relation fetched successfully", true)
}
func TestGetParticularUserRelationErrorUserRelationNotFound(t *testing.T) {
	url := fmt.Sprintf("/api/protected/ur/get?id=%v", "randomId")
	TestGetRequest(t, url, 403, "User relation does not exists or it is not associated with you", true)
}
func TestGetParticularUserRelationErrorInvalidColumn(t *testing.T) {
	res, err := services.GetRows("userRelation", 0, []string{"id"}, config.GetValidUserColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	url := fmt.Sprintf("/api/protected/ur/get?id=%v&columns=invalid", res[0]["id"])
	TestGetRequest(t, url, 400, "no valid columns selected", true)
}
func TestGetParticularUserRelationErrorUnauthorized(t *testing.T) {
	res, err := services.GetRows("userRelation", 0, []string{"id"}, config.GetValidUserColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	url := fmt.Sprintf("/api/protected/ur/get?id=%v", res[0]["id"])
	TestGetRequest(t, url, 401, "User relation fetched successfully", false)
}
func TestGetParticularUserRelationErrorNotAllowed(t *testing.T) { //can be test when logged in using third user //test in the last
	res, err := services.GetRows("userRelation", 0, []string{"id"}, config.GetValidUserColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	url := fmt.Sprintf("/api/protected/ur/get?id=%v", res[0]["id"])
	TestGetRequest(t, url, 403, "Only user and related user are allowed to view this relation", true)
}

// update user relation tests
func TestUpdateUserRelation(t *testing.T) { //their is only one user relation exist which is sent by logged in user and sender can not change status so this will be tested in last
	res, err := services.GetRows("userRelation", 0, []string{"id"}, config.GetValidUserColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	url := fmt.Sprintf("/api/protected/ur/update?id=%v", res[0]["id"])
	userRelation := `{
		"status":  "friends"
	}`
	TestPatchRequest(t, userRelation, url, 200, "User relation updated successfully", true)
}
func TestUpdateUserRelationErrorStatusNotProvided(t *testing.T) {
	res, err := services.GetRows("userRelation", 0, []string{"id"}, config.GetValidUserColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	url := fmt.Sprintf("/api/protected/ur/update?id=%v", res[0]["id"])
	userRelation := `{
		"invalid":  "friends"
	}`
	TestPatchRequest(t, userRelation, url, 400, "Status is required", true)
}
func TestUpdateUserRelationErrorNotFound(t *testing.T) {
	url := fmt.Sprintf("/api/protected/ur/update?id=%v", "randomId")
	userRelation := `{
		"invalid":  "friends"
	}`
	TestPatchRequest(t, userRelation, url, 403, "User relation does not exists or it is not associated with you", true)
}
func TestUpdateUserRelationErrorUnauthorized(t *testing.T) {
	res, err := services.GetRows("userRelation", 0, []string{"id"}, config.GetValidUserColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	url := fmt.Sprintf("/api/protected/ur/update?id=%v", res[0]["id"])
	userRelation := `{
		"status":  "friends"
	}`
	TestPatchRequest(t, userRelation, url, 401, "login required", false)
}
func TestUpdateUserRelationErrorNotAllowed(t *testing.T) { //sender can not update status to friends
	res, err := services.GetRows("userRelation", 0, []string{"id"}, config.GetValidUserColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	url := fmt.Sprintf("/api/protected/ur/update?id=%v", res[0]["id"])
	userRelation := `{
		"status":  "friends"
	}`
	TestPatchRequest(t, userRelation, url, 403, "Only related user can update status to friends", true)
}
func TestUpdateUserRelationErrorNotAllowedToRevertToPendingStatus(t *testing.T) {
	res, err := services.GetRows("userRelation", 0, []string{"id"}, config.GetValidUserColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	url := fmt.Sprintf("/api/protected/ur/update?id=%v", res[0]["id"])
	userRelation := `{
		"status":  "pending"
	}`
	TestPatchRequest(t, userRelation, url, 406, "User relation status is already updated to friends or blocked, you can't update status to pending again", true)
}

// test for deleting users
func TestDeleteUserRelation(t *testing.T) {
	res, err := services.GetRows("userRelation", 0, []string{"id"}, config.GetValidUserColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	url := fmt.Sprintf("/api/protected/ur/delete?id=%v", res[0]["id"])
	TestDeleteRequest(t, url, 200, "User relation deleted successfully", true)
}
func TestDeleteUserRelationErrorNotFound(t *testing.T) {
	url := fmt.Sprintf("/api/protected/ur/delete?id=%v", "randomId")
	TestDeleteRequest(t, url, 403, "userRelation not found", true)
}
func TestDeleteUserRelationErrorUnauthorised(t *testing.T) {
	res, err := services.GetRows("userRelation", 0, []string{"id"}, config.GetValidUserColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	url := fmt.Sprintf("/api/protected/ur/delete?id=%v", res[0]["id"])
	TestDeleteRequest(t, url, 401, "login required", false)
}
func TestDeleteUserRelationErrorNotAllowed(t *testing.T) { //can be test when logged in using third user //test in the last
	res, err := services.GetRows("userRelation", 0, []string{"id"}, config.GetValidUserColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	url := fmt.Sprintf("/api/protected/ur/delete?id=%v", res[0]["id"])
	TestDeleteRequest(t, url, 403, "Only user and related user are allowed to delete this request", true)
}
