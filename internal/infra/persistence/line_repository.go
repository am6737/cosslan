package persistence

import (
	"context"
	"github.com/cossteam/cosslan/internal/domain/entity"
	"github.com/cossteam/cosslan/internal/domain/filter"
	"go.mongodb.org/mongo-driver/mongo"
)

type LineRepository struct {
	client     *mongo.Client
	db         *mongo.Database
	collection *mongo.Collection
}

const (
	lineCollectionName = "lines"
)

func NewLineRepository(client *mongo.Client, dbName string) *LineRepository {
	db := client.Database(dbName)
	collection := db.Collection(lineCollectionName)
	m := &LineRepository{
		client:     client,
		db:         db,
		collection: collection,
	}
	return m
}

func (l *LineRepository) Create(ctx context.Context, line *entity.Line) error {
	//TODO implement me
	panic("implement me")
}

func (l *LineRepository) Get(ctx context.Context, id string) (*entity.Line, error) {
	//TODO implement me
	panic("implement me")
}

func (l *LineRepository) Update(ctx context.Context, line *entity.Line) error {
	//TODO implement me
	panic("implement me")
}

func (l *LineRepository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func (l *LineRepository) Find(ctx context.Context, filter *filter.LineFilter) ([]*entity.Line, error) {
	//TODO implement me
	panic("implement me")
}
