package po

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserID    primitive.ObjectID `bson:"user_id,omitempty"`
	Name      string             `bson:"name"`
	Password  string             `bson:"password"`
	Email     string             `bson:"email"`
	Status    uint8              `bson:"status"`
	Role      uint8              `bson:"role"`
	CreatedAt int64              `bson:"created_at"`
	UpdatedAt int64              `bson:"updated_at"`
	DeletedAt int64              `bson:"deleted_at"`
}
