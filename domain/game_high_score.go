package domain

// GameHighScore This object represents one row of the high scores table for a game.
// And that's about all we've got for now.
// If you've got any questions, please check out our Bot FAQ ».
type GameHighScore struct {
	// Position in high score table for the game.
	Position int `json:"position"`
	// User.
	User User `json:"user"`
	// Score.
	Score int `json:"score"`
}
