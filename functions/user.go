package functions

import (
	"github.com/piyush7833/Chat-Api/services"
	"github.com/piyush7833/Chat-Api/types"
)

func UpdateUser(user types.UpdateUserType, id string) (interface{}, types.ErrorType) {
	var where *string
	var whereClause = "id = " + "'" + id + "'"
	where = &whereClause
	validColumns := map[string]bool{
		"name":  true,
		"email": true,
		"image": true,
		"phone": true,
	}
	rowsAffected, err := services.UpdateRows("User", user, where, validColumns)
	if err.StatusCode != 0 {
		return nil, types.ErrorType{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
	} else if rowsAffected == 0 {
		return nil, types.ErrorType{
			Message:    "No user found",
			StatusCode: 404,
		}
	}
	return nil, error
}

func GetAllUser(page int, columns []string) (interface{}, types.ErrorType) {
	validColumns := map[string]bool{
		"id":        true,
		"name":      true,
		"username":  true,
		"email":     true,
		"password":  true,
		"image":     true,
		"phone":     true,
		"createdAt": true,
	}
	results, err := services.GetRows("User", page, columns, validColumns, nil)
	if err.StatusCode != 0 {
		return nil, types.ErrorType{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
	}
	return results, types.ErrorType{}
}

func GetParticularUserById(id string, columns []string) (interface{}, types.ErrorType) {

	validColumns := map[string]bool{
		"id":        true,
		"name":      true,
		"username":  true,
		"email":     true,
		"password":  true,
		"image":     true,
		"phone":     true,
		"createdAt": true,
	}
	var where *string
	var whereClause = "id = " + "'" + id + "'"
	where = &whereClause
	results, err := services.GetRows("User", 0, columns, validColumns, where)
	if err.StatusCode != 0 {
		return nil, types.ErrorType{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
	}
	return results, error
	// return res, error
}

func GetUsersByUsername(userName string, columns []string) (interface{}, types.ErrorType) {
	validColumns := map[string]bool{
		"id":        true,
		"name":      true,
		"username":  true,
		"email":     true,
		"password":  true,
		"image":     true,
		"phone":     true,
		"createdAt": true,
	}
	var where *string
	var whereClause = "username = " + "'" + userName + "'"
	where = &whereClause
	results, err := services.GetRows("User", 0, columns, validColumns, where)
	if err.StatusCode != 0 {
		return nil, types.ErrorType{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
	}
	return results, error
}
func DeleteUser(id string) (interface{}, types.ErrorType) {
	var where *string
	var whereClause = "id = " + "'" + id + "'"
	where = &whereClause
	rowsAffected, err := services.DeleteRow("User", *where)
	if err.StatusCode != 0 {
		return nil, types.ErrorType{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
	} else if rowsAffected == 0 {
		return nil, types.ErrorType{
			Message:    "No user found",
			StatusCode: 404,
		}
	}
	return nil, error
}
