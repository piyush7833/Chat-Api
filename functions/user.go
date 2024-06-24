package functions

import (
	"github.com/piyush7833/Chat-Api/config"
	"github.com/piyush7833/Chat-Api/services"
	"github.com/piyush7833/Chat-Api/types"
)

func UpdateUser(user types.UpdateUserType, id string) (interface{}, types.ErrorType) {
	var where *string
	var whereClause = "id = " + "'" + id + "'"
	where = &whereClause
	rowsAffected, err := services.UpdateRows("users", user, where, config.UpdateValidUserColumns())
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
	return nil, types.ErrorType{}
}

func GetAllUser(page int, columns []string, orderBy []string, isDesc bool) (interface{}, types.ErrorType) {
	results, err := services.GetRows("users", page, columns, config.GetValidUserColumns(), nil, nil, nil, nil, true)
	if err.StatusCode != 0 {
		return nil, types.ErrorType{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
	}
	return results, types.ErrorType{}
}

func GetParticularUserById(id string, columns []string) (interface{}, types.ErrorType) {
	var whereClause = "id = " + "'" + id + "'"
	results, err := services.GetRows("users", 0, columns, config.GetValidUserColumns(), &whereClause, nil, nil, nil, true)
	if err.StatusCode != 0 {
		return nil, types.ErrorType{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
	}
	return results[0], types.ErrorType{}
	// return res, error
}

func GetUsersByUsername(userName string, columns []string) (interface{}, types.ErrorType) {
	var whereClause = "username = " + "'" + userName + "'"
	results, err := services.GetRows("users", 0, columns, config.GetValidUserColumns(), &whereClause, nil, nil, nil, true)
	if err.StatusCode != 0 {
		return nil, types.ErrorType{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
	}
	return results[0], types.ErrorType{}
}

func DeleteUser(id string) (interface{}, types.ErrorType) {
	var where *string
	var whereClause = "id = " + "'" + id + "'"
	where = &whereClause

	rowsAffected, err := services.DeleteRow("users", *where)
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
	return nil, types.ErrorType{}
}
