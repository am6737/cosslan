package persistence

import (
	"context"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/domain/filter"
	"go.mongodb.org/mongo-driver/mongo"
)

type NodeRepository struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	nodeCollectionName = "nodes"
)

func NewNodeRepository(client *mongo.Client, dbName string) *NodeRepository {
	db := client.Database(dbName)
	collection := db.Collection(nodeCollectionName)
	m := &NodeRepository{
		client:     client,
		db:         db,
		collection: collection,
	}
	return m
}

func (n *NodeRepository) Find(ctx context.Context, filter *filter.NodeFilter) ([]*entity.Node, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NodeRepository) Create(ctx context.Context, node *entity.Node) error {
	//TODO implement me
	panic("implement me")
}

func (n *NodeRepository) Get(ctx context.Context, id string) (*entity.Node, error) {
	//TODO implement me
	panic("implement me")
}

func (n *NodeRepository) Update(ctx context.Context, node *entity.Node) error {
	//TODO implement me
	panic("implement me")
}

func (n *NodeRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}
