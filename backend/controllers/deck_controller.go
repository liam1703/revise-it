package controllers

import (
	"context"
	"fmt"
	"revise-it/backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func AddDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// add deck to db with ref to user
		var deck models.Deck
		fmt.Println(ctx, "deck", deck)
		if err := c.BindJSON(&deck); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(deck)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			return
		}

		deck.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		deck.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		deck.ID = primitive.NewObjectID()

		fmt.Println(ctx, "deck", deck)

		defer cancel()
	}
}

// func for deleting deck

// get all decks for a user
