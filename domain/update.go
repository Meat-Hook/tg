package domain

// Update This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	// New incoming inline query.
	//
	// Optional.
	InlineQuery *InlineQuery `json:"inline_query"`
	// New incoming callback query.
	//
	// Optional.
	CallbackQuery *CallbackQuery `json:"callback_query"`
	// New incoming pre-checkout query.
	// Contains full information about checkout.
	//
	// Optional.
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query"`
	// New poll state.
	// Bots receive only updates about stopped polls and polls, which are sent by the bot.
	//
	// Optional.
	Poll *Poll `json:"poll"`
	// A user changed their answer in a non-anonymous poll.
	// Bots receive new votes only in polls that were sent by the bot itself.
	//
	// Optional.
	PollAnswer *PollAnswer `json:"poll_answer"`
	// A chat member's status was updated in a chat.
	// The bot must be an administrator in the chat and must explicitly specify “chat_member” in the list of allowed_updates to receive these updates.
	//
	// Optional.
	ChatMember *ChatMemberUpdated `json:"chat_member"`
	// New incoming message of any kind — text, photo, sticker, etc.
	//
	// Optional.
	Message *Message `json:"message"`
	// New version of a message that is known to the bot and was edited.
	//
	// Optional.
	EditedMessage *Message `json:"edited_message"`
	// New incoming channel post of any kind — text, photo, sticker, etc.
	//
	// Optional.
	ChannelPost *Message `json:"channel_post"`
	// The result of an inline query that was chosen by a user and sent to their chat partner.
	// Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	//
	// Optional.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result"`
	// The update's unique identifier.
	// Update identifiers start from a certain positive number and increase sequentially.
	// This ID becomes especially handy if you're using Webhooks, since it allows you to ignore repeated updates or to restore the correct update sequence, should they get out of order.
	// If there are no new updates for at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	UpdateId int `json:"update_id"`
	// New version of a channel post that is known to the bot and was edited.
	//
	// Optional.
	EditedChannelPost *Message `json:"edited_channel_post"`
	// New incoming shipping query.
	// Only for invoices with flexible price.
	//
	// Optional.
	ShippingQuery *ShippingQuery `json:"shipping_query"`
	// The bot's chat member status was updated in a chat.
	// For private chats, this update is received only when the bot is blocked or unblocked by the user.
	//
	// Optional.
	MyChatMember *ChatMemberUpdated `json:"my_chat_member"`
}
