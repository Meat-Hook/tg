package domain

// ChatMemberOwner Represents a chat member that owns the chat and has all administrator privileges.
type ChatMemberOwner struct {
	// The member's status in the chat, always “creator”.
	Status string `json:"status"`
	// Information about the user.
	User User `json:"user"`
	// True, if the user's presence in the chat is hidden.
	IsAnonymous bool `json:"is_anonymous"`
	// Custom title for this user.
	//
	// Optional.
	CustomTitle string `json:"custom_title"`
}

func (c ChatMemberOwner) isChatMember() {}
