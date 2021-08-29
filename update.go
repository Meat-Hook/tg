package tg

import (
	"context"
	"fmt"

	"github.com/Meat-Hook/tg/domain"
)

// UpdateRequest is args for method Bot.Updates.
type UpdateRequest struct {
	// Identifier of the first update to be returned.
	// Must be greater by one than the highest among the identifiers of previously received updates.
	// By default, updates starting with the earliest unconfirmed update are returned.
	// An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id.
	// The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates queue.
	// All previous updates will forgotten.
	//
	// Optional.
	Offset int `json:"offset,omitempty"`
	// Limits the number of updates to be retrieved. Values between 1-100 are accepted.
	// Defaults to 100.
	//
	// Optional.
	Limit int `json:"limit,omitempty"`
	// Timeout in seconds for long polling.
	// Defaults to 0, i.e. usual short polling. Should be positive, short polling should be used for testing purposes only.
	//
	// Optional.
	Timeout int `json:"timeout,omitempty"`
	// A JSON-serialized list of the update types you want your bot to receive.
	// For example, specify [“message”, “edited_channel_post”, “callback_query”] to only receive updates of these types.
	// See Update for a complete list of available update types.
	// Specify an empty list to receive all update types except chat_member (default).
	// If not specified, the previous setting will be used.
	//
	// Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted updates may be received for a short period of time.
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// Updates get list updates.
// Not work if set web-hook.
// There are two mutually exclusive ways of receiving updates
// for your bot — the getUpdates method on one hand and Webhooks on the other.
// Incoming updates are stored on the server until the bot receives them either way,
// but they will not be kept longer than 24 hours.
func (b *Bot) Updates(ctx context.Context, req UpdateRequest) ([]domain.Update, error) {
	var res []domain.Update
	err := makeRequest(ctx, b.client, getUpdates, b.token, req, &res)
	if err != nil {
		return nil, fmt.Errorf("makeRequest: %w", err)
	}

	return res, nil
}
