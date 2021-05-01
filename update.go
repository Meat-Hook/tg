package tg

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/valyala/fastjson"
)

var (
	_ json.Unmarshaler = &Update{}
)

// Errors.
var (
	// UnknownField about unknown field from update message.
	UnknownField = errors.New("unknown field")
)

const (
	_ UpdateKind = iota
	// StdMessage new incoming message of any kind — text, photo, sticker, etc.
	StdMessage
	// EditedMessage mew version of a message that is known to the bot and was edited.
	EditedMessage
	// ChannelPost new incoming channel post of any kind — text, photo, sticker, etc.
	ChannelPost
	// EditedChannelPost new incoming inline query.
	EditedChannelPost
	// InlineQuery new incoming inline query.
	InlineQuery
	// ChosenInlineResult the result of an inline query that was chosen by a user and sent to their chat partner.
	// Please see our documentation on the feedback collecting for details on how to enable these updates for your bot.
	ChosenInlineResult
	// CallbackQuery new incoming callback query.
	CallbackQuery
	// ShippingQuery new incoming shipping query. Only for invoices with flexible price.
	ShippingQuery
	// PreCheckoutQuery new incoming shipping query. Only for invoices with flexible price.
	PreCheckoutQuery
	// Poll new poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot.
	Poll
	// PollAnswer a user changed their answer in a non-anonymous poll.
	// Bots receive new votes only in polls that were sent by the bot itself.
	PollAnswer
	// MyChatMember the bot's chat member status was updated in a chat.
	// For private chats, this update is received only when the bot is blocked or unblocked by the user.
	MyChatMember
	// ChatMember a chat member's status was updated in a chat.
	// The bot must be an administrator in the chat and must explicitly specify “chat_member”
	// in the list of allowed_updates to receive these updates.
	ChatMember
)

type (
	// UpdateOption for build update config.
	UpdateOption func(cfg *updateConfig)
	// contains all args for getting update.
	updateConfig struct {
		Offset  int `json:"offset,omitempty"`
		Limit   int `json:"limit,omitempty"`
		Timeout int `json:"timeout,omitempty"`
	}
	// UpdateKind contains information about action kind.
	UpdateKind uint8
	// Update is an update response, from GetUpdates.
	// Has only one of field.
	Update struct {
		// UpdateID is the update's unique identifier.
		// Update identifiers start from a certain positive number and increase sequentially.
		// This ID becomes especially handy if you're using Webhooks,
		// since it allows you to ignore repeated updates or to restore
		// the correct update sequence, should they get out of order.
		// If there are no new updates for at least a week, then identifier
		// of the next update will be chosen randomly instead of sequentially.
		UpdateID int
		// Kind contains information about update kind.
		Kind UpdateKind
		// Body contains body message.
		Body json.RawMessage
	}
)

// UpdateWithOffset set offset for update config.
func UpdateWithOffset(offset int) UpdateOption {
	return func(cfg *updateConfig) {
		cfg.Offset = offset
	}
}

// UpdateWithLimit set offset for update config.
func UpdateWithLimit(limit int) UpdateOption {
	return func(cfg *updateConfig) {
		cfg.Limit = limit
	}
}

// UpdateWithTimeout set offset for update config.
func UpdateWithTimeout(timeout int) UpdateOption {
	return func(cfg *updateConfig) {
		cfg.Timeout = timeout
	}
}

// For unmarshal json.
var (
	id     = `update_id`
	fields = map[string]UpdateKind{
		"message":              StdMessage,
		"edited_message":       EditedMessage,
		"channel_post":         ChannelPost,
		"edited_channel_post":  EditedChannelPost,
		"inline_query":         InlineQuery,
		"chosen_inline_result": ChosenInlineResult,
		"callback_query":       CallbackQuery,
		"shipping_query":       ShippingQuery,
		"pre_checkout_query":   PreCheckoutQuery,
		"poll":                 Poll,
		"poll_answer":          PollAnswer,
		"my_chat_member":       MyChatMember,
		"chat_member":          ChatMember,
	}
)

// UnmarshalJSON implements json.Unmarshaler.
func (u *Update) UnmarshalJSON(b []byte) error {
	v, err := fastjson.Parse(string(b))
	if err != nil {
		return fmt.Errorf("fastjson parse: %w", err)
	}

	u.UpdateID = v.GetInt(id)
	for name, kind := range fields {
		value := v.GetObject(name)
		if value == nil {
			continue
		}

		u.Body = value.MarshalTo(nil)
		u.Kind = kind
		break
	}

	if u.Body == nil {
		return UnknownField
	}

	return nil
}

// Updates get list updates.
// Not work if set web-hook.
// There are two mutually exclusive ways of receiving updates
// for your bot — the getUpdates method on one hand and Webhooks on the other.
// Incoming updates are stored on the server until the bot receives them either way,
// but they will not be kept longer than 24 hours.
func (b *Bot) Updates(ctx context.Context, opts ...UpdateOption) ([]Update, error) {
	cfg := updateConfig{}
	for i := range opts {
		opts[i](&cfg)
	}

	var res []Update
	err := b.makeRequest(ctx, getUpdates, cfg, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
