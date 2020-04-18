package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUser(database *mongo.Database) FindUserFunc {
	/** Initialize collection **/
	collection := database.Collection("users")
	return func(criteria Criteria) (*UserRow, error) {
		result := UserRow{}
		filter := bson.M{}
		for key, value := range criteria {
			filter[key] = value
		}

		err := collection.FindOne(context.Background(), filter).Decode(&result)
		if err != nil {
			return nil, &RepositoryError{Message: fmt.Sprintf("User not found"), Cause: err}
		}
		return &result, nil
	}

}
func InsertUser(database *mongo.Database) InsertUserFunc {
	/** Initialize collection **/
	collection := database.Collection("users")
	return func(criteria Criteria) (*UserRow, error) {
		parameters := bson.M{}
		for key, value := range criteria {
			parameters[key] = value
		}
		res, err := collection.InsertOne(context.Background(), parameters)
		if err != nil {
			return nil, &RepositoryError{Message: fmt.Sprintf("Cannot create user"), Cause: err}
		}
		newUser := UserRow{
			ID:       res.InsertedID.(primitive.ObjectID),
			Email:    parameters["email"].(string),
			Password: parameters["password"].(string),
		}
		return &newUser, nil
	}
}
