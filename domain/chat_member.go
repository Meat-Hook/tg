package domain

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Errors.
var (
	// ErrUnknownChatMemberKind unknown ChatMemberKind value.
	ErrUnknownChatMemberKind = errors.New("unknown ChatMemberKind")
)

var (
	chatMemberKind                  = ChatMemberKind(0)
	_              json.Marshaler   = chatMemberKind
	_              json.Unmarshaler = &chatMemberKind
)

// ChatMemberKind enum.
type ChatMemberKind uint8

// UnmarshalJSON implements json.Unmarshaler.
func (i *ChatMemberKind) UnmarshalJSON(bytes []byte) error {
	str := string(bytes)

	var res ChatMemberKind
	switch str {
	case Creator.String():
		res = Creator
	case Administrator.String():
		res = Administrator
	case Member.String():
		res = Member
	case Restricted.String():
		res = Restricted
	case Left.String():
		res = Left
	case Kicked.String():
		res = Kicked
	default:
		return fmt.Errorf("%w: %s", ErrUnknownChatMemberKind, str)
	}

	*i = res

	return nil
}

// MarshalJSON implements json.Marshaler.
func (i ChatMemberKind) MarshalJSON() ([]byte, error) {
	return []byte(i.String()), nil
}

const (
	_             ChatMemberKind = iota
	Creator                      // creator
	Administrator                // administrator
	Member                       // member
	Restricted                   // restricted
	Left                         // left
	Kicked                       // kicked
)

// ChatMember this object contains information about one member of a chat.
type ChatMember struct {
	// User information about the user.
	User User `json:"user"`
	// Status the member's status in the chat.
	Status ChatMemberKind `json:"status"`
	// CustomTitle owner and administrators only. Custom title for this user.
	CustomTitle string `json:"custom_title"`
	// IsAnonymous owner and administrators only.
	// True, if the user's presence in the chat is hidden.
	IsAnonymous bool `json:"is_anonymous"`
	// CanBeEdited administrators only.
	// True, if the bot is allowed to edit administrator privileges of that user.
	CanBeEdited bool `json:"can_be_edited"`
	// CanManageChat administrators only.
	// True, if the administrator can access the chat event log, chat statistics, message statistics in channels,
	// see channel members, see anonymous administrators in supergroups and ignore slow mode.
	// Implied by any other administrator privilege
	CanManageChat bool `json:"can_manage_chat"`
	// CanPostMessages administrators only.
	// True, if the administrator can post in the channel; channels only.
	CanPostMessages bool `json:"can_post_messages"`
	// CanEditMessages administrators only.
	// True, if the administrator can edit messages of other users and can pin messages; channels only.
	CanEditMessages bool `json:"can_edit_messages"`
	// CanDeleteMessages administrators only.
	// True, if the administrator can delete messages of other users.
	CanDeleteMessages bool `json:"can_delete_messages"`
	// CanManageVoiceChats administrators only.
	// True, if the administrator can manage voice chats.
	CanManageVoiceChats bool `json:"can_manage_voice_chats"`
	// CanRestrictMembers administrators only.
	// True, if the administrator can restrict, ban or unban chat members.
	CanRestrictMembers bool `json:"can_restrict_members"`
	// CanPromoteMembers administrators only.
	// True, if the administrator can add new administrators with a subset of
	// their own privileges or demote administrators that he has promoted, directly or indirectly
	// (promoted by administrators that were appointed by the user).
	CanPromoteMembers bool `json:"can_promote_members"`
	// CanChangeInfo administrators and restricted only.
	// True, if the user is allowed to change the chat title, photo and other settings.
	CanChangeInfo bool `json:"can_change_info"`
	// CanInviteUsers administrators and restricted only.
	// True, if the user is allowed to invite new users to the chat.
	CanInviteUsers bool `json:"can_invite_users"`
	// CanPinMessages administrators and restricted only.
	// True, if the user is allowed to pin messages; groups and supergroups only.
	CanPinMessages bool `json:"can_pin_messages"`
	// IsMember restricted only.
	// True, if the user is a member of the chat at the moment of the request.
	IsMember bool `json:"is_member"`
	// CanSendMessages restricted only.
	// True, if the user is allowed to send text messages, contacts, locations and venues.
	CanSendMessages bool `json:"can_send_messages"`
	// CanSendMediaMessages restricted only.
	// True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes.
	CanSendMediaMessages bool `json:"can_send_media_messages"`
	// CanSendPolls restricted only.
	// True, if the user is allowed to send polls.
	CanSendPolls bool `json:"can_send_polls"`
	// CanSendOtherMessages restricted only.
	// True, if the user is allowed to send animations, games, stickers and use inline bots.
	CanSendOtherMessages bool `json:"can_send_other_messages"`
	// CanAddWebPagePreviews restricted only.
	// True, if the user is allowed to add web page previews to their messages.
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	// UntilDate restricted and kicked only.
	// Date when restrictions will be lifted for this user; unix time
	UntilDate int `json:"until_date"`
}
