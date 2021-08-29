package domain

// BotCommandScopeChatAdministrators Represents the scope of bot commands, covering all administrators of a specific group or supergroup chat.
type BotCommandScopeChatAdministrators struct {
	// Scope type, must be chat_administrators.
	Type string `json:"type"`
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatId string `json:"chat_id"`
}
