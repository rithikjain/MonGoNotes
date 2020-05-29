package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title string             `bson:"title,omitempty" json:"title,omitempty"`
	Body  string             `bson:"body,omitempty" json:"body,omitempty"`
}
