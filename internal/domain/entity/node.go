package entity

type Node struct {
	NodeID    string     `bson:"node_id,omitempty"`
	Name      string     `bson:"name"`
	IP        string     `bson:"ip"`
	UserID    string     `bson:"user_id,omitempty"`
	NetworkID string     `bson:"network_id,omitempty"`
	PublicIP  string     `bson:"public_ip"`
	Port      uint16     `bson:"port"`
	Status    NodeStatus `bson:"status"`
	CreatedAt int64      `bson:"created_at"`
	UpdatedAt int64      `bson:"updated_at"`
	DeletedAt int64      `bson:"deleted_at"`
}

type NodeStatus uint8

const (
	Online  NodeStatus = iota + 1 //在线
	Offline                       // 离线
)
