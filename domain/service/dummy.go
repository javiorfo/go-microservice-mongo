package service

import (
	"context"

	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice-lib/tracing"
	"github.com/javiorfo/go-microservice-mongo/domain/model"
	"github.com/javiorfo/go-microservice-mongo/domain/repository"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type DummyService interface {
	FindById(context.Context, string) (*model.Dummy, error)
	FindAll(context.Context, pagination.Page) ([]model.Dummy, error)
	Create(context.Context, *model.Dummy) error
}

type dummyService struct {
	repository repository.DummyRepository
	tracer     trace.Tracer
}

func NewDummyService(r repository.DummyRepository) DummyService {
	return &dummyService{
		repository: r,
		tracer:     otel.Tracer(tracing.Name()),
	}
}

func (service *dummyService) FindById(ctx context.Context, id string) (*model.Dummy, error) {
	_, span := service.tracer.Start(ctx, tracing.Name())
	defer span.End()

	return service.repository.FindById(ctx, id)
}

func (service *dummyService) FindAll(ctx context.Context, page pagination.Page) ([]model.Dummy, error) {
	_, span := service.tracer.Start(ctx, tracing.Name())
	defer span.End()

	return service.repository.FindAll(ctx, page)
}

func (service *dummyService) Create(ctx context.Context, d *model.Dummy) error {
	_, span := service.tracer.Start(ctx, tracing.Name())
	defer span.End()

	return service.repository.Create(ctx, d)
}
