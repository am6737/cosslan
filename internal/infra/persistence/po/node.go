package po

import "go.mongodb.org/mongo-driver/bson/primitive"

type Node struct {
	NodeID    primitive.ObjectID `bson:"node_id"`
	Name      string             `bson:"name"`
	IP        string             `bson:"ip"`
	PublicIP  string             `bson:"public_ip"`
	Port      uint16             `bson:"port"`
	UserID    primitive.ObjectID `bson:"user_id"`
	NetworkID primitive.ObjectID `bson:"network_id"`
	Status    uint8              `bson:"status"`
	CreatedAt int64              `bson:"created_at"`
	UpdatedAt int64              `bson:"updated_at"`
	DeletedAt int64              `bson:"deleted_at"`
}
