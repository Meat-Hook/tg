package domain

// ChatMemberUpdated This object represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	// Chat the user belongs to.
	Chat Chat `json:"chat"`
	// Performer of the action, which resulted in the change.
	From User `json:"from"`
	// Date the change was done in Unix time.
	Date int `json:"date"`
	// Previous information about the chat member.
	OldChatMember ChatMember `json:"old_chat_member"`
	// New information about the chat member.
	NewChatMember ChatMember `json:"new_chat_member"`
	// Chat invite link, which was used by the user to join the chat; for joining by invite link events only.
	//
	// Optional.
	InviteLink *ChatInviteLink `json:"invite_link"`
}
