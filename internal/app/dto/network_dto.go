package dto

type CreateNetworkDTO struct {
	Address    string `json:"address"`
	SubnetMask string `json:"subnet_mask"`
	UserID     string `json:"user_id"`
}
