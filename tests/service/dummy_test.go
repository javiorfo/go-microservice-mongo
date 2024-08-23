package dummy_test

import (
	"errors"
	"testing"

	"github.com/javiorfo/go-microservice-lib/pagination"
	"github.com/javiorfo/go-microservice/domain/model"
	"github.com/javiorfo/go-microservice/domain/service"
	"github.com/javiorfo/go-microservice/tests/mocks"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestFindById(t *testing.T) {
	mockRepo := new(mocks.MockDummyRepository)
	dummyService := service.NewDummyService(mockRepo)

	id := "1"
	expectedDummy := &model.Dummy{ID: primitive.NewObjectID()}

	mockRepo.On("FindById", id).Return(expectedDummy, nil)

	result, err := dummyService.FindById(id)

	assert.NoError(t, err)
	assert.Equal(t, expectedDummy, result)
	mockRepo.AssertExpectations(t)
}

func TestFindByIdNotFound(t *testing.T) {
	mockRepo := new(mocks.MockDummyRepository)
	dummyService := service.NewDummyService(mockRepo)

	id := "1"

	mockRepo.On("FindById", id).Return(nil, errors.New("not found"))

	result, err := dummyService.FindById(id)

	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(mocks.MockDummyRepository)
	dummyService := service.NewDummyService(mockRepo)

	page := pagination.Page{Page: 1, Size: 10, SortBy: "id", SortOrder: "asc"}
	expectedDummies := []model.Dummy{
		{ID: primitive.NewObjectID()},
		{ID: primitive.NewObjectID()},
	}

	mockRepo.On("FindAll", page).Return(expectedDummies, nil)

	result, err := dummyService.FindAll(page)

	assert.NoError(t, err)
	assert.Equal(t, expectedDummies, result)
	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := new(mocks.MockDummyRepository)
	dummyService := service.NewDummyService(mockRepo)

	newDummy := &model.Dummy{ID: primitive.NewObjectID()}

	mockRepo.On("Create", newDummy).Return(nil)

	err := dummyService.Create(newDummy)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}