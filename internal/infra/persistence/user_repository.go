package persistence

import (
	"context"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/domain/filter"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	userCollectionName = "users"
)

func NewUserRepository(client *mongo.Client, dbName string) *UserRepository {
	db := client.Database(dbName)
	collection := db.Collection(userCollectionName)
	m := &UserRepository{
		client:     client,
		db:         db,
		collection: collection,
	}
	return m
}

func (u *UserRepository) Create(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Get(ctx context.Context, id string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Find(ctx context.Context, filter *filter.UserFilter) ([]*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Update(ctx context.Context, user *entity.User) error {
	//TODO implement me
	panic("implement me")
}
