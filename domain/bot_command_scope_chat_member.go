package domain

// BotCommandScopeChatMember Represents the scope of bot commands, covering a specific member of a group or supergroup chat.
type BotCommandScopeChatMember struct {
	// Scope type, must be chat_member.
	Type string `json:"type"`
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatId string `json:"chat_id"`
	// Unique identifier of the target user.
	UserId int `json:"user_id"`
}
