package functions

import (
	"fmt"

	"github.com/piyush7833/Chat-Api/config"
	alias "github.com/piyush7833/Chat-Api/config/Alias"
	"github.com/piyush7833/Chat-Api/services"
	"github.com/piyush7833/Chat-Api/types"
)

func CreateTag(tag types.CreateTagType) (interface{}, types.ErrorType) {
	res, err := services.InsertRow("tags", tag)
	if err.StatusCode != 0 {
		return nil, err
	}
	return res, types.ErrorType{}
}

func GetAllTags(page int, columns []string, userId string, relation_type string, orderBy []string, isDesc bool) ([]map[string]interface{}, types.ErrorType) {
	var whereClause string
	urAlias := alias.GetTagsAlias()
	joinColumns := config.GetJoinTagColumns(urAlias["tags"], urAlias["user"], urAlias["message"], urAlias["reminder"])
	joinClause := fmt.Sprintf(`JOIN "users" %s ON %s."userId" = %s.id JOIN "messages" %s ON %s."relatedUserId" = %s.id JOIN "messages" %s ON %s."relatedUserId" = %s.id`, urAlias["user"], urAlias["tags"], urAlias["user"], urAlias["message"], urAlias["tags"], urAlias["message"], urAlias["reminder"], urAlias["tags"], urAlias["reminder"])
	return services.GetRows("tags", page, columns, config.GetValidTagsColumns(), &whereClause, &joinClause, joinColumns, orderBy, isDesc)
}
