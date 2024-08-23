package repository

import (
	"context"
	"errors"

	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type DummyRepository interface {
	FindById(id string) (*model.Dummy, error)
	FindAll(pagination.Page) ([]model.Dummy, error)
	Create(*model.Dummy) error
}

type dummyRepository struct {
	*mongo.Collection
}

func NewDummyRepository(collection *mongo.Collection) DummyRepository {
	return &dummyRepository{collection}
}

func (repository *dummyRepository) FindById(id string) (*model.Dummy, error) {
	var dummy model.Dummy

    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, err
    }

	err = repository.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&dummy)
    if err != nil {
        if err == mongo.ErrNoDocuments {
		    return nil, errors.New("Dummy not found")
        } else {
            return nil, err
        }
    } else {
        return &dummy, nil
    }
}

func (repository *dummyRepository) FindAll(page pagination.Page) ([]model.Dummy, error) {
	var dummies []model.Dummy

    // TODO
	return dummies, nil
}

func (repository *dummyRepository) Create(d *model.Dummy) error {
    d.ID = primitive.NewObjectID()
    // TODO auditory

	_, err := repository.InsertOne(context.Background(), d)
	if err != nil {
		return err
	}
	return nil
}
