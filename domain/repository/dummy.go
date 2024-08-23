package repository

import (
	"context"
	"errors"
	"time"

	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice-mongo/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	cursor, err := repository.Find(context.Background(), bson.D{}, getPagination(page))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.TODO()) {
		var dummy model.Dummy
		_ = cursor.Decode(&dummy)
		dummies = append(dummies, dummy)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return dummies, nil
}

func getPagination(page pagination.Page) *options.FindOptions {
    order := 1
    if page.SortOrder == "desc" {
        order = -1
    }

    sort := bson.D{{Key: page.SortBy, Value: order}}
	return options.Find().
		SetSort(sort).
		SetSkip(int64(page.Page-1) * int64(page.Size)).
		SetLimit(int64(page.Size))
}

func (repository *dummyRepository) Create(d *model.Dummy) error {
	d.ID = primitive.NewObjectID()
	d.CreateDate = time.Now().UTC()

	_, err := repository.InsertOne(context.Background(), d)
	if err != nil {
		return err
	}
	return nil
}
