package tg

type (
	// User represents a Telegram user or bot.
	User struct {
		// ID is a ID for user.
		ID int `json:"id"`
		// FirstName user's or bot's first name.
		FirstName string `json:"first_name"`
		// LastName user's or bot's last name.
		LastName string `json:"last_name"`
		// UserName user's or bot's username.
		UserName string `json:"username"`
		// LanguageCode IETF language.
		// info: https://en.wikipedia.org/wiki/IETF_language_tag
		LanguageCode string `json:"language_code"`
		// IsBot true, if this user is a bot.
		IsBot bool `json:"is_bot"`
	}
)
