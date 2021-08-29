package domain

// BotCommandScopeChat Represents the scope of bot commands, covering a specific chat.
type BotCommandScopeChat struct {
	// Scope type, must be a chat.
	Type string `json:"type"`
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatId string `json:"chat_id"`
}
