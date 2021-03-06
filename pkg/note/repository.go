package note

import (
	"context"
	"github.com/rithikjain/MongoNotes/pkg"
	"github.com/rithikjain/MongoNotes/pkg/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Repository interface {
	CreateNote(note *entities.Note) (*entities.Note, error)

	GetAllNotes() (*[]entities.Note, error)

	UpdateNote(note *entities.Note) error

	DeleteNote(id string) error
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
	note.ID = primitive.NewObjectID()
	note.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	note.UpdatedAt = note.CreatedAt
	_, err := r.Coll.InsertOne(context.Background(), note)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return note, nil
}

func (r *repo) GetAllNotes() (*[]entities.Note, error) {
	var notes []entities.Note
	cursor, err := r.Coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	for cursor.Next(context.TODO()) {
		var note entities.Note
		_ = cursor.Decode(&note)
		notes = append(notes, note)
	}
	return &notes, nil
}

func (r *repo) UpdateNote(note *entities.Note) error {
	note.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	_, err := r.Coll.UpdateOne(context.Background(), bson.M{"_id": note.ID}, bson.M{"$set": note})
	if err != nil {
		return pkg.ErrDatabase
	}
	return nil
}

func (r *repo) DeleteNote(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return pkg.ErrDatabase
	}
	_, err = r.Coll.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		return pkg.ErrDatabase
	}
	return nil
}
