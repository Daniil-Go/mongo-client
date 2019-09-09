package client

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DefaultDatabase is name of database
const DefaultDatabase = "clientbase"

//CollectionName is name of collection
const CollectionName = "client"

//MongoHandler is type of mongoHandler struct
type MongoHandler struct {
	client   *mongo.Client
	database string
}

//NewHandler is MongoHandler Constructor
func NewHandler(address string) *MongoHandler {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cl, _ := mongo.Connect(ctx, options.Client().ApplyURI(address))
	mh := &MongoHandler{
		client:   cl,
		database: DefaultDatabase,
	}
	return mh
}

//GetOne will get one element from collection
func (mh *MongoHandler) GetOne(c *Client, filter interface{}) error {
	//Will automatically create a collection if not available
	collection := mh.client.Database(mh.database).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	err := collection.FindOne(ctx, filter).Decode(c)
	return err
}

//Get will get all elements from collection
func (mh *MongoHandler) Get(filter interface{}) []*Client {
	collection := mh.client.Database(mh.database).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var result []*Client
	for cur.Next(ctx) {
		client := &Client{}
		er := cur.Decode(client)
		if er != nil {
			log.Fatal(er)
		}
		result = append(result, client)
	}
	return result
}

//AddOne will add one element to collection
func (mh *MongoHandler) AddOne(c *Client) (*mongo.InsertOneResult, error) {
	collection := mh.client.Database(mh.database).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, c)
	return result, err
}

//Update will update information in collection
func (mh *MongoHandler) Update(c *Client, filter interface{}, update interface{}) (*mongo.UpdateResult, error) {
	collection := mh.client.Database(mh.database).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.UpdateMany(ctx, filter, update)
	return result, err
}

//RemoveOne will remove one element from collection
func (mh *MongoHandler) RemoveOne(filter interface{}) (*mongo.DeleteResult, error) {
	collection := mh.client.Database(mh.database).Collection(CollectionName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	result, err := collection.DeleteOne(ctx, filter)
	return result, err
}
