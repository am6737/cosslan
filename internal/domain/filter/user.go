package filter

import (
	"github.com/cossteam/cosslan/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
)

// UserFilter 是用于查询用户的过滤条件
type UserFilter struct {
	ID     string
	Name   string
	Email  string
	Status entity.UserStatus
	Role   entity.UserRole
	// 可以添加其他可能的查询条件
}

// ToBSON 将 UserFilter 转换为 MongoDB 查询文档
func (filter *UserFilter) ToBSON() bson.M {
	query := bson.M{}

	if filter.ID != "" {
		query["_id"] = filter.ID
	}

	if filter.Name != "" {
		query["name"] = filter.Name
	}

	if filter.Email != "" {
		query["email"] = filter.Email
	}

	if filter.Status != 0 {
		query["status"] = filter.Status
	}

	if filter.Role != 0 {
		query["role"] = filter.Role
	}

	// 可以添加其他可能的查询条件

	return query
}
