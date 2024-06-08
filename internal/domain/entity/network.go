package entity

type Network struct {
	NetworkID  string `bson:"network_id,omitempty"`
	Address    string `bson:"address"`
	SubnetMask string `bson:"subnet_mask"`
	UserID     string `bson:"user_id,omitempty"`
	DeletedAt  int64  `bson:"deleted_at"`
	CreatedAt  int64  `bson:"created_at"`
	UpdatedAt  int64  `bson:"updated_at"`
}
