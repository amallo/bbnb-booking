package repository

import "go.mongodb.org/mongo-driver/mongo"

type RepositoryError struct {
	Cause   error
	Message string
}

func (e *RepositoryError) Error() string {
	return e.Message
}

func (e *RepositoryError) ErrNoDocuments() bool {
	return e.Cause == mongo.ErrNoDocuments
}
