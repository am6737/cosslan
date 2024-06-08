package po

import "go.mongodb.org/mongo-driver/bson/primitive"

type Network struct {
	NetworkID  primitive.ObjectID `bson:"network_id,omitempty"`
	Address    string             `bson:"address"`
	SubnetMask string             `bson:"subnet_mask"`
	UserID     primitive.ObjectID `bson:"user_id,omitempty"`
	DeletedAt  int64              `bson:"deleted_at"`
	CreatedAt  int64              `bson:"created_at"`
	UpdatedAt  int64              `bson:"updated_at"`
}
