package test

import (
	"net/http"
	"testing"
)

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

func TestLoginUser(t *testing.T) {
	userData := `{
        "username": "test_user",
		"password": "password"
    }`
	TestPostRequest(t, userData, "/api/signin", 200, "Login successfull")

}

func TestIncorrectPasswordError(t *testing.T) {
	userData := `{
        "username": "test_user",
		"password": "incorrect"
    }`
	TestPostRequest(t, userData, "/api/signin", 403, "Incorrect Password")

}
func TestNoUserErrorSignin(t *testing.T) {
	userData := `{
        "username": "test_user_not_found",
		"password": "password"
    }`
	TestPostRequest(t, userData, "/api/signin", 404, "User not found")
}

func TestMissingFieldErrorSignin(t *testing.T) {
	userData := `{
        "username": "test_user"
    }`
	TestPostRequest(t, userData, "/api/signin", 500, "Field not found")
}
