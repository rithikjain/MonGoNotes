package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Title string             `bson:"title" json:"title"`
	Body  string             `bson:"body,omitempty" json:"body,omitempty"`
}
