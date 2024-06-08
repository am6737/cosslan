package filter

import "go.mongodb.org/mongo-driver/bson"

// NodeFilter 是用于查询节点的过滤条件
type NodeFilter struct {
	NodeID    string
	Name      string
	IP        string
	UserID    string
	Location  string
	NetworkID string
}

// ToBSON 将 NodeFilter 转换为 MongoDB 查询文档
func (filter *NodeFilter) ToBSON() bson.M {
	query := bson.M{}

	if filter.NodeID != "" {
		query["_id"] = filter.NodeID
	}

	if filter.Name != "" {
		query["name"] = filter.Name
	}

	if filter.IP != "" {
		query["ip"] = filter.IP
	}

	if filter.UserID != "" {
		query["user_id"] = filter.UserID
	}

	if filter.Location != "" {
		query["location"] = filter.Location
	}

	if filter.NetworkID != "" {
		query["network_id"] = filter.NetworkID
	}

	return query
}
