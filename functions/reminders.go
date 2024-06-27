package functions

import (
	"fmt"
	"strings"

	"github.com/piyush7833/Chat-Api/config"
	alias "github.com/piyush7833/Chat-Api/config/Alias"
	"github.com/piyush7833/Chat-Api/services"
	"github.com/piyush7833/Chat-Api/types"
)

func CreateReminder(reminder types.CreateReminderType) (interface{}, types.ErrorType) {
	// fmt.Println("Reminder:", reminder)

	//allowed only is user have relation with the receiver

	res, err := services.InsertRow("reminders", reminder)
	if err.StatusCode != 0 {
		if strings.Contains(err.Message, "violates foreign key constraint") {
			err.Message = "Reciever does not exists"
			err.StatusCode = 404
		}
		return nil, err
	}
	return res, types.ErrorType{}
}

func GetAllReminder(userId string, page int, reminder_type string, selectedColumns []string, orderByColumns []string, isDesc bool) (interface{}, types.ErrorType) {
	rAlias := alias.GetReminderAlias()
	var whereClause string
	if reminder_type == "all" {
		whereClause = fmt.Sprintf(`%s."senderId"= '%s' OR %s."receiverId" = '%s'`, rAlias["reminder"], userId, rAlias["reminder"], userId)
	} else if reminder_type == "sent" {
		whereClause = fmt.Sprintf(` %s."senderId" = '%s'`, rAlias["reminder"], userId)
	} else if reminder_type == "received" {
		whereClause = fmt.Sprintf(` %s."receiverId" ='%s'`, rAlias["reminder"], userId)
	}

	joinClause := fmt.Sprintf(`JOIN "users" %s ON %s."senderId" = %s.id JOIN "users" %s ON %s."receiverId" = %s.id`, rAlias["sender"], rAlias["reminder"], rAlias["sender"], rAlias["receiver"], rAlias["reminder"], rAlias["receiver"])
	joinColumns := config.GetJoinReminderColumns(rAlias["reminder"], rAlias["sender"], rAlias["receiver"])

	res, err := services.GetRows("reminders", page, selectedColumns, config.GetValidReminderColumns(), &whereClause, &joinClause, joinColumns, orderByColumns, isDesc)

	if err.StatusCode != 0 {
		if strings.Contains(err.Message, "No reminders found") {
			err.StatusCode = 200
		}
		return nil, err
	}
	return res, types.ErrorType{}
}

func GetParticularReminder(userId string, id string, page int, selectedColumns []string, orderByColumns []string, isDesc bool) (interface{}, types.ErrorType) {
	rAlias := alias.GetReminderAlias()

	whereClause := fmt.Sprintf(`%s.id = '%s' and (%s."senderId"='%s' or %s."receiverId"='%s')`, rAlias["reminder"], id, rAlias["reminder"], userId, rAlias["reminder"], userId)

	joinClause := fmt.Sprintf(`JOIN "users" %s ON %s."senderId" = %s.id JOIN "users" %s ON %s."receiverId" = %s.id`, rAlias["sender"], rAlias["reminder"], rAlias["sender"], rAlias["receiver"], rAlias["reminder"], rAlias["receiver"])

	joinColumns := config.GetJoinReminderColumns(rAlias["reminder"], rAlias["sender"], rAlias["receiver"])

	res, err := services.GetRows("reminders", page, selectedColumns, config.GetValidReminderColumns(), &whereClause, &joinClause, joinColumns, orderByColumns, isDesc)

	if err.StatusCode != 0 {
		if strings.Contains(err.Message, "No reminders found") {
			return nil, types.ErrorType{
				Message:    "Reminder does not exists or it is not associated with you",
				StatusCode: 403,
			}
		}
		return nil, err
	}

	return res[0], types.ErrorType{}
}

func UpdateReminder(reminder types.UpdateReminderType, id string, userId string) (interface{}, types.ErrorType) {
	whereClause := fmt.Sprintf(`id = '%s' and ("senderId"='%s' or "receiverId"='%s')`, id, userId, userId)

	res, err := services.UpdateRows("reminders", reminder, &whereClause, config.UpdateValidReminderColumns())
	if err.StatusCode != 0 {
		if strings.Contains(err.Message, "No rows updated") {
			err.Message = "You can't update the reminder which is not associated to you"
			err.StatusCode = 403
		}
		return nil, err
	}
	return res, types.ErrorType{}
}

func DeleteReminder(id string, userId string) (interface{}, types.ErrorType) {
	whereClause := fmt.Sprintf(`id = '%s' and ("senderId"='%s' or "receiverId"='%s')`, id, userId, userId)

	res, err := services.DeleteRow("reminders", whereClause)
	if err.StatusCode != 0 {
		return nil, err
	}
	if res == 0 {
		return nil, types.ErrorType{
			Message:    "You can't delete the reminder which is not associated to you",
			StatusCode: 403,
		}
	}
	return res, types.ErrorType{}
}
