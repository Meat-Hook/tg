package domain

// InlineQuery This object represents an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	// Offset of the results to be returned, can be controlled by the bot.
	Offset string `json:"offset"`
	// Type of the chat, from which the inline query was sent.
	// Can be either “sender” for a private chat with the inline query sender, “private”, “group”, “supergroup”, or “channel”.
	// The chat type should be always known for requests sent from official clients and most third-party clients, unless the request was sent from a secret chat.
	//
	// Optional.
	ChatType string `json:"chat_type"`
	// Sender location, only for bots that request user location.
	//
	// Optional.
	Location *Location `json:"location"`
	// Unique identifier for this query.
	Id string `json:"id"`
	// Sender.
	From User `json:"from"`
	// Text of the query (up to 256 characters).
	Query string `json:"query"`
}
