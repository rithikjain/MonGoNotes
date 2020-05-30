package user

import (
	"context"
	"github.com/rithikjain/MongoNotes/pkg"
	"github.com/rithikjain/MongoNotes/pkg/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

type Repository interface {
	FindByID(id string) (*entities.User, error)

	Register(user *entities.User) (*entities.User, error)

	DoesEmailExist(email string) (bool, error)

	FindByEmail(email string) (*entities.User, error)
}

type repo struct {
	Coll *mongo.Collection
}

func NewRepo(coll *mongo.Collection) {
	return &repo{
		Coll: coll,
	}
}

func (r *repo) FindByID(id string) (*entities.User, error) {
	user := &entities.User{}
	objID, _ := primitive.ObjectIDFromHex(id)
	err := r.Coll.FindOne(context.Background(), bson.M{"_id": objID}).Decode(user)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return user, nil
}

func (r *repo) Register(user *entities.User) (*entities.User, error) {
	user.ID = primitive.NewObjectID()
	_, err := r.Coll.InsertOne(context.Background(), user)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return user, nil
}

func (r *repo) DoesEmailExist(email string) (bool, error) {
	count, err := r.Coll.CountDocuments(context.Background(), bson.M{"email": email})
	if err != nil {
		return true, pkg.ErrDatabase
	}
	if count == 0 {
		return false, nil
	}
	return true, nil
}

func (r *repo) FindByEmail(email string) (*entities.User, error) {
	user := &entities.User{}
	err := r.Coll.FindOne(context.Background(), bson.M{"email": email}).Decode(user)
	if err != nil {
		return nil, pkg.ErrDatabase
	}
	return user, nil
}
