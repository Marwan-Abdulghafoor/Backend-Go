package dao

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"murwan.net/fiephrs-backend/utils"
)

// Connect to MongoDB
func ConnectDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Insert a profile into the collection
func InsertProfile(client *mongo.Client, profile *utils.ProfileInfo) error {
	collection := client.Database("apsd").Collection("patients")
	_, err := collection.InsertOne(context.Background(), profile)
	return err
}

func FindProfileById(client *mongo.Client, id int) (*utils.ProfileInfo, error) {
	collection := client.Database("apsd").Collection("patients")
	var result utils.ProfileInfo
	err := collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

// Update a profile
func UpdateProfile(client *mongo.Client, id int, updateProfile *utils.ProfileInfo) error {
	collection := client.Database("apsd").Collection("patients")
	_, err := collection.UpdateOne(context.Background(), bson.M{"id": id}, bson.M{"$set": updateProfile})
	return err
}

// Delete a profile by first name
func DeleteProfile(client *mongo.Client, firstName string) error {
	collection := client.Database("test").Collection("profileInfo")
	_, err := collection.DeleteOne(context.Background(), bson.M{"firstName": firstName})
	return err
}
