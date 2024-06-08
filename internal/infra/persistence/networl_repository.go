package persistence

import (
	"context"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/domain/filter"
	"github.com/cossteam/cosslan/internal/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ repository.NetworkRepository = &NetworkRepository{}

type NetworkRepository struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	networkCollectionName = "networks"
)

func NewNetworkRepository(client *mongo.Client, dbName string) *NetworkRepository {
	db := client.Database(dbName)
	collection := db.Collection(networkCollectionName)
	m := &NetworkRepository{
		client:     client,
		db:         db,
		collection: collection,
	}
	return m
}

func (n *NetworkRepository) Update(ctx context.Context, network *entity.Network) error {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepository) Get(ctx context.Context, id string) {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepository) Create(ctx context.Context, network *entity.Network) error {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (n *NetworkRepository) Find(ctx context.Context, filter *filter.NetworkFilter) ([]*entity.Network, error) {
	//TODO implement me
	panic("implement me")
}
