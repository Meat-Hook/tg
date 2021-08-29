package domain

// ChatPermissions Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// True, if the user is allowed to send animations, games, stickers and use inline bots, implies can_send_media_messages.
	//
	// Optional.
	CanSendOtherMessages bool `json:"can_send_other_messages"`
	// True, if the user is allowed to add web page previews to their messages, implies can_send_media_messages.
	//
	// Optional.
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	// True, if the user is allowed to change the chat title, photo and other settings.
	// Ignored in public supergroups.
	//
	// Optional.
	CanChangeInfo bool `json:"can_change_info"`
	// True, if the user is allowed to invite new users to the chat.
	//
	// Optional.
	CanInviteUsers bool `json:"can_invite_users"`
	// True, if the user is allowed to pin messages.
	// Ignored in public supergroups.
	//
	// Optional.
	CanPinMessages bool `json:"can_pin_messages"`
	// True, if the user is allowed to send text messages, contacts, locations and venues.
	//
	// Optional.
	CanSendMessages bool `json:"can_send_messages"`
	// True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes, implies can_send_messages.
	//
	// Optional.
	CanSendMediaMessages bool `json:"can_send_media_messages"`
	// True, if the user is allowed to send polls, implies can_send_messages.
	//
	// Optional.
	CanSendPolls bool `json:"can_send_polls"`
}
