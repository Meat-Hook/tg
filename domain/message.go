package domain

// Message This object represents a message.
type Message struct {
	// For replies, the original message.
	// Note that the Message object in this field will not contain further reply_to_message fields even if it itself is a reply.
	//
	// Optional.
	ReplyToMessage *Message `json:"reply_to_message"`
	// Caption for the animation, audio, document, photo, video or voice, 0-1024 characters.
	//
	// Optional.
	Caption string `json:"caption"`
	// New members that were added to the group or supergroup and information about them (the bot itself may be one of these members).
	//
	// Optional.
	NewChatMembers []User `json:"new_chat_members"`
	// A chat title was changed to this value.
	//
	// Optional.
	NewChatTitle string `json:"new_chat_title"`
	// Service message: voice chat started.
	//
	// Optional.
	VoiceChatStarted *VoiceChatStarted `json:"voice_chat_started"`
	// Date the message was last edited in Unix time.
	//
	// Optional.
	EditDate int `json:"edit_date"`
	// Message is a photo, available sizes of the photo.
	//
	// Optional.
	Photo []PhotoSize `json:"photo"`
	// Message is a shared contact, information about the contact.
	//
	// Optional.
	Contact *Contact `json:"contact"`
	// Specified message was pinned.
	// Note that the Message object in this field will not contain further reply_to_message fields even if it is itself a reply.
	//
	// Optional.
	PinnedMessage *Message `json:"pinned_message"`
	// Telegram Passport data.
	//
	// Optional.
	PassportData *PassportData `json:"passport_data"`
	// Conversation the message belongs to.
	Chat Chat `json:"chat"`
	// For messages forwarded from channels or from anonymous administrators, information about the original sender chat.
	//
	// Optional.
	ForwardFromChat *Chat `json:"forward_from_chat"`
	// For messages forwarded from channels, signature of the post author if present.
	//
	// Optional.
	ForwardSignature string `json:"forward_signature"`
	// For text messages, the actual UTF-8 text of the message, 0-4096 characters.
	//
	// Optional.
	Text string `json:"text"`
	// Message is a video, information about the video.
	//
	// Optional.
	Video *Video `json:"video"`
	// For forwarded messages, sender of the original message.
	//
	// Optional.
	ForwardFrom *User `json:"forward_from"`
	// For forwarded messages, date the original message was sent in Unix time.
	//
	// Optional.
	ForwardDate int `json:"forward_date"`
	// For messages with a caption, special entities like usernames, URLs, bot commands, etc.
	// that appear in the caption.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Sender of the message, sent on behalf of a chat.
	// The channel itself for channel messages.
	// The supergroup itself for messages from anonymous group administrators.
	// The linked channel for messages automatically forwarded to the discussion group.
	//
	// Optional.
	SenderChat *Chat `json:"sender_chat"`
	// Message is a game, information about the game.
	// More about games ».
	//
	// Optional.
	Game *Game `json:"game"`
	// The supergroup has been migrated from a group with the specified identifier.
	// This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	//
	// Optional.
	MigrateFromChatId int `json:"migrate_from_chat_id"`
	// Unique message identifier inside this chat.
	MessageId int `json:"message_id"`
	// Signature of the post author for messages in channels, or the custom title of an anonymous group administrator.
	//
	// Optional.
	AuthorSignature string `json:"author_signature"`
	// Message is a general file, information about the file.
	//
	// Optional.
	Document *Document `json:"document"`
	// Message is a dice with random value.
	//
	// Optional.
	Dice *Dice `json:"dice"`
	// Message is an invoice for a payment, information about the invoice.
	// More about payments ».
	//
	// Optional.
	Invoice *Invoice `json:"invoice"`
	// Service message: voice chat ended.
	//
	// Optional.
	VoiceChatEnded *VoiceChatEnded `json:"voice_chat_ended"`
	// For messages forwarded from channels, identifier of the original message in the channel.
	//
	// Optional.
	ForwardFromMessageId int `json:"forward_from_message_id"`
	// Sender's name for messages forwarded from users who disallow adding a link to their account in forwarded messages.
	//
	// Optional.
	ForwardSenderName string `json:"forward_sender_name"`
	// Message is an audio file, information about the file.
	//
	// Optional.
	Audio *Audio `json:"audio"`
	// Message is a sticker, information about the sticker.
	//
	// Optional.
	Sticker *Sticker `json:"sticker"`
	// Service message: the chat photo was deleted.
	//
	// Optional.
	DeleteChatPhoto bool `json:"delete_chat_photo"`
	// Service message: the channel has been created.
	// This field can't be received in a message coming through updates, because bot can't be a member of a channel when it is created.
	// It can only be found in reply_to_message if someone replies to a very first message in a channel.
	//
	// Optional.
	ChannelChatCreated bool `json:"channel_chat_created"`
	// Message is an animation, information about the animation.
	// For backward compatibility, when this field is set, the document field will also be set.
	//
	// Optional.
	Animation *Animation `json:"animation"`
	// Message is a video note, information about the video message.
	//
	// Optional.
	VideoNote *VideoNote `json:"video_note"`
	// A member was removed from the group, information about them (this member may be the bot itself).
	//
	// Optional.
	LeftChatMember *User `json:"left_chat_member"`
	// The domain name of the website on which the user has logged in.
	// More about Telegram Login ».
	//
	// Optional.
	ConnectedWebsite string `json:"connected_website"`
	// Service message: voice chat scheduled.
	//
	// Optional.
	VoiceChatScheduled *VoiceChatScheduled `json:"voice_chat_scheduled"`
	// Service message: new participants invited to a voice chat.
	//
	// Optional.
	VoiceChatParticipantsInvited *VoiceChatParticipantsInvited `json:"voice_chat_participants_invited"`
	// A chat photo was change to this value.
	//
	// Optional.
	NewChatPhoto []PhotoSize `json:"new_chat_photo"`
	// Message is a venue, information about the venue.
	// For backward compatibility, when this field is set, the location field will also be set.
	//
	// Optional.
	Venue *Venue `json:"venue"`
	// Service message: the group has been created.
	//
	// Optional.
	GroupChatCreated bool `json:"group_chat_created"`
	// Service message: auto-delete timer settings changed in the chat.
	//
	// Optional.
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	// Inline keyboard attached to the message.
	// login_url buttons are represented as ordinary url buttons.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// Date the message was sent in Unix time.
	Date int `json:"date"`
	// For text messages, special entities like usernames, URLs, bot commands, etc.
	// that appear in the text.
	//
	// Optional.
	Entities []MessageEntity `json:"entities"`
	// Message is a voice message, information about the file.
	//
	// Optional.
	Voice *Voice `json:"voice"`
	// Bot through which the message was sent.
	//
	// Optional.
	ViaBot *User `json:"via_bot"`
	// Message is a native poll, information about the poll.
	//
	// Optional.
	Poll *Poll `json:"poll"`
	// Service message: the supergroup has been created.
	// This field can't be received in a message coming through updates, because bot can't be a member of a supergroup when it is created.
	// It can only be found in reply_to_message if someone replies to a very first message in a directly created supergroup.
	//
	// Optional.
	SupergroupChatCreated bool `json:"supergroup_chat_created"`
	// The unique identifier of a media message group this message belongs to.
	//
	// Optional.
	MediaGroupId string `json:"media_group_id"`
	// Message is a service message about a successful payment, information about the payment.
	// More about payments ».
	//
	// Optional.
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment"`
	// Service message.
	// A user in the chat triggered another user's proximity alert while sharing Live Location.
	//
	// Optional.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered"`
	// Sender, empty for messages sent to channels.
	//
	// Optional.
	From *User `json:"from"`
	// Message is a shared location, information about the location.
	//
	// Optional.
	Location *Location `json:"location"`
	// The group has been migrated to a supergroup with the specified identifier.
	// This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	//
	// Optional.
	MigrateToChatId int `json:"migrate_to_chat_id"`
}
