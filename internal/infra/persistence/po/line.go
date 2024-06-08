package po

import "go.mongodb.org/mongo-driver/bson/primitive"

type Line struct {
	LineID    primitive.ObjectID `bson:"line_id"`
	Name      string             `bson:"name"`
	Type      string             `bson:"type"`
	UserID    primitive.ObjectID `bson:"user_id"`
	NetworkID primitive.ObjectID `bson:"network_id"`
	CreatedAt int64              `bson:"created_at"`
	UpdatedAt int64              `bson:"updated_at"`
	DeletedAt int64              `bson:"deleted_at"`
}
