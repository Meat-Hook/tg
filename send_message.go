package tg

import (
	"context"
	"fmt"

	"github.com/Meat-Hook/tg/domain"
)

// SendMessageRequest is args for method Bot.SendMessage.
type SendMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID string `json:"chat_id"`
	// Text of the message to be sent, 1-4096 characters after entities parsing.
	Text string `json:"text"`
	// Mode for parsing entities in the message text. See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode,omitempty"`
	// List of special entities that appear in message text, which can be specified instead of parse_mode.
	//
	// Optional.
	Entities []domain.MessageEntity `json:"entities,omitempty"`
	// Disables link previews for links in this message.
	//
	// Optional.
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	// Sends the message silently. Users will receive a notification with no sound.
	//
	// Optional.
	DisableNotification bool `json:"disable_notification,omitempty"`
	// If the message is a reply, ID of the original message.
	//
	// Optional.
	ReplyToMessageID int `json:"reply_to_message_id,omitempty"`
	// Pass True, if the message should be sent even if the specified replied-to message is not found.
	//
	// Optional.
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Additional interface options.
	// A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard or to force a reply from the user.
	//
	// Optional.
	ReplyMarkup domain.ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendMessage use this method to send text messages.
// Use this method to send text messages. On success, the sent Message is returned.
func (b *Bot) SendMessage(ctx context.Context, req SendMessageRequest) (*domain.Message, error) {
	res := &domain.Message{}
	err := makeRequest(ctx, b.client, sendMessage, b.token, req, &res)
	if err != nil {
		return nil, fmt.Errorf("makeRequest: %w", err)
	}

	return res, nil
}
