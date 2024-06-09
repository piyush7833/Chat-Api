package test

import "testing"

func TestGetAllUser(t *testing.T) {
	TestGetRequest(t, "/api/protected/user/get-all?page=0", 200, "Users retrieved successfully")
}
