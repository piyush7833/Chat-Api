package alias

var mainAlias = "t"

func GetUserRelationAlias() map[string]string {
	return map[string]string{
		"userRelation": mainAlias,
		"user":         "u1",
		"relatedUser":  "u2",
	}
}
