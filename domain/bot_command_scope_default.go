package domain

// BotCommandScopeDefault Represents the default scope of bot commands.
// Default commands are used if no commands with a narrower scope are specified for the user.
type BotCommandScopeDefault struct {
	// Scope type, must be default.
	Type string `json:"type"`
}
