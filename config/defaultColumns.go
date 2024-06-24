package config

func GetUserDefaultColumns() []string {
	return []string{"id", "name", "username", "email", "phone", "image", "createdAt"}
}

func GetDefaultUserRelationColumns() []string {
	return []string{"id", "userId", "relatedUserId", "user_name", "related_user_name", "status", "createdAt", "updatedAt"}
}
