package domain

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Errors.
var (
	// ErrUnknownChatType unknown ChatType value.
	ErrUnknownChatType = errors.New("unknown ChatType")
)

var (
	chatType                  = ChatType(0)
	_        json.Marshaler   = chatType
	_        json.Unmarshaler = &chatType
)

// ChatType enum.
type ChatType uint8

const (
	_          ChatType = iota
	Sender              // sender
	Private             // private
	Group               // group
	Supergroup          // supergroup
	Channel             // channel
)

// UnmarshalJSON implements json.Unmarshaler.
func (i *ChatType) UnmarshalJSON(bytes []byte) error {
	str := string(bytes)

	var res ChatType
	switch str {
	case Sender.String():
		res = Sender
	case Private.String():
		res = Private
	case Group.String():
		res = Group
	case Supergroup.String():
		res = Supergroup
	case Channel.String():
		res = Channel
	default:
		return fmt.Errorf("%w: %s", ErrUnknownChatType, str)
	}

	*i = res

	return nil
}

// MarshalJSON implements json.Marshaler.
func (i ChatType) MarshalJSON() ([]byte, error) {
	return []byte(i.String()), nil
}

// InlineQuery this object represents an incoming inline query.
// When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	// ID unique identifier for this query.
	ID string `json:"id"`
	// From sender.
	From User `json:"from"`
	// Query text of the query (up to 256 characters).
	Query string `json:"query"`
	// Offset of the results to be returned, can be controlled by the bot.
	Offset string `json:"offset"`
	// ChatType type of the chat, from which the inline query was sent.
	// The chat type should be always known for requests sent from official clients and most third-party clients,
	// unless the request was sent from a secret chat
	ChatType string `json:"chat_type"`
	// Location sender location, only for bots that request user location.
	Location Location `json:"location"`
}
