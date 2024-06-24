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
