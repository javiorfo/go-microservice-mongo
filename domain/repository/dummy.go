package repository

import (
	"context"
	"errors"
	"time"

	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice-lib/tracing"
	"github.com/javiorfo/go-microservice-mongo/domain/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type DummyRepository interface {
	FindById(context.Context, string) (*model.Dummy, error)
	FindAll(context.Context, pagination.Page) ([]model.Dummy, error)
	Create(context.Context, *model.Dummy) error
}

type dummyRepository struct {
	*mongo.Collection
	tracer trace.Tracer
}

func NewDummyRepository(collection *mongo.Collection) DummyRepository {
	return &dummyRepository{Collection: collection, tracer: otel.Tracer(tracing.Name())}
}

func (repository *dummyRepository) FindById(ctx context.Context, id string) (*model.Dummy, error) {
	ctx, span := repository.tracer.Start(ctx, tracing.Name())
	defer span.End()

	var dummy model.Dummy
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = repository.FindOne(ctx, bson.M{"_id": objID}).Decode(&dummy)
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

func (repository *dummyRepository) FindAll(ctx context.Context, page pagination.Page) ([]model.Dummy, error) {
	ctx, span := repository.tracer.Start(ctx, tracing.Name())
	defer span.End()

	var dummies []model.Dummy
	cursor, err := repository.Find(ctx, bson.D{}, getPagination(page))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
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

func (repository *dummyRepository) Create(ctx context.Context, d *model.Dummy) error {
	ctx, span := repository.tracer.Start(ctx, tracing.Name())
	defer span.End()

	d.ID = primitive.NewObjectID()
	d.CreateDate = time.Now().UTC()

	_, err := repository.InsertOne(ctx, d)
	if err != nil {
		return err
	}
	return nil
}
