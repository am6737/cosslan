package converter

import (
	"fmt"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/infra/persistence/po"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToEntityLine(l *po.Line) *entity.Line {

	return &entity.Line{
		LineID:    l.LineID.String(),
		Name:      l.Name,
		Type:      l.Type,
		UserID:    l.UserID.String(),
		CreatedAt: l.CreatedAt,
		UpdatedAt: l.UpdatedAt,
		DeletedAt: l.DeletedAt,
		NetworkID: l.NetworkID.String(),
	}
}

func ToPOLine(p *entity.Line) *po.Line {
	lid, err := primitive.ObjectIDFromHex(p.LineID)
	if err != nil {
		fmt.Println("error：", err)
		return nil
	}
	if lid == primitive.NilObjectID {
		lid = primitive.NewObjectID()
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

	return &po.Line{
		LineID:    lid,
		Name:      p.Name,
		Type:      p.Type,
		UserID:    userID,
		NetworkID: networkID,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
}
