package domain

// BotCommandScopeAllPrivateChats Represents the scope of bot commands, covering all private chats.
type BotCommandScopeAllPrivateChats struct {
	// Scope type, must be all_private_chats.
	Type string `json:"type"`
}
