package line

import "go.mongodb.org/mongo-driver/bson"

// LineFilter 是用于查询线路的过滤条件
type LineFilter struct {
	ID         string
	Name       string
	Type       string
	UserID     string
	NextNodeID string
}

// ToBSON 将 LineFilter 转换为 MongoDB 查询文档
func (filter *LineFilter) ToBSON() bson.M {
	query := bson.M{}

	if filter.ID != "" {
		query["_id"] = filter.ID
	}

	if filter.Name != "" {
		query["name"] = filter.Name
	}

	if filter.Type != "" {
		query["type"] = filter.Type
	}

	if filter.UserID != "" {
		query["user_id"] = filter.UserID
	}

	if filter.NextNodeID != "" {
		query["next_node_id"] = filter.NextNodeID
	}

	return query
}
