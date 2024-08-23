package config

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/javiorfo/go-microservice-lib/env"
	"github.com/javiorfo/go-microservice-lib/security"
	"github.com/javiorfo/go-microservice-mongo/internal/database"
)

// Keycloak configuration
var KeycloakConfig = security.KeycloakConfig{
	Keycloak:     gocloak.NewClient(env.GetEnvOrDefault("KEYCLOAK_HOST", "http://localhost:8081")),
	Realm:        "javi",
	ClientID:     "srv-client",
	ClientSecret: env.GetEnvOrDefault("KEYCLOAK_CLIENT_SECRET", "RqaTlO0d2OnBbeRuImNnbLWm5yZL66Mo"),
	Enabled:      true,
}

var DBDataConnection = database.DBDataConnection{
	Host:     env.GetEnvOrDefault("DB_HOST", "localhost"),
	Port:     env.GetEnvOrDefault("DB_PORT", "27017"),
	DBName:   env.GetEnvOrDefault("DB_NAME", "db_dummy"),
	User:     env.GetEnvOrDefault("DB_USER", "admin"),
	Password: env.GetEnvOrDefault("DB_PASSWORD", "admin"),
}

// App configuration
const AppName = "go-microservice-mongo"
const AppPort = ":8080"
const AppContextPath = "/app"

// Tracing server configuration
var TracingHost = env.GetEnvOrDefault("TRACING_HOST", "http://localhost:4318")

// Swagger configuration
var SwaggerEnabled = env.GetEnvOrDefault("SWAGGER_ENABLED", "true")
