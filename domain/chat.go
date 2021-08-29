package domain

// Chat This object represents a chat.
type Chat struct {
	// The most recent pinned message (by sending date).
	// Returned only in getChat.
	//
	// Optional.
	PinnedMessage *Message `json:"pinned_message"`
	// Default chat member permissions, for groups and supergroups.
	// Returned only in getChat.
	//
	// Optional.
	Permissions *ChatPermissions `json:"permissions"`
	// The time after which all messages sent to the chat will be automatically deleted; in seconds.
	// Returned only in getChat.
	//
	// Optional.
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
	// For supergroups, name of group sticker set.
	// Returned only in getChat.
	//
	// Optional.
	StickerSetName string `json:"sticker_set_name"`
	// Unique identifier for this chat.
	// This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	Id int `json:"id"`
	// Username, for private chats, supergroups and channels if available.
	//
	// Optional.
	Username string `json:"username"`
	// Primary invite link, for groups, supergroups and channel chats.
	// Returned only in getChat.
	//
	// Optional.
	InviteLink string `json:"invite_link"`
	// For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user.
	// Returned only in getChat.
	//
	// Optional.
	SlowModeDelay int `json:"slow_mode_delay"`
	// True, if the bot can change the group sticker set.
	// Returned only in getChat.
	//
	// Optional.
	CanSetStickerSet bool `json:"can_set_sticker_set"`
	// Unique identifier for the linked chat, i.e. the discussion group identifier for a channel and vice versa; for supergroups and channel chats.
	// This identifier may be greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it.
	// But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for storing this identifier.
	// Returned only in getChat.
	//
	// Optional.
	LinkedChatId int `json:"linked_chat_id"`
	// First name of the other party in a private chat.
	//
	// Optional.
	FirstName string `json:"first_name"`
	// Last name of the other party in a private chat.
	//
	// Optional.
	LastName string `json:"last_name"`
	// Description, for groups, supergroups and channel chats.
	// Returned only in getChat.
	//
	// Optional.
	Description string `json:"description"`
	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”.
	Type string `json:"type"`
	// Title, for supergroups, channels and group chats.
	//
	// Optional.
	Title string `json:"title"`
	// Chat photo.
	// Returned only in getChat.
	//
	// Optional.
	Photo *ChatPhoto `json:"photo"`
	// Bio of the other party in a private chat.
	// Returned only in getChat.
	//
	// Optional.
	Bio string `json:"bio"`
	// For supergroups, the location to which the supergroup is connected.
	// Returned only in getChat.
	//
	// Optional.
	Location *ChatLocation `json:"location"`
}
