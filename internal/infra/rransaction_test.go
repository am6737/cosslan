package infra

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"testing"
)

// TransactionFunc 定义操作函数类型
type TransactionFunc func(sessCtx mongo.SessionContext) error

// ExecuteTransaction 执行事务
func ExecuteTransaction(client *mongo.Client, operations ...TransactionFunc) error {
	// 开始一个会话
	session, err := client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(context.TODO())

	// 定义事务函数
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		for _, operation := range operations {
			if err := operation(sessCtx); err != nil {
				return nil, err
			}
		}
		return nil, nil
	}

	// 执行事务
	_, err = session.WithTransaction(context.TODO(), callback)
	return err
}

// 示例操作1
func operation1(sessCtx mongo.SessionContext) error {
	collection := sessCtx.Client().Database("testdb").Collection("testcollection")
	_, err := collection.InsertOne(sessCtx, bson.D{{"name", "Alice"}})
	return err
}

// 示例操作2
func operation2(sessCtx mongo.SessionContext) error {
	collection := sessCtx.Client().Database("testdb").Collection("testcollection")
	_, err := collection.InsertOne(sessCtx, bson.D{{"name", "Bob"}})
	return err
}

func TestTransaction(t *testing.T) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// 执行事务
	err = ExecuteTransaction(client, operation1, operation2)
	if err != nil {
		log.Fatalf("Transaction failed: %v", err)
	} else {
		fmt.Println("Transaction succeeded")
	}
}
