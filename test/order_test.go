package test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	Init()
	code := m.Run()
	Close()
	os.Exit(code)
}

func TestOrder(t *testing.T) {

	t.Run("TestCreateUser", TestCreateUser)
	t.Run("TestDuplicateUserError", TestDuplicateUserError)
	t.Run("TestMissingFieldErrorSignup", TestMissingFieldErrorSignup)
	t.Run("TestLoginUser", TestLoginUser)
	t.Run("TestIncorrectPasswordError", TestIncorrectPasswordError)
	t.Run("TestNoUserErrorSignin", TestNoUserErrorSignin)
	t.Run("TestMissingFieldErrorSignin", TestMissingFieldErrorSignin)
}
