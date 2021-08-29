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
	client := &http.Client{}
	botInfo := domain.User{}

	err := makeRequest(ctx, client, getMe, token, nil, &botInfo)
	if err != nil {
		return nil, fmt.Errorf("makeRequest: %w", err)
	}

	b := &Bot{
		token:  token,
		client: client, // TODO: Add custom client.
		info:   botInfo,
	}

	return b, nil
}

func makeRequest(ctx context.Context, client *http.Client, method, token string, body interface{}, unmarshal interface{}) (err error) {
	endpoint, err := url.Parse(fmt.Sprintf("%s/%s/%s", apiURL, fmt.Sprintf("bot%s", token), method))
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

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("b.client.Do: %w", err)
	}
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
		return fmt.Errorf("r.Description: %s", r.Description)
	}

	return json.Unmarshal(r.Result, unmarshal)
}
