package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type FindUserByEmailAndPasswordFunc = func(email string, password string) (user *UserRow, err error)
type InsertUserFunc = func(email string, password string) (user *UserRow, err error)
type UserRow struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}
