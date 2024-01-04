package controllers

import (
	"github.com/go-playground/validator/v10"

	"revise-it/backend/database"

	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()
