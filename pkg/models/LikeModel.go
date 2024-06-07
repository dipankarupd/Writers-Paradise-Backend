package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	ID        primitive.ObjectID `bson:"_id"`
	Comment   primitive.ObjectID `bson:"comment,omitempty" json:"comment"`
	LikedBy   primitive.ObjectID `bson:"liked_by,omitempty" json:"liked_by"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
