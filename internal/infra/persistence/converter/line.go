package converter

import (
	"fmt"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/infra/persistence/po"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToPOLine(l *entity.Line) *po.Line {
	return &po.Line{
		ID:         l.ID.String(),
		Name:       l.Name,
		Type:       l.Type,
		Subnets:    l.Subnets,
		UserID:     l.UserID.String(),
		NextNodeID: l.NextNodeID.String(),
		CreatedAt:  l.CreatedAt,
		UpdatedAt:  l.UpdatedAt,
		DeletedAt:  l.DeletedAt,
	}
}

func ToEntityLine(p *po.Line) *entity.Line {
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

	nextNodeID, err := primitive.ObjectIDFromHex(p.NextNodeID)
	if err != nil {
		fmt.Println("error：", err)
		return nil
	}

	return &entity.Line{
		ID:         id,
		Name:       p.Name,
		Type:       p.Type,
		Subnets:    p.Subnets,
		UserID:     userID,
		NextNodeID: nextNodeID,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
		DeletedAt:  p.DeletedAt,
	}
}
