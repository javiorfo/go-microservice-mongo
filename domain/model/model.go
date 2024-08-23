package model

import (
	"github.com/javiorfo/go-microservice-lib/auditory"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Dummy represents a dada structure
type Dummy struct {
	ID   primitive.ObjectID `json:"id"`
	Info string             `json:"info"`
	auditory.Auditable
}
