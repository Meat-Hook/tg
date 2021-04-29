package tg

import "encoding/json"

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

	// Contains data from the Telegram API with the result.
	response struct {
		Ok          bool                `json:"ok"`
		Result      json.RawMessage     `json:"result"`
		ErrorCode   int                 `json:"error_code"`
		Description string              `json:"description"`
		Parameters  *responseParameters `json:"parameters"`
	}
	responseParameters struct {
		MigrateToChatID int64 `json:"migrate_to_chat_id"` // optional
		RetryAfter      int   `json:"retry_after"`        // optional
	}
)
