package persistence

import (
	"context"
	"fmt"
	"github.com/cossteam/cosslan/config"
	"github.com/cossteam/cosslan/internal/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"sync"
)

var (
	repos *Repositories
	once  sync.Once
)

type Repositories struct {
	UserRepo    repository.UserRepository
	LineRepo    repository.LineRepository
	NodeRepo    repository.NodeRepository
	NetworkRepo repository.NetworkRepository
	Db          *mongo.Client
}

func newRepositories(db *mongo.Client, dbName string) *Repositories {
	return &Repositories{
		Db:          db,
		UserRepo:    NewUserRepository(db, dbName),
		LineRepo:    NewLineRepository(db, dbName),
		NodeRepo:    NewNodeRepository(db, dbName),
		NetworkRepo: NewNetworkRepository(db, dbName),
	}
}

func GetRepositories(cfg config.Config) *Repositories {
	once.Do(func() {
		fmt.Println("初始化数据库")
		mongoDB, err := connectMongoDB(cfg.Persistence.Url)
		if err != nil {
			panic(err)
		}
		repos = newRepositories(mongoDB, cfg.Persistence.DB)
	})
	return repos
}

func connectMongoDB(uri string) (*mongo.Client, error) {
	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the server to confirm a successful connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	return client, nil
}
