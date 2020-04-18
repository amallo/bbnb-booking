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
	return func(email string, password string) (*UserRow, error) {
		res, err := collection.InsertOne(context.Background(), bson.M{"email": email, "password": password})
		if err != nil {
			return nil, &RepositoryError{Message: fmt.Sprintf("Cannot create user %s", email), Cause: err}
		}
		newUser := UserRow{
			ID:       res.InsertedID.(primitive.ObjectID),
			Email:    email,
			Password: password,
		}
		return &newUser, nil
	}
}
