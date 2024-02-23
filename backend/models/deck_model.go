package models

// deck represents a deck of flashcards.
type Deck struct {
	ID         primitive.ObjectID `bson:"_id"`
	Name       string
	Description string
	user_id     string
	Flashcards []Flashcard
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
}
