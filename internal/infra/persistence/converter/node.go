package converter

import (
	"fmt"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/infra/persistence/po"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToPONode(n *entity.Node) *po.Node {
	return &po.Node{
		ID:        n.ID.String(),
		Name:      n.Name,
		IP:        n.IP,
		UserID:    n.UserID.String(),
		Location:  n.Location,
		CreatedAt: n.CreatedAt,
		UpdatedAt: n.UpdatedAt,
		DeletedAt: n.DeletedAt,
	}
}

func ToEntityNode(p *po.Node) *entity.Node {
	id, err := primitive.ObjectIDFromHex(p.ID)
	if err != nil {
		fmt.Println("error：", err)
		return nil
	}

	userID, err := primitive.ObjectIDFromHex(p.UserID)
	if err != nil {
		fmt.Println("error：", err)
		return nil
	}

	return &entity.Node{
		ID:        id,
		Name:      p.Name,
		IP:        p.IP,
		UserID:    userID,
		Location:  p.Location,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
}
