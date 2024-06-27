package test

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	err := Init()
	if err != nil {
		log.Fatal(err)
	}
	code := m.Run()
	Close()
	os.Exit(code)
}

func TestOrder(t *testing.T) {

	//tests for main user logged in
	t.Run("TestCreateUser", TestCreateUser)
	t.Run("TestDuplicateUserError", TestDuplicateUserError)
	t.Run("TestMissingFieldErrorSignup", TestMissingFieldErrorSignup)
	t.Run("TestLoginUser", TestLoginUser)

	t.Run("TestIncorrectPasswordError", TestIncorrectPasswordError)
	t.Run("TestNoUserErrorSignin", TestNoUserErrorSignin)
	t.Run("TestMissingFieldErrorSignin", TestMissingFieldErrorSignin)

	t.Run("TestGetAllUser", TestGetAllUser)
	t.Run("TestGetAllUserErrorUnAuthorized", TestGetAllUserErrorUnAuthorized)
	t.Run("TestGetAllUserErrorInvalidColumns", TestGetAllUserErrorInvalidColumns)

	t.Run("TestGetUserById", TestGetUserById)
	t.Run("TestGetUserByIdErrorInvalidColumns", TestGetUserByIdErrorInvalidColumns)
	t.Run("TestGetUserByIdErrorUnauthorized", TestGetUserByIdErrorUnauthorized)
	t.Run("TestGetUserByIdErrorUserNotFound", TestGetUserByIdErrorUserNotFound)

	t.Run("TestGetUserByUsername", TestGetUserByUsername)
	t.Run("TestGetUserByUsernameErrorInvalidColumns", TestGetUserByUsernameErrorInvalidColumns)
	t.Run("TestGetUserByUsernameErrorUnauthorized", TestGetUserByUsernameErrorUnauthorized)
	t.Run("TestGetUserByUsernameErrorUserNotFound", TestGetUserByUsernameErrorUserNotFound)

	t.Run("TestGetAuthenticatedUser", TestGetAuthenticatedUser)
	t.Run("TestGetAuthenticatedUserErrorUnAuthorized", TestGetAuthenticatedUserErrorUnauthorized)
	t.Run("TestGetAuthenticatedUserErrorInvalidColumns", TestGetAuthenticatedUserErrorInvalidColumns)

	t.Run("TestUpdateUserErrorInvalidColumns", TestUpdateUserErrorInvalidColumns)
	t.Run("TestUpdateUserErrorUnauthorized", TestUpdateUserErrorUnauthorized)

	// // create other users after testing user routes
	t.Run("TestCreateOtherUsers", TestCreateOtherUsers)

	t.Run("TestGetAllUserRelationNoneFound", TestGetAllUserRelationNoneFound)

	t.Run("TestCreateUserRelation", TestCreateUserRelation)
	t.Run("TestCreateUserRelationErrorDuplicateRelation", TestCreateUserRelationErrorDuplicateRelation)
	t.Run("TestCreateUserRelationErrorMissingField", TestCreateUserRelationErrorMissingField)
	// t.Run("TestCreateUserRelationErrorRelationAlreadyExist",TestCreateUserRelationErrorRelationAlreadyExist) //need to handle
	t.Run("TestCreateUserRelationErrorUnauthorized", TestCreateUserRelationErrorUnauthorized)
	t.Run("TestCreateUserRelationErrorInvalidRelatedUserId", TestCreateUserRelationErrorInvalidRelatedUserId)

	t.Run("TestGetAllUserRelation", TestGetAllUserRelation)
	t.Run("TestGetAllUserRelationParticularStatus", TestGetAllUserRelationParticularStatus)
	t.Run("TestGetAllUserRelationParticularType", TestGetAllUserRelationParticularType)
	t.Run("TestGetAllUserRelationParticularTypeParticularStatus", TestGetAllUserRelationParticularTypeParticularStatus)
	t.Run("TestGetAllUserRelationErrorUnauthorized", TestGetAllUserRelationErrorUnauthorized)
	t.Run("TestGetAllUserRelationErrorInvalidColumns", TestGetAllUserRelationErrorInvalidColumns)
	t.Run("TestGetAllUserRelationErrorInvalidStatus", TestGetAllUserRelationErrorInvalidStatus)
	t.Run("TestGetAllUserRelationErrorInvalidType", TestGetAllUserRelationErrorInvalidType)

	t.Run("TestGetParticularUserRelation", TestGetParticularUserRelation)
	t.Run("TestGetParticularUserRelationErrorUserRelationNotFound", TestGetParticularUserRelationErrorUserRelationNotFound)
	t.Run("TestGetParticularUserRelationErrorInvalidColumn", TestGetParticularUserRelationErrorInvalidColumn)
	t.Run("TestGetParticularUserRelationErrorUnauthorized", TestGetParticularUserRelationErrorUnauthorized)

	t.Run("TestUpdateUserRelationErrorNotAllowed", TestUpdateUserRelationErrorNotAllowed)
	t.Run("TestUpdateUserRelationErrorStatusNotProvided", TestUpdateUserRelationErrorStatusNotProvided)
	t.Run("TestUpdateUserRelationErrorNotFound", TestUpdateUserRelationErrorNotFound)
	t.Run("TestUpdateUserRelationErrorUnauthorized", TestUpdateUserRelationErrorUnauthorized)

	t.Run("TestGetAllReminderNoneFound", TestGetAllReminderNoneFound)
	t.Run("TestCreateReminder", TestCreateReminder)
	t.Run("TestCreateReminderErrorMissingField", TestCreateReminderErrorMissingField)
	t.Run("TestCreateReminderErrorInvalidTime", TestCreateReminderErrorInvalidTime)
	t.Run("TestCreateReminderErrorInvalidReceiverId", TestCreateReminderErrorInvalidReceiverId)
	t.Run("TestCreateReminderErrorUnauthorized", TestCreateReminderErrorUnauthorized)

	t.Run("TestGetAllReminder", TestGetAllReminder)
	t.Run("TestGetAllReminderErrorUnauthorized", TestGetAllReminderErrorUnauthorized)
	t.Run("TestGetAllReminderErrorInvalidColumns", TestGetAllReminderErrorInvalidColumn)

	t.Run("TestGetParticularReminder", TestGetParticularReminder)
	t.Run("TestGetParticularReminderErrorUnauthorized", TestGetParticularReminderErrorUnauthorized)
	t.Run("TestGetParticularReminderErrorInvalidColumns", TestGetParticularReminderErrorInvalidColumns)
	t.Run("TestGetParticularReminderErrorReminderNotFound", TestGetParticularReminderErrorNotFound)

	t.Run("TestUpdateReminder", TestUpdateReminder)
	t.Run("TestUpdateReminderErrorUnauthorized", TestUpdateReminderErrorUnauthorized)
	t.Run("TestUpdateReminderErrorInvalidColumns", TestUpdateReminderErrorInvalidColumn)
	t.Run("TestUpdateReminderErrorReminderNotFound", TestUpdateReminderErrorNotFound)

	// //tests for related user logged in
	t.Run("Test login related user", TestLoginRelatedUserUsingEmail)

	t.Run("TestUpdateUserRelation", TestUpdateUserRelation)
	t.Run("TestUpdateUserRelationErrorNotAllowedToRevertToPendingStatus", TestUpdateUserRelationErrorNotAllowedToRevertToPendingStatus)

	t.Run("TestGetParticularUserRelation", TestGetParticularUserRelation)
	t.Run("TestUpdateReminder", TestUpdateReminder)

	//tests for third user who is not allowed for things
	t.Run("Test login not allowed user", TestLoginNotAllowedUserUsingPhone)

	t.Run("TestUpdateUserRelationErrorNotAllowed", TestUpdateUserRelationErrorNotAllowed)
	t.Run("TestGetParticularUserRelationErrorNotAllowed", TestGetParticularUserRelationErrorNotAllowed)
	t.Run("TestDeleteUserRelationErrorNotAllowed", TestDeleteUserRelationErrorNotAllowed)

	t.Run("TestGetParticularReminderErrorNotAllowed", TestGetParticularReminderErrorNotAllowed)
	t.Run("TestUpdateReminderErrorNotAllowed", TestUpdateReminderErrorNotAllowed)
	t.Run("TestDeleteReminderErrorNotAllowd", TestDeleteReminderErrorNotAllowed)

	// //delete/update tests for main user logged in user
	t.Run("Test login main user", TestLoginUser)
	t.Run("TestUpdateUser", TestUpdateUser)

	t.Run("TestDeleteUserRelationErrorNotFound", TestDeleteUserRelationErrorNotFound)
	t.Run("TestDeleteUserRelationErrorUnauthorised", TestDeleteUserRelationErrorUnauthorised)
	t.Run("TestDeleteUserRelation", TestDeleteUserRelation)

	t.Run("TestDeleteReminderErrorUnauthorized", TestDeleteReminderErrorUnauthorized)
	t.Run("TestDeleteReminderErrorReminderNotFound", TestDeleteReminderErrorNotFound)
	t.Run("TestDeleteReminder", TestDeleteReminder)

	t.Run("TestDeleteUser", TestDeleteUser)

	fmt.Println("All tests passed successfully")
}
