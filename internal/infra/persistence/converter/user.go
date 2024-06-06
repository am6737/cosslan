package converter

import (
	"fmt"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/infra/persistence/po"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ToPOUser(u *entity.User) *po.User {
	return &po.User{
		UserID:    u.UserID.String(),
		Name:      u.Name,
		Password:  u.Password,
		Email:     u.Email,
		Status:    uint8(u.Status),
		Role:      uint8(u.Role),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}
}

func ToEntityUser(p *po.User) *entity.User {
	userID, err := primitive.ObjectIDFromHex(p.UserID)
	if err != nil {
		fmt.Println("error：", err)
		return nil
	}

	return &entity.User{
		UserID:    userID,
		Name:      p.Name,
		Password:  p.Password,
		Email:     p.Email,
		Status:    entity.UserStatus(p.Status),
		Role:      entity.UserRole(p.Role),
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
		DeletedAt: p.DeletedAt,
	}
}
