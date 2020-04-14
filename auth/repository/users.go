package repository

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func FindUser(database *mongo.Database) FindUserFunc {
	/** Initialize collection **/
	collection := database.Collection("users")
	return func(email string, password string) (*User, error) {
		result := User{}
		filter := bson.D{{email, password}}

		/** Return an error if not user found **/
		err := collection.FindOne(context.Background(), filter).Decode(&result)
		if err != nil {
			return nil, fmt.Errorf("User %s not matching email or password", email)
		}
		return &result, nil
	}

}
