package config

func GetJoinUserRelationColumns(ogAlias string, userAlias string, relatedUserAlias string) map[string]string {
	return map[string]string{
		"id":                ogAlias + ".id",
		"userId":            ogAlias + ".userId",
		"user_name":         userAlias + ".name",
		"relatedUserId":     ogAlias + ".relatedUserId",
		"related_user_name": relatedUserAlias + ".name",
		"status":            ogAlias + ".status",
		"createdAt":         ogAlias + ".createdAt",
		"updatedAt":         ogAlias + ".updatedAt",
	}
}
func GetJoinTagColumns(ogAlias string, userAlias string, messageAlias string, reminderAlias string) map[string]string {
	return map[string]string{
		"id":              ogAlias + ".id",
		"userId":          ogAlias + ".userId",
		"user_name":       userAlias + ".name",
		"messageId":       ogAlias + ".messageId",
		"message_content": messageAlias + ".content",
		"message_type":    messageAlias + ".type",
		"reminderId":      ogAlias + ".reminderId",
		"reminder_title":  reminderAlias + ".title",
		"createdAt":       ogAlias + ".createdAt",
		"updatedAt":       ogAlias + ".updatedAt",
	}
}
func GetJoinReminderColumns(ogAlias string, senderAlias string, recieverrAlias string) map[string]string {
	return map[string]string{
		"id":            ogAlias + ".id",
		"message":       ogAlias + ".message",
		"time":          ogAlias + ".time",
		"senderId":      ogAlias + ".senderId",
		"sender_name":   senderAlias + ".name",
		"receiverId":    ogAlias + ".receiverId",
		"reciever_name": recieverrAlias + ".name",
		"tune":          ogAlias + ".tune",
		"createdAt":     ogAlias + ".createdAt",
		"updatedAt":     ogAlias + ".updatedAt",
	}
}
