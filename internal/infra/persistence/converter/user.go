package converter

import (
	"fmt"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/infra/persistence/po"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToEntityUser(u *po.User) *entity.User {

	return &entity.User{
		UserID:    u.UserID.String(),
		Name:      u.Name,
		Password:  u.Password,
		Email:     u.Email,
		Status:    entity.UserStatus(u.Status),
		Role:      entity.UserRole(u.Role),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}

func ToPOUser(p *entity.User) *po.User {

	userID, err := primitive.ObjectIDFromHex(p.UserID)
	if err != nil {
		fmt.Println("error：", err)
		return nil
	}
	if userID == primitive.NilObjectID {
		userID = primitive.NewObjectID()
	}

	return &po.User{
		UserID:    userID,
		Name:      p.Name,
		Password:  p.Password,
		Email:     p.Email,
		Status:    uint8(p.Status),
		Role:      uint8(p.Role),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
}
