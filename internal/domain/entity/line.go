package entity

type Line struct {
	LineID     string `bson:"line_id,omitempty"`
	Name       string `bson:"name"`
	Type       string `bson:"type"`
	UserID     string `bson:"user_id,omitempty"`
	NextNodeID string `bson:"next_node_id,omitempty"`
	NetworkID  string `bson:"network_id,omitempty"`
	CreatedAt  int64  `bson:"created_at"`
	UpdatedAt  int64  `bson:"updated_at"`
	DeletedAt  int64  `bson:"deleted_at"`
}

// 判断是否还有下一个节点
func (l *Line) HasNextNode() bool {
	return l.NextNodeID != ""
}
