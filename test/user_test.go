package test

import (
	"fmt"
	"testing"

	"github.com/piyush7833/Chat-Api/config"
	"github.com/piyush7833/Chat-Api/services"
)

// get all user
func TestGetAllUser(t *testing.T) {
	TestGetRequest(t, "/api/protected/user/get-all?page=0", 200, "Users fetched successfully", true)
}
func TestGetAllUserErrorUnAuthorized(t *testing.T) {
	TestGetRequest(t, "/api/protected/user/get-all?page=0", 401, "login required", false)
}
func TestGetAllUserErrorInvalidColumns(t *testing.T) {
	TestGetRequest(t, "/api/protected/user/get-all?page=0&columns=invalid", 400, "no valid columns selected", true)
}

// get user by id test
func TestGetUserById(t *testing.T) {
	var whereClause = "username = " + "'test_user'"
	id, err := services.GetRows("users", 0, []string{"id"}, config.GetValidUserColumns(), &whereClause, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf("Error in fetching user id")
		return
	}
	url := fmt.Sprintf("/api/protected/user/get?id=%v", id[0]["id"])
	TestGetRequest(t, url, 200, "Users fetched successfully", true)
}

func TestGetUserByIdErrorInvalidColumns(t *testing.T) {
	var whereClause = "username = " + "'test_user'"
	id, err := services.GetRows("users", 0, []string{"id"}, config.GetValidUserColumns(), &whereClause, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf("Error in fetching user id")
		return
	}
	url := fmt.Sprintf("/api/protected/user/get?id=%v&columns=invalid", id[0]["id"])
	TestGetRequest(t, url, 400, "no valid columns selected", true)
}
func TestGetUserByIdErrorUnauthorized(t *testing.T) {
	var whereClause = "username = " + "'test_user'"
	id, err := services.GetRows("users", 0, []string{"id"}, config.GetValidUserColumns(), &whereClause, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf("Error in fetching user id")
		return
	}
	url := fmt.Sprintf("/api/protected/user/get?id=%v", id[0]["id"])
	TestGetRequest(t, url, 401, "login required", false)

}
func TestGetUserByIdErrorUserNotFound(t *testing.T) {
	url := fmt.Sprintf("/api/protected/user/get?id=%v", "e82fb13f-a118-4eac-98c5-629ac01ef8a2") //id which is not in db
	TestGetRequest(t, url, 404, "No users found", true)
}

// get user by username test
func TestGetUserByUsername(t *testing.T) {
	TestGetRequest(t, "/api/protected/user/get?username=test_user", 200, "Users fetched successfully", true)
}
func TestGetUserByUsernameErrorUserNotFound(t *testing.T) {
	TestGetRequest(t, "/api/protected/user/get?username=user_not_found", 404, "No users found", true)
}
func TestGetUserByUsernameErrorInvalidColumns(t *testing.T) {
	TestGetRequest(t, "/api/protected/user/get?username=test_user&columns=invalid", 400, "no valid columns selected", true)
}
func TestGetUserByUsernameErrorUnauthorized(t *testing.T) {
	TestGetRequest(t, "/api/protected/user/get?username=test_user", 401, "login required", false)
}

// get authenticated user
func TestGetAuthenticatedUser(t *testing.T) {
	TestGetRequest(t, "/api/protected/user/get", 200, "Users fetched successfully", true)
}
func TestGetAuthenticatedUserErrorUnauthorized(t *testing.T) {
	TestGetRequest(t, "/api/protected/user/get", 401, "login required", false)
}
func TestGetAuthenticatedUserErrorInvalidColumns(t *testing.T) {
	TestGetRequest(t, "/api/protected/user/get?columns=invalid", 400, "no valid columns selected", true)
}

// update user
func TestUpdateUser(t *testing.T) {
	userData := `{
        "name": "updated_test_user"
    }`
	TestPatchRequest(t, userData, "/api/protected/user/update", 200, "User updated successfully", true)
}
func TestUpdateUserErrorUnauthorized(t *testing.T) {
	userData := `{
        "name": "updated_test_user"
    }`
	TestPatchRequest(t, userData, "/api/protected/user/update", 401, "login required", false)
}
func TestUpdateUserErrorInvalidColumns(t *testing.T) {
	userData := `{
        "username": "updated_test_user"
    }`
	TestPatchRequest(t, userData, "/api/protected/user/update", 400, "no valid columns selected", true)
}

// delete user
func TestDeleteUser(t *testing.T) {
	TestDeleteRequest(t, "/api/protected/user/delete", 200, "User deleted successfully", true)
}
func TestDeleteUserErrorUnauthorized(t *testing.T) {
	TestDeleteRequest(t, "/api/protected/user/delete", 401, "login required", false)
}
