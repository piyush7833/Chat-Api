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
	TestPostRequest(t, userData, "/api/signup", 201, "User created successfully", true)
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
	TestPostRequest(t, userData, "/api/signup", http.StatusForbidden, "User with given user name or email or phone already exist", true)
}
func TestMissingFieldErrorSignup(t *testing.T) {
	userData := `{
        "username": "test_user_2",
        "name": "Test User",
		"password": "password",
        "phone": "+1234567892",
        "image": "https://example.com/avatar.png"
    }`
	TestPostRequest(t, userData, "/api/signup", 500, "Missing field", true)
}

//signin

func TestLoginUser(t *testing.T) {
	userData := `{
        "username": "test_user",
        "password": "password"
    }`
	TestPostRequest(t, userData, "/api/signin", 200, "Login successfull", true)
}

func TestIncorrectPasswordError(t *testing.T) {
	userData := `{
        "username": "test_user",
		"password": "incorrect"
    }`
	TestPostRequest(t, userData, "/api/signin", 403, "Incorrect Password", true)

}
func TestNoUserErrorSignin(t *testing.T) {
	userData := `{
        "username": "test_user_not_found",
		"password": "password"
    }`
	TestPostRequest(t, userData, "/api/signin", 404, "User not found", true)
}

func TestMissingFieldErrorSignin(t *testing.T) {
	userData := `{
        "username": "test_user"
    }`
	TestPostRequest(t, userData, "/api/signin", 500, "Field not found", true)
}

func TestCreateOtherUsers(t *testing.T) {
	userData := `{
        "username": "related_user",
        "name": "related_user",
        "email": "related_user@example.com",
        "password": "password",
        "phone": "+1234557890",
        "image": "https://example.com/avatar.png"
    }`
	TestPostRequest(t, userData, "/api/signup", 201, "User created successfully", true)
	userData2 := `{
        "username": "not_allowed_user",
        "name": "not_allowed_user",
        "email": "not_allowed_user@example.com",
        "password": "password",
        "phone": "+1234457890",
        "image": "https://example.com/avatar.png"
    }`
	TestPostRequest(t, userData2, "/api/signup", 201, "User created successfully", true)
}

// logged in related user
func TestLoginRelatedUserUsingEmail(t *testing.T) {
	userData := `{
        "email": "related_user@example.com",
		"password": "password"
    }`
	TestPostRequest(t, userData, "/api/signin", 200, "Login successfull", true)
}
func TestLoginNotAllowedUserUsingPhone(t *testing.T) {
	userData := `{
        "phone": "+1234457890",
		"password": "password"
    }`
	TestPostRequest(t, userData, "/api/signin", 200, "Login successfull", true)
}
