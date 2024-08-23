package model

import (
	"github.com/javiorfo/go-microservice-lib/auditory"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Dummy represents a dada structure
type Dummy struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Info string             `json:"info" bson:"info"`
	auditory.Auditable
}
