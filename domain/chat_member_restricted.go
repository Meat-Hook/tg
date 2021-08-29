package domain

// ChatMemberRestricted Represents a chat member that is under certain restrictions in the chat.
// Supergroups only.
type ChatMemberRestricted struct {
	// True, if the user is a member of the chat at the moment of the request.
	IsMember bool `json:"is_member"`
	// True, if the user is allowed to change the chat title, photo and other settings.
	CanChangeInfo bool `json:"can_change_info"`
	// True, if the user is allowed to invite new users to the chat.
	CanInviteUsers bool `json:"can_invite_users"`
	// True, if the user is allowed to send polls.
	CanSendPolls bool `json:"can_send_polls"`
	// True, if the user is allowed to add web page previews to their messages.
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	// Information about the user.
	User User `json:"user"`
	// True, if the user is allowed to pin messages.
	CanPinMessages bool `json:"can_pin_messages"`
	// True, if the user is allowed to send text messages, contacts, locations and venues.
	CanSendMessages bool `json:"can_send_messages"`
	// True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes.
	CanSendMediaMessages bool `json:"can_send_media_messages"`
	// True, if the user is allowed to send animations, games, stickers and use inline bots.
	CanSendOtherMessages bool `json:"can_send_other_messages"`
	// Date when restrictions will be lifted for this user; unix time.
	// If 0, then the user is restricted forever.
	UntilDate int `json:"until_date"`
	// The member's status in the chat, always “restricted”.
	Status string `json:"status"`
}

func (c ChatMemberRestricted) isChatMember() {}
