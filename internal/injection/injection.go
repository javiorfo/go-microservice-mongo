package injection

import (
	"github.com/gofiber/fiber/v2"
	"github.com/javiorfo/go-microservice/api/routes"
	"github.com/javiorfo/go-microservice/config"
	"github.com/javiorfo/go-microservice/domain/repository"
	"github.com/javiorfo/go-microservice/domain/service"
	"github.com/javiorfo/go-microservice/internal/database"
)

func Inject(api fiber.Router) {
	// MongoDB
	db := database.DBinstance

	// Dummy: Repository, Servicer and Routes
    dummyCollection := db.Collection("dummies")
	dummyRepository := repository.NewDummyRepository(dummyCollection)
	dummyService := service.NewDummyService(dummyRepository)
	routes.Dummy(api, config.KeycloakConfig, dummyService)
}
