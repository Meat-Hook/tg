package domain

// BotCommandScopeAllGroupChats Represents the scope of bot commands, covering all group and supergroup chats.
type BotCommandScopeAllGroupChats struct {
	// Scope type, must be all_group_chats.
	Type string `json:"type"`
}
