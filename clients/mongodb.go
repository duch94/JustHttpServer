package clients

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

var (
	mongoHost    = "localhost"
	mongoPort    = "27017"
	mongoAddress = "mongodb://" + mongoHost + ":" + mongoPort
)

// SendDocument is function for sending document doc to DB dbName and collection collName
func SendDocument(dbName string, collName string, doc map[string]interface{}) (interface{}, error) {
	client, err := mongo.NewClient(mongoAddress)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	collection := client.Database(dbName).Collection(collName)
	res, err := collection.InsertOne(ctx, bson.M(doc))
	if err != nil {
		return nil, err
	}

	err = client.Disconnect(ctx)
	if err != nil {
		return nil, err
	}
	return res.InsertedID, nil
}

// GetDocumentByLogin is function for finding users by user login
func GetDocumentByLogin(dbName string, collName string, login string) (interface{}, error) {
	client, err := mongo.NewClient(mongoAddress)
	if err != nil {
		return nil, err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	filter := make(map[string]interface{})
	filter["login"] = login
	collection := client.Database(dbName).Collection(collName)
	res := collection.FindOne(ctx, bson.M(filter))

	err = client.Disconnect(ctx)
	if err != nil {
		return nil, err
	}

	var resDoc interface{}
	res.Decode(resDoc)

	return resDoc, nil
}