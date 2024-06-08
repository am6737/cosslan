package converter

import (
	"fmt"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/infra/persistence/po"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToEntityNode(n *po.Node) *entity.Node {
	return &entity.Node{
		NodeID:    n.NodeID.String(),
		Name:      n.Name,
		IP:        n.IP,
		UserID:    n.UserID.String(),
		NetworkID: n.NetworkID.String(),
		Status:    entity.NodeStatus(n.Status),
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
		DeletedAt: n.DeletedAt,
	}
}

func ToPONode(p *entity.Node) *po.Node {
	id, err := primitive.ObjectIDFromHex(p.NodeID)
	if err != nil {
		fmt.Println("error：", err)
		return nil
	}

	if id == primitive.NilObjectID {
		id = primitive.NewObjectID()
	}

	userID, err := primitive.ObjectIDFromHex(p.UserID)
	if err != nil {
		fmt.Println("error：", err)
		return nil
	}

	networkID, err := primitive.ObjectIDFromHex(p.NetworkID)
	if err != nil {
		fmt.Println("error：", err)
		return nil
	}

	return &po.Node{
		NodeID:    id,
		Name:      p.Name,
		IP:        p.IP,
		PublicIP:  p.PublicIP,
		Port:      p.Port,
		UserID:    userID,
		Status:    uint8(p.Status),
		NetworkID: networkID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
}
