package config

func GetUserDefaultColumns() []string {
	return []string{"id", "name", "username", "email", "phone", "image", "createdAt"}
}

func GetDefaultUserRelationColumns() []string {
	return []string{"id", "userId", "relatedUserId", "user_name", "related_user_name", "status", "createdAt", "updatedAt"}
}

func GetDefaultTagsColumns() []string {
	return []string{"id", "userId", "user_name", "messageId", "message_content", "message_type", "reminderId", "reminder_title", "createdAt", "updatedAt"}
}

func GetDefaultReminderColumns() []string {
	return []string{"id", "message", "time", "senderId", "sender_name", "receiverId", "reciever_name", "tune", "createdAt", "updatedAt"}
}
