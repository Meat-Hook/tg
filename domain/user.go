package domain

// User This object represents a Telegram user or bot.
type User struct {
	// User's or bot's first name.
	FirstName string `json:"first_name"`
	// User's or bot's last name.
	//
	// Optional.
	LastName string `json:"last_name"`
	// User's or bot's username.
	//
	// Optional.
	Username string `json:"username"`
	// True, if the bot can be invited to groups.
	// Returned only in getMe.
	//
	// Optional.
	CanJoinGroups bool `json:"can_join_groups"`
	// True, if privacy mode is disabled for the bot.
	// Returned only in getMe.
	//
	// Optional.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages"`
	// Unique identifier for this user or bot.
	// This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	Id int `json:"id"`
	// True, if this user is a bot.
	IsBot bool `json:"is_bot"`
	// IETF language tag of the user's language.
	//
	// Optional.
	LanguageCode string `json:"language_code"`
	// True, if the bot supports inline queries.
	// Returned only in getMe.
	//
	// Optional.
	SupportsInlineQueries bool `json:"supports_inline_queries"`
}
