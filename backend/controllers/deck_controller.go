package controllers

import (
	"context"
	"fmt"
	"net/http"
	"revise-it/backend/database"
	"revise-it/backend/models"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var deckCollection *mongo.Collection = database.OpenCollection(database.Client, "decks")

func AddDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// add deck to db with ref to user
		var deck models.Deck
		fmt.Println(ctx, "deck", deck)
		if err := c.BindJSON(&deck); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error line 22": err.Error()})
			return
		}
		fmt.Println("inside adddeck")
		validationErr := validate.Struct(deck)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		deck.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		deck.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		deck.ID = primitive.NewObjectID()

		fmt.Println(ctx, "deck", deck)
		fmt.Println("Made it to the end", deckCollection)

		_, insertErr := deckCollection.InsertOne(ctx, deck)
		if insertErr != nil {
			msg := "Deck item was not created"
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg, "insertErr": insertErr.Error()})
			return
		}
		defer cancel()
	}
}

// get all decks for a user

// func for deleting deck


