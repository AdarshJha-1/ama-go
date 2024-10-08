package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at"`
}
