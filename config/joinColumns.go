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
