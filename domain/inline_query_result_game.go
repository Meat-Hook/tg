package domain

// InlineQueryResultGame Represents a Game.
// Note: This will only work in Telegram versions released after October 1, 2016.
// Older clients will not display any inline results if a game result is among them.
type InlineQueryResultGame struct {
	// Type of the result, must be game.
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
	// Short name of the game.
	GameShortName string `json:"game_short_name"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}
