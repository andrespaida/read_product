package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
    ID          primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
    Name        string             `json:"name" bson:"name"`
    Description string             `json:"description" bson:"description"`
    Price       float64            `json:"price" bson:"price"`
    Stock       int                `json:"stock" bson:"stock"`
    ImageURL    string             `json:"image_url" bson:"image_url"`
    CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}