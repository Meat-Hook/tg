package domain

// ChatMemberBanned Represents a chat member that was banned in the chat and can't return to the chat or view chat messages.
type ChatMemberBanned struct {
	// The member's status in the chat, always “kicked”.
	Status string `json:"status"`
	// Information about the user.
	User User `json:"user"`
	// Date when restrictions will be lifted for this user; unix time.
	// If 0, then the user is banned forever.
	UntilDate int `json:"until_date"`
}

func (c ChatMemberBanned) isChatMember() {}
