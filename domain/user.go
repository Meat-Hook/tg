package domain

// User represents a Telegram user or bot.
type User struct {
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
	// CanJoinGroups true, if the bot can be invited to groups.
	// Returned only in getMe.
	CanJoinGroups bool `json:"can_join_groups"`
	// CanReadAllGroupMessages true, if privacy mode is disabled for the bot.
	// Returned only in getMe.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages"`
	// SupportsInlineQueries true, if the bot supports inline queries.
	// Returned only in getMe.
	SupportsInlineQueries bool `json:"supports_inline_queries"`
}
