package controllers

import (
	"context"
	"revise-it/backend/models"
	"time"

	"github.com/gin-gonic/gin"
)

func addDeck() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		// add deck to db with ref to user
		var deck models.Deck
	}
}

// func for deleting deck

// get all decks for a user
