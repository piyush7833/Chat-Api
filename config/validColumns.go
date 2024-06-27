package config

func GetValidUserColumns() map[string]bool {
	return map[string]bool{
		"id":            true,
		"name":          true,
		"username":      true,
		"email":         true,
		"phone":         true,
		"image":         true,
		"lastSeen":      true,
		"isOnline":      true,
		"phoneVerified": true,
		"emailVerified": true,
		"createdAt":     true,
	}
}

func UpdateValidUserColumns() map[string]bool {
	return map[string]bool{
		"name":  true,
		"email": true,
		"image": true,
		"phone": true,
	}
}
func GetValidUserRelationColumns() map[string]bool {
	return map[string]bool{
		"id":                true,
		"userId":            true,
		"user_name":         true,
		"relatedUserId":     true,
		"related_user_name": true,
		"status":            true,
		"createdAt":         true,
		"updatedAt":         true,
	}
}
func GetValidTagsColumns() map[string]bool {
	return map[string]bool{
		"id":              true,
		"userId":          true,
		"user_name":       true,
		"messageId":       true,
		"message_content": true,
		"message_type":    true,
		"reminderId":      true,
		"reminder_title":  true,
		"title":           true,
		"createdAt":       true,
		"updatedAt":       true,
	}
}

func GetValidReminderColumns() map[string]bool {
	return map[string]bool{
		"id":            true,
		"message":       true,
		"time":          true,
		"senderId":      true,
		"sender_name":   true,
		"receiverId":    true,
		"reciever_name": true,
		"tune":          true,
		"createdAt":     true,
		"updatedAt":     true,
	}
}
func UpdateValidReminderColumns() map[string]bool {
	return map[string]bool{
		"message": true,
		"time":    true,
		"tune":    true,
	}
}
