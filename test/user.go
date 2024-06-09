package test

import (
	"log"
)

func init() {
	err := Init()
	if err != nil {
		log.Fatal(err)
	}
}

// func TestGetAllUser(t *testing.T) {
// 	TestGetRequest(t, "/api/protected/user/get-all?page=0", 200, "Users retrieved successfully")
// }
