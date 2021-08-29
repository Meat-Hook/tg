package domain

// BotCommandScopeAllChatAdministrators Represents the scope of bot commands, covering all group and supergroup chat administrators.
type BotCommandScopeAllChatAdministrators struct {
	// Scope type, must be all_chat_administrators.
	Type string `json:"type"`
}
