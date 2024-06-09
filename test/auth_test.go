package test

import (
	"log"
	"net/http"
	"testing"
)

// signup
func init() {
	err := Init()
	if err != nil {
		log.Fatal(err)
	}
}

func TestCreateUser(t *testing.T) {
	userData := `{
        "username": "test_user",
        "name": "Test User",
        "email": "test@example.com",
        "password": "password",
        "phone": "+1234567890",
        "image": "https://example.com/avatar.png"
    }`
	TestPostRequest(t, userData, "/api/signup", 201, "User created successfully")
}

func TestDuplicateUserError(t *testing.T) {
	userData := `{
        "username": "test_user",
        "name": "Test User",
        "email": "test@example.com",
        "password": "password",
        "phone": "+1234567890",
        "image": "https://example.com/avatar.png"
    }`
	TestPostRequest(t, userData, "/api/signup", http.StatusForbidden, "User with given user name or email or phone already exist")

}
func TestMissingFieldErrorSignup(t *testing.T) {
	userData := `{
        "username": "test_user_2",
        "name": "Test User",
		"password": "password",
        "phone": "+1234567892",
        "image": "https://example.com/avatar.png"
    }`
	TestPostRequest(t, userData, "/api/signup", 500, "Missing field")
}

//signin

func LoginUser(t *testing.T) {
	userData := `{
        "username": "test_user_2",
		"password": "password",
    }`
	TestPostRequest(t, userData, "/api/signin", 200, "Login successfull")

}

func IncorrectPasswordError(t *testing.T) {
	userData := `{
        "username": "test_user_2",
		"password": "incorrect",
    }`
	TestPostRequest(t, userData, "/api/signin", 403, "Incorrect Password")

}
func NoUserError(t *testing.T) {
	userData := `{
        "username": "test_user_not_found",
		"password": "password",
    }`
	TestPostRequest(t, userData, "/api/signin", 404, "User not found")
}

func MissingFieldErrorSignin(t *testing.T) {
	userData := `{
        "username": "test_user_2",
    }`
	TestPostRequest(t, userData, "/api/signin", 500, "Field not found")
}
