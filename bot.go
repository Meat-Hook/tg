package tg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Meat-Hook/tg/domain"
)

const (
	apiURL = "https://api.telegram.org"
)

// endpoints
const (
	selfData   = "getMe"
	getUpdates = "getUpdates"
)

type (
	// Bot allows you to interact with the Telegram Bot API.
	Bot struct {
		token  string
		client *http.Client
		info   domain.User
	}
	// Contains data from the Telegram API with the result.
	response struct {
		Ok          bool                `json:"ok"`
		Result      json.RawMessage     `json:"result"`
		ErrorCode   int                 `json:"error_code"`
		Description string              `json:"description"`
		Parameters  *responseParameters `json:"parameters"`
	}
	responseParameters struct {
		MigrateToChatID int64 `json:"migrate_to_chat_id"` // optional
		RetryAfter      int   `json:"retry_after"`        // optional
	}
)

// Info returns Bot information.
func (b *Bot) Info() domain.User {
	return b.info
}

// NewBot returns new Bot instance.
func NewBot(ctx context.Context, token string) (*Bot, error) {
	b := &Bot{
		token:  token,
		client: &http.Client{}, // TODO: Add custom client.
		info:   domain.User{},
	}

	err := b.init(ctx)
	if err != nil {
		return nil, fmt.Errorf("init bot: %w", err)
	}

	return b, nil
}

func (b *Bot) init(ctx context.Context) error {
	u := domain.User{}
	err := b.makeRequest(ctx, selfData, nil, &u)
	if err != nil {
		return fmt.Errorf("b.makeRequest: %w", err)
	}

	b.info = u

	return nil
}

func (b *Bot) makeRequest(ctx context.Context, method string, body interface{}, unmarshal interface{}) (err error) {
	endpoint, err := url.Parse(fmt.Sprintf("%s/%s/%s", apiURL, fmt.Sprintf("bot%s", b.token), method))
	if err != nil {
		return fmt.Errorf("url.Parse: %w", err)
	}

	buf, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), bytes.NewBuffer(buf))
	if err != nil {
		return fmt.Errorf("http.NewRequestWithContext: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := b.client.Do(req)
	if err != nil {
		return fmt.Errorf("b.client.Do: %w", err)
	}
	defer res.Body.Close() // TODO: Add error handler.
	defer func() {
		closeErr := res.Body.Close()
		if closeErr != nil {
			err = fmt.Errorf("%w: %s", err, closeErr)
		}
	}()

	r := &response{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return fmt.Errorf("json.NewDecoder(res.Body).Decode: %w", err)
	}

	if !r.Ok {
		return fmt.Errorf("%s", r.Description)
	}

	return json.Unmarshal(r.Result, unmarshal)
}
