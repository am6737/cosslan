package converter

import (
	"fmt"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/infra/persistence/po"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToEntityNetwork(e *po.Network) *entity.Network {
	// 将 entity.Network 转换为 po.Network
	return &entity.Network{
		NetworkID:  e.NetworkID.String(),
		Address:    e.Address,
		SubnetMask: e.SubnetMask,
		UserID:     e.UserID.String(),
		DeletedAt:  e.DeletedAt,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
	}

}

func ToPONetwork(p *entity.Network) *po.Network {
	networkID, err := primitive.ObjectIDFromHex(p.NetworkID)
	if err != nil {
		fmt.Println("error：", err)
		return nil
	}

	if networkID == primitive.NilObjectID {
		networkID = primitive.NewObjectID()
	}

	userID, err := primitive.ObjectIDFromHex(p.UserID)
	if err != nil {
		fmt.Println("error：", err)
		return nil
	}

	return &po.Network{
		NetworkID:  networkID,
		Address:    p.Address,
		SubnetMask: p.SubnetMask,
		UserID:     userID,
		DeletedAt:  p.DeletedAt,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}
