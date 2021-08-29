package domain

// ChatMemberMember Represents a chat member that has no additional privileges or restrictions.
type ChatMemberMember struct {
	// The member's status in the chat, always “member”.
	Status string `json:"status"`
	// Information about the user.
	User User `json:"user"`
}

func (c ChatMemberMember) isChatMember() {}
