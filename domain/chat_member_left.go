package domain

// ChatMemberLeft Represents a chat member that isn't currently a member of the chat, but may join it themselves.
type ChatMemberLeft struct {
	// The member's status in the chat, always “left”.
	Status string `json:"status"`
	// Information about the user.
	User User `json:"user"`
}

func (c ChatMemberLeft) isChatMember() {}
