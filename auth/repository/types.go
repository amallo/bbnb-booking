package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type FindUserFunc = func(criteria Criteria) (user *UserRow, err error)
type InsertUserFunc = func(email string, password string) (user *UserRow, err error)
type UserRow struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

type Criteria map[string]interface{}
