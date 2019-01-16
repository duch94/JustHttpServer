package clients

import (
	"context"
	"time"

	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// MongoClient is struct with methods for working with MongoDB
type MongoClient struct {
	mongoHost string
	mongoPort string
	client *mongo.Client
	ctx context.Context
}

// NewMongoClient is constructor of MongoClient object
func NewMongoClient(mongoHost string, mongoPort string) (*MongoClient, error) {
	var mc MongoClient
	mongoAddress := "mongodb://" + mongoHost + ":" + mongoPort
	var err error
	mc.client, err = mongo.NewClient(mongoAddress)
	if err != nil {
		return nil, err
	}

	// context must be created from request and go through all the abstraction 
	mc.ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	err = mc.client.Connect(mc.ctx)
	if err != nil {
		return nil, err
	}
	return &mc, nil
}

// Disconnect is function for disconnecting from db
func (mc *MongoClient) Disconnect() error {
	err := mc.client.Disconnect(mc.ctx)
	if err != nil {
		return err
	}
	return nil
}

// SendDocument is function for sending document doc to DB dbName and collection collName
func (mc *MongoClient) SendDocument(ctx context.Context, dbName string, collName string, doc map[string]interface{}) (interface{}, error) {
	collection := mc.client.Database(dbName).Collection(collName)
	res, err := collection.InsertOne(ctx, bson.M(doc))
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

// GetDocumentByLogin is function for finding users by user login
func (mc *MongoClient) GetDocumentByLogin(ctx context.Context, dbName string, collName string, login string) (interface{}, error) {
	filter := make(map[string]interface{})
	filter["login"] = login
	collection := mc.client.Database(dbName).Collection(collName)
	res := collection.FindOne(ctx, bson.M(filter))

	var resDoc interface{}
	err := res.Decode(&resDoc)
	if err != nil {
		return nil, err
	}

	return resDoc, nil
}

// UpdateDocumentByLogin is function for updating user by login
func (mc *MongoClient) UpdateDocumentByLogin(ctx context.Context, dbName string, collName string, login string) (interface{}, error) {
	return nil, nil
}

// DeleteDocumentByLogin is function for deleting user by login
func (mc *MongoClient) DeleteDocumentByLogin(ctx context.Context, dbName string, collName string, login string) (interface{}, error) {
	return nil, nil
}

// GetAllDocuments is function for getting all the user documents
func (mc *MongoClient) GetAllDocuments(ctx context.Context, dbName string, collName string) ([]interface{}, error) {
	filter := make(map[string]interface{})
	collection := mc.client.Database(dbName).Collection(collName)
	cursor, err := collection.Find(ctx, bson.M(filter))
	if err != nil {
		return nil, err
	}

	userDoc := struct {
		Login string
		Dob string
		Name string
		Password string
	}{}
	var resultSlice []interface{}
	for {
		if cursor.Next(ctx) {
			err = cursor.Decode(&userDoc)
			if err != nil {
				return nil, err
			}
			resultSlice = append(resultSlice, userDoc)
		} else {
			break
		}
	}
	
	return resultSlice, nil
}