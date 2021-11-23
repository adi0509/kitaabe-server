package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var Context context.Context

func Close(client *mongo.Client, ctx context.Context,
	cancel context.CancelFunc) {

	defer cancel()

	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func Connect(uri string) (*mongo.Client, context.Context,
	context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)
	cnt, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	Client = cnt
	Context = ctx
	return cnt, ctx, cancel, err
}

func AddIndex(dbName, collection, id string) error {
	serviceCollection := Client.Database(dbName).Collection(collection)
	indexName, err := serviceCollection.Indexes().CreateOne(mtest.Background, mongo.IndexModel{
		Keys:    bson.D{{Key: id, Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}
	fmt.Println("index name: " + indexName)
	return nil
}

func InsertOne(dataBase, col string, doc interface{}) (*mongo.InsertOneResult, error) {
	collection := Client.Database(dataBase).Collection(col)
	result, err := collection.InsertOne(context.TODO(), doc)
	return result, err
}

func UpdateOne(dataBase, col string, filter, update interface{}) (result *mongo.UpdateResult, err error) {
	collection := Client.Database(dataBase).Collection(col)
	result, err = collection.UpdateOne(context.TODO(), filter, update)
	return
}

func Query(dataBase, col string, query, field interface{}) (result *mongo.Cursor, err error) {
	collection := Client.Database(dataBase).Collection(col)

	result, err = collection.Find(context.TODO(), query, options.Find().SetProjection(field))
	return
}
