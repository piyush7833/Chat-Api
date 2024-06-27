package test

import (
	"fmt"
	"testing"

	"github.com/piyush7833/Chat-Api/config"
	"github.com/piyush7833/Chat-Api/services"
)

// test for create reminder
func TestCreateReminder(t *testing.T) {
	var whereClause = "username = " + "'related_user'"
	res, err := services.GetRows("users", 0, []string{"id"}, config.GetValidUserColumns(), &whereClause, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	reminderData := `{
		"message": "Test Reminder",
		"time": "2025-05-01T12:00:00Z",
		"tune" : "https://example.com/tune.mp3",
		"receiverId":  "` + fmt.Sprintf("%v", res[0]["id"]) + `"
	}`
	TestPostRequest(t, reminderData, "/api/protected/reminder/create", 201, "Reminder created successfully", true)
}

func TestCreateReminderErrorMissingField(t *testing.T) {
	reminderData := `{
		"message": "Test Reminder",
		"time": "2025-05-01T12:00:00Z",
	}`
	TestPostRequest(t, reminderData, "/api/protected/reminder/create", 500, "Missing field", true)
}

func TestCreateReminderErrorInvalidTime(t *testing.T) {
	var whereClause = "username = " + "'related_user'"
	res, err := services.GetRows("users", 0, []string{"id"}, config.GetValidUserColumns(), &whereClause, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	reminderData := `{
		"message": "Test Reminder",
		"time": "01 03",
		"tune" : "https://example.com/tune.mp3",
		"receiverId":   "` + fmt.Sprintf("%v", res[0]["id"]) + `"
	}`
	TestPostRequest(t, reminderData, "/api/protected/reminder/create", 500, "Incorrect time format", true)
}

func TestCreateReminderErrorInvalidReceiverId(t *testing.T) {
	reminderData := `{
		"message": "Test Reminder",
		"time": "2025-05-01T12:00:00Z",
		"tune" : "https://example.com/tune.mp3",
		"receiverId":  "randomId"
	}`
	TestPostRequest(t, reminderData, "/api/protected/reminder/create", 404, "Reciever does not exists", true)
}

func TestCreateReminderErrorUnauthorized(t *testing.T) {
	reminderData := `{
		"message": "Test Reminder",
		"time": "2025-05-01T12:00:00Z",
		"tune" : "https://example.com/tune.mp3",
		"receiverId": "randomId"
	}`
	TestPostRequest(t, reminderData, "/api/protected/reminder/create", 401, "login required", false)
}

// test for get particular reminder
func TestGetParticularReminder(t *testing.T) {
	res, err := services.GetRows("reminders", 0, []string{"id"}, config.GetValidReminderColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	TestGetRequest(t, "/api/protected/reminder/get?id="+fmt.Sprintf("%v", res[0]["id"]), 200, "Reminder fetched successfully", true)
}

func TestGetParticularReminderErrorInvalidColumns(t *testing.T) {
	res, err := services.GetRows("reminders", 0, []string{"id"}, config.GetValidReminderColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	TestGetRequest(t, "/api/protected/reminder/get?id="+fmt.Sprintf("%v", res[0]["id"])+"&columns=invalid", 400, "no valid columns selected", true)
}

func TestGetParticularReminderErrorNotFound(t *testing.T) {
	TestGetRequest(t, "/api/protected/reminder/get?id=randomId", 403, "Reminder does not exists or it is not associated with you", true)
}

func TestGetParticularReminderErrorUnauthorized(t *testing.T) {
	res, err := services.GetRows("reminders", 0, []string{"id"}, config.GetValidReminderColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	TestGetRequest(t, "/api/protected/reminder/get?id="+fmt.Sprintf("%v", res[0]["id"]), 401, "login required", false)
}

func TestGetParticularReminderErrorNotAllowed(t *testing.T) { //logged in by third user
	res, err := services.GetRows("reminders", 0, []string{"id"}, config.GetValidReminderColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	TestGetRequest(t, "/api/protected/reminder/get?id="+fmt.Sprintf("%v", res[0]["id"]), 403, "Reminder does not exists or it is not associated with you", true)
}

// test for get all reminder
func TestGetAllReminder(t *testing.T) {
	TestGetRequest(t, "/api/protected/reminder/get-all", 200, "Reminder fetched successfully", true)
}
func TestGetAllReminderSent(t *testing.T) {
	TestGetRequest(t, "/api/protected/reminder/get-all?type=sent", 200, "Reminder fetched successfully", true)
}
func TestGetAllReminderRecieved(t *testing.T) {
	TestGetRequest(t, "/api/protected/reminder/get-all?type=received", 200, "Reminder fetched successfully", true)
}

func TestGetAllReminderErrorUnauthorized(t *testing.T) {
	TestGetRequest(t, "/api/protected/reminder/get-all", 401, "login required", false)
}
func TestGetAllReminderNoneFound(t *testing.T) {
	TestGetRequest(t, "/api/protected/reminder/get-all", 200, "No reminder found", true)
}
func TestGetAllReminderErrorInvalidColumn(t *testing.T) {
	TestGetRequest(t, "/api/protected/reminder/get-all?columns=invalid", 400, "no valid columns selected", true)
}

// test for update reminder
func TestUpdateReminder(t *testing.T) {
	res, err := services.GetRows("reminders", 0, []string{"id"}, config.GetValidReminderColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	reminderData := `{
		"message": "Test Reminder Updated",
		"time": "2025-05-01T12:00:00Z"
	}`
	TestPatchRequest(t, reminderData, "/api/protected/reminder/update?id="+fmt.Sprintf("%v", res[0]["id"]), 200, "Reminder updated successfully", true)
}

func TestUpdateReminderErrorInvalidColumn(t *testing.T) {
	res, err := services.GetRows("reminders", 0, []string{"id"}, config.GetValidReminderColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	reminderData := `{
		"receiverId": "randomId"
	}`
	TestPatchRequest(t, reminderData, "/api/protected/reminder/update?id="+fmt.Sprintf("%v", res[0]["id"]), 400, "no valid columns selected", true)
}

func TestUpdateReminderErrorUnauthorized(t *testing.T) {
	res, err := services.GetRows("reminders", 0, []string{"id"}, config.GetValidReminderColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	reminderData := `{
		"message": "Test Reminder Updated",
		"time": "2025-05-01T12:00:00Z"
	}`
	TestPatchRequest(t, reminderData, "/api/protected/reminder/update?id="+fmt.Sprintf("%v", res[0]["id"]), 401, "login required", false)
}

func TestUpdateReminderErrorNotAllowed(t *testing.T) { //logged in by third user
	res, err := services.GetRows("reminders", 0, []string{"id"}, config.GetValidReminderColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	reminderData := `{
		"message": "Test Reminder Updated",
		"time": "2025-05-01T12:00:00Z"
	}`
	TestPatchRequest(t, reminderData, "/api/protected/reminder/update?id="+fmt.Sprintf("%v", res[0]["id"]), 403, "You can't update the reminder which is not associated to you", true)
}

func TestUpdateReminderErrorNotFound(t *testing.T) {
	reminderData := `{
		"message": "Test Reminder Updated",
		"time": "2025-05-01T12:00:00Z"
	}`
	TestPatchRequest(t, reminderData, "/api/protected/reminder/update?id=randomId", 403, "You can't update the reminder which is not associated to you", true)
}

// test for delete reminder
func TestDeleteReminder(t *testing.T) {
	res, err := services.GetRows("reminders", 0, []string{"id"}, config.GetValidReminderColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	TestDeleteRequest(t, "/api/protected/reminder/delete?id="+fmt.Sprintf("%v", res[0]["id"]), 200, "Reminder deleted successfully", true)
}

func TestDeleteReminderErrorUnauthorized(t *testing.T) {
	res, err := services.GetRows("reminders", 0, []string{"id"}, config.GetValidReminderColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	TestDeleteRequest(t, "/api/protected/reminder/delete?id="+fmt.Sprintf("%v", res[0]["id"]), 401, "login required", false)
}

func TestDeleteReminderErrorNotAllowed(t *testing.T) { //logged in by third user
	res, err := services.GetRows("reminders", 0, []string{"id"}, config.GetValidReminderColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		t.Errorf(err.Message)
		return
	}
	TestDeleteRequest(t, "/api/protected/reminder/delete?id="+fmt.Sprintf("%v", res[0]["id"]), 403, "You can't delete the reminder which is not associated to you", true)
}

func TestDeleteReminderErrorNotFound(t *testing.T) {
	TestDeleteRequest(t, "/api/protected/reminder/delete?id=randomId", 403, "You can't delete the reminder which is not associated to you", true)
}
