package dummy_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/javiorfo/go-microservice-mongo/domain/model"
	"github.com/javiorfo/go-microservice-mongo/domain/repository"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var repo repository.DummyRepository

func TestMain(m *testing.M) {
	ctx := context.Background()

	req := testcontainers.ContainerRequest{
		Image:        "mongo:latest",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor:   wait.ForListeningPort("27017/tcp"),
	}

	mongoContainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		log.Fatalf("Failed to start container: %s", err)
	}
	defer mongoContainer.Terminate(ctx)

	host, err := mongoContainer.Host(ctx)
	if err != nil {
		log.Fatalf("Failed to get container host: %s", err)
	}
	port, err := mongoContainer.MappedPort(ctx, "27017")
	if err != nil {
		log.Fatalf("Failed to get container port: %s", err)
	}

	clientOptions := options.Client().ApplyURI("mongodb://" + host + ":" + port.Port())
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %s", err)
	}
	defer client.Disconnect(ctx)

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %s", err)
	}
	collection := client.Database("testdb").Collection("dummies")

	repo = repository.NewDummyRepository(collection)

	// Run the tests
	code := m.Run()

	os.Exit(code)
}

func TestDummy(t *testing.T) {

	dummyRecord := model.Dummy{Info: "testname"}

	ctx := context.Background()
	if err := repo.Create(ctx, &dummyRecord); err != nil {
		t.Fatalf("Failed to insert record: %v", err)
	}

	dummy, err := repo.FindById(ctx, dummyRecord.ID.Hex())
	if err != nil {
		t.Fatalf("Failed to query record: %v", err)
	}

	if dummy.Info != "testname" {
		t.Errorf("Expected name to be 'testname', got '%s'", dummy.Info)
	}
}
