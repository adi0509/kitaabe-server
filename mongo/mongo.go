package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
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
	_, err := serviceCollection.Indexes().CreateOne(mtest.Background, mongo.IndexModel{
		Keys:    bson.D{{Key: id, Value: 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}
	return nil
}

func AddTextIndexItem(dbName, collection string) error {
	coll := Client.Database(dbName).Collection(collection)
	index := []mongo.IndexModel{

		{
			Keys: bsonx.Doc{
				{Key: "item_name", Value: bsonx.String("text")},
				{Key: "item_description", Value: bsonx.String("text")},
				{Key: "university", Value: bsonx.String("text")},
				{Key: "available_in_city", Value: bsonx.String("text")}},
			Options: options.Index().SetWeights(bson.D{
				{"item_name", 5},
				{"description", 3},
			}),
		},
	}
	opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
	_, err := coll.Indexes().CreateMany(context.Background(), index, opts)
	if err != nil {
		fmt.Println(err)
		return err
	}
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
