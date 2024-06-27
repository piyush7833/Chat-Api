package alias

var mainAlias = "t"

func GetUserRelationAlias() map[string]string {
	return map[string]string{
		"userRelation": mainAlias,
		"user":         "u1",
		"relatedUser":  "u2",
	}
}
func GetReminderAlias() map[string]string {
	return map[string]string{
		"reminder": mainAlias,
		"sender":   "s",
		"receiver": "r",
	}
}
func GetTagsAlias() map[string]string {
	return map[string]string{
		"tags":     mainAlias,
		"user":     "u",
		"message":  "m",
		"reminder": "r",
	}
}
