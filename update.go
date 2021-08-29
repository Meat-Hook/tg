package tg

import (
	"context"

	"github.com/Meat-Hook/tg/domain"
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

// Updates get list updates.
// Not work if set web-hook.
// There are two mutually exclusive ways of receiving updates
// for your bot — the getUpdates method on one hand and Webhooks on the other.
// Incoming updates are stored on the server until the bot receives them either way,
// but they will not be kept longer than 24 hours.
func (b *Bot) Updates(ctx context.Context, opts ...UpdateOption) ([]domain.Update, error) {
	cfg := updateConfig{}
	for i := range opts {
		opts[i](&cfg)
	}

	var res []domain.Update
	err := b.makeRequest(ctx, getUpdates, cfg, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
