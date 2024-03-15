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

var deckCollection *mongo.Collection = database.OpenCollection(database.Client, "deck")

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
		fmt.Println("Made it to the end")

		// resultInsertionNumber, insertErr := userCollection.InsertOne(ctx, user)
		// if insertErr != nil {
		// 	msg := fmt.Sprintf("Deck item was not created")
		// 	c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		// 	return
		// }
		defer cancel()
	}
}

// func for deleting deck

// get all decks for a user
