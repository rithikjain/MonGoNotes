package note

import (
	"context"
	"github.com/rithikjain/MongoNotes/pkg"
	"github.com/rithikjain/MongoNotes/pkg/entities"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	CreateNote(note *entities.Note) (*entities.Note, error)
}

type repo struct {
	Coll *mongo.Collection
}

func NewRepo(coll *mongo.Collection) Repository {
	return &repo{
		Coll: coll,
	}
}

func (r *repo) CreateNote(note *entities.Note) (*entities.Note, error) {
	_, err := r.Coll.InsertOne(context.Background(), note)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return note, nil
}
