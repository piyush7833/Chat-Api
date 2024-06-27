package functions

import (
	"fmt"
	"strings"

	"github.com/piyush7833/Chat-Api/config"
	alias "github.com/piyush7833/Chat-Api/config/Alias"
	"github.com/piyush7833/Chat-Api/services"
	"github.com/piyush7833/Chat-Api/types"
)

func CreateUserRelation(relation types.CreateUserRelationType) (interface{}, types.ErrorType) {
	if *relation.RelatedUserId == *relation.UserId {
		return nil, types.ErrorType{
			Message:    "User and related user can't be same",
			StatusCode: 400,
		}
	}
	res, err := services.InsertRow("userRelation", relation)
	if err.StatusCode != 0 {
		if strings.Contains(err.Message, "unique_user_related_user") {
			err.Message = "Relation already exists"
			err.StatusCode = 400
		} else if strings.Contains(err.Message, "violates foreign key constraint") {
			err.Message = "Related user does not exists"
			err.StatusCode = 404
		}
		return nil, err
	}
	return res, types.ErrorType{}
}

func UpdateUserRelation(relation types.UpdateUserRelationType, id string, userId string) (interface{}, types.ErrorType) {
	validColumns := map[string]bool{
		"status": true,
	}
	var whereClause = "id = " + "'" + id + "'"
	userRelation, err := GetParticularUserRelation(id, []string{"userId", "relatedUserId"}, userId)
	if err.StatusCode != 0 {
		if strings.Contains(err.Message, "No userRelation found") {
			err.Message = "User relation does not exists or it is not associated with you"
			err.StatusCode = 403
		}
		return nil, err
	}
	userRelationMap := userRelation.(map[string]interface{})
	if relation.Status == nil {
		return nil, types.ErrorType{
			Message:    "Status is required",
			StatusCode: 400,
		}
	}
	fmt.Println(userRelationMap["relatedUserId"] != userId && *relation.Status == "friends")
	if userRelationMap["relatedUserId"] != userId && *relation.Status == "friends" {
		return nil, types.ErrorType{
			Message:    "Only related user can update status to friends",
			StatusCode: 403,
		}
	}
	if userRelationMap["status"] != "pending" && *relation.Status == "pending" {
		return nil, types.ErrorType{
			Message:    "User relation status is already updated to friends or blocked, you can't update status to pending again",
			StatusCode: 406,
		}
	}
	rowsAffected, err := services.UpdateRows("userRelation", relation, &whereClause, validColumns)
	if err.StatusCode != 0 {
		if strings.Contains(err.Message, "No userRelation found") {
			err.Message = "User relation does not exists or it is not associated with you"
			err.StatusCode = 403
		}
		return nil, types.ErrorType{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
	} else if rowsAffected == 0 {
		return nil, types.ErrorType{
			Message:    "No relation found",
			StatusCode: 404,
		}
	}
	return rowsAffected, types.ErrorType{}
}

// need to handle multiple status and relation type i.e. array of status and relation type
func GetAllUserRelations(status string, page int, columns []string, userId string, relation_type string, orderBy []string, isDesc bool) (interface{}, types.ErrorType) {
	var whereClause string
	urAlias := alias.GetUserRelationAlias()
	joinColumns := config.GetJoinUserRelationColumns(urAlias["userRelation"], urAlias["user"], urAlias["relatedUser"])

	if userId == "" {
		return nil, types.ErrorType{
			Message:    "userId is required",
			StatusCode: 400,
		}
	}
	if status != "pending" && status != "blocked" && status != "friends" && status != "all" && status != "" {
		return nil, types.ErrorType{
			Message:    "Invalid status",
			StatusCode: 400,
		}
	}
	if relation_type != "sent" && relation_type != "received" && relation_type != "" && relation_type != "all" {
		return nil, types.ErrorType{
			Message:    "Invalid relation type",
			StatusCode: 400,
		}
	}

	userRelationAlias := urAlias["userRelation"]
	relatedUserIdCondition := fmt.Sprintf(`%s."relatedUserId" = '%s'`, userRelationAlias, userId)
	userIdCondition := fmt.Sprintf(`%s."userId" = '%s'`, userRelationAlias, userId)

	if status == "all" || status == "" {
		switch relation_type {
		case "all", "":
			whereClause = fmt.Sprintf(`(%s OR %s)`, relatedUserIdCondition, userIdCondition)
		case "received":
			whereClause = relatedUserIdCondition
		case "sent":
			whereClause = userIdCondition
		}
	} else {
		statusCondition := fmt.Sprintf(`%s.status = '%s'`, userRelationAlias, status)
		switch relation_type {
		case "all", "":
			whereClause = fmt.Sprintf(`%s AND (%s OR %s)`, statusCondition, relatedUserIdCondition, userIdCondition)
		case "received":
			whereClause = fmt.Sprintf(`%s AND %s`, statusCondition, relatedUserIdCondition)
		case "sent":
			whereClause = fmt.Sprintf(`%s AND %s`, statusCondition, userIdCondition)
		}
	}

	joinClause := fmt.Sprintf(`JOIN "users" %s ON %s."userId" = %s.id JOIN "users" %s ON %s."relatedUserId" = %s.id`, urAlias["user"], urAlias["userRelation"], urAlias["user"], urAlias["relatedUser"], urAlias["userRelation"], urAlias["relatedUser"])
	res, err := services.GetRows("userRelation", page, columns, config.GetValidUserRelationColumns(), &whereClause, &joinClause, joinColumns, orderBy, isDesc)
	if err.StatusCode != 0 {
		if strings.Contains(err.Message, "No userRelation found") {
			err.StatusCode = 200
		}
		return nil, err
	}
	return res, types.ErrorType{}
}

func GetParticularUserRelation(id string, columns []string, userId string) (interface{}, types.ErrorType) {
	urAlias := alias.GetUserRelationAlias()

	joinColumns := config.GetJoinUserRelationColumns(urAlias["userRelation"], urAlias["user"], urAlias["relatedUser"])

	// whereClause := urAlias["userRelation"] + ".id = " + "'" + id + "'"
	whereClause := fmt.Sprintf(`%s.id = '%s' and (%s."userId"='%s' or %s."relatedUserId"='%s')`, urAlias["userRelation"], id, urAlias["userRelation"], userId, urAlias["userRelation"], userId)

	joinClause := fmt.Sprintf(`JOIN "users" %s ON %s."userId" = %s.id JOIN "users" %s ON %s."relatedUserId" = %s.id`, urAlias["user"], urAlias["userRelation"], urAlias["user"], urAlias["relatedUser"], urAlias["userRelation"], urAlias["relatedUser"])

	res, err := services.GetRows("userRelation", 0, columns, config.GetValidUserRelationColumns(), &whereClause, &joinClause, joinColumns, nil, true)

	if err.StatusCode != 0 {
		if strings.Contains(err.Message, "No userRelation found") {
			return nil, types.ErrorType{
				Message:    "User relation does not exists or it is not associated with you",
				StatusCode: 403,
			}
		}
		return nil, err
	}

	return res[0], types.ErrorType{}
}

func DeleteUserRelations(id string, userId string) (interface{}, types.ErrorType) {
	var whereClause = fmt.Sprintf(`id = '%s' and ("userId"='%s' or "relatedUserId"='%s')`, id, userId, userId)
	// userRelation, err := GetParticularUserRelation(id, []string{"userId", "relatedUserId"}, userId)
	// if err.StatusCode != 0 {
	// 	return nil, err
	// }
	// userRelationMap := userRelation.(map[string]interface{})
	// if userRelationMap["userId"] != userId && userRelationMap["relatedUserId"] != userId {
	// 	return nil, types.ErrorType{
	// 		Message:    "Only user and related user are allowed to delete this request",
	// 		StatusCode: 403,
	// 	}
	// }
	rowsAffected, err := services.DeleteRow("userRelation", whereClause)
	if err.StatusCode != 0 {
		return nil, types.ErrorType{
			Message:    err.Message,
			StatusCode: err.StatusCode,
		}
	} else if rowsAffected == 0 {
		return nil, types.ErrorType{
			Message:    "You can't delete the relation which is not associated to you",
			StatusCode: 403,
		}
	}
	return rowsAffected, types.ErrorType{}
}
