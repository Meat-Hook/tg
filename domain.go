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
	// ChatPhoto this object represents a chat photo.
	ChatPhoto struct {
		// SmallFileID file identifier of small (160x160) chat photo.
		// This file_id can be used only for photo download and only for as long as the photo is not changed.
		SmallFileID string `json:"small_file_id"`
		// SmallFileUniqueID unique file identifier of small (160x160) chat photo,
		// which is supposed to be the same over time and for different bots.
		// Can't be used to download or reuse the file.
		SmallFileUniqueID string `json:"small_file_unique_id"`
		// BigFileID File identifier of big (640x640) chat photo.
		// This file_id can be used only for photo download and only for as long as the photo is not changed.
		BigFileID string `json:"big_file_id"`
		// BigFileUniqueID unique file identifier of big (640x640) chat photo,
		// which is supposed to be the same over time and for different bots.
		// Can't be used to download or reuse the file.
		BigFileUniqueID string `json:"big_file_unique_id"`
	}
)
