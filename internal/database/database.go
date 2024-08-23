package database

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DBDataConnection struct {
	Host     string
	Port     string
	DBName   string
	User     string
	Password string
}

var DBinstance *mongo.Database

func (db DBDataConnection) Connect() (context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	dsn := fmt.Sprintf("mongodb://%s:%s@%s:%s",
		db.User,
		db.Password,
		db.Host,
		db.Port)

	clientOptions := options.Client().ApplyURI(dsn)

	client, err := mongo.Connect(ctx, clientOptions)
	mongoDB := client.Database(db.DBName)
	if err != nil {
		cancel()
		return nil, err
	}

	log.Info("Connected to MongoDB!")
	DBinstance = mongoDB
	return cancel, nil
}
