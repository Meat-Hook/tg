package domain

// ChatMemberAdministrator Represents a chat member that has some additional privileges.
type ChatMemberAdministrator struct {
	// True, if the bot is allowed to edit administrator privileges of that user.
	CanBeEdited bool `json:"can_be_edited"`
	// True, if the user's presence in the chat is hidden.
	IsAnonymous bool `json:"is_anonymous"`
	// True, if the administrator can access the chat event log, chat statistics, message statistics in channels, see channel members, see anonymous administrators in supergroups and ignore slow mode.
	// Implied by any other administrator privilege.
	CanManageChat bool `json:"can_manage_chat"`
	// True, if the administrator can manage voice chats.
	CanManageVoiceChats bool `json:"can_manage_voice_chats"`
	// True, if the user is allowed to pin messages; groups and supergroups only.
	//
	// Optional.
	CanPinMessages bool `json:"can_pin_messages"`
	// Custom title for this user.
	//
	// Optional.
	CustomTitle string `json:"custom_title"`
	// Information about the user.
	User User `json:"user"`
	// True, if the administrator can restrict, ban or unban chat members.
	CanRestrictMembers bool `json:"can_restrict_members"`
	// True, if the user is allowed to change the chat title, photo and other settings.
	CanChangeInfo bool `json:"can_change_info"`
	// True, if the user is allowed to invite new users to the chat.
	CanInviteUsers bool `json:"can_invite_users"`
	// True, if the administrator can edit messages of other users and can pin messages; channels only.
	//
	// Optional.
	CanEditMessages bool `json:"can_edit_messages"`
	// True, if the administrator can delete messages of other users.
	CanDeleteMessages bool `json:"can_delete_messages"`
	// True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that he has promoted, directly or indirectly (promoted by administrators that were appointed by the user).
	CanPromoteMembers bool `json:"can_promote_members"`
	// The member's status in the chat, always “administrator”.
	Status string `json:"status"`
	// True, if the administrator can post in the channel; channels only.
	//
	// Optional.
	CanPostMessages bool `json:"can_post_messages"`
}

func (c ChatMemberAdministrator) isChatMember() {}
