package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CreatedAt primitive.DateTime `bson:"created_at,omitempty" json:"created_at,omitempty"`
	Title     string             `bson:"title,omitempty" json:"title,omitempty"`
	Body      string             `bson:"body,omitempty" json:"body,omitempty"`
}
