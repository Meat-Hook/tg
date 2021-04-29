package tg

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	apiURL = "https://api.telegram.org"
)

// endpoints
const (
	selfData = "getMe"
)

// Bot allows you to interact with the Telegram Bot API.
type Bot struct {
	token  string
	client *http.Client
	info   User
}

// Info returns Bot information.
func (b *Bot) Info() User {
	return b.info
}

// NewBot returns new Bot instance.
func NewBot(ctx context.Context, token string) (*Bot, error) {
	b := &Bot{
		token:  token,
		client: &http.Client{}, // TODO: Add custom client.
		info:   User{},
	}

	err := b.init(ctx)
	if err != nil {
		return nil, fmt.Errorf("init bot: %w", err)
	}

	return b, nil
}

func (b *Bot) init(ctx context.Context) error {
	u := User{}
	err := makeRequest(ctx, b.token, selfData, b.client, nil, &u)
	if err != nil {
		return fmt.Errorf("make request: %w", err)
	}

	b.info = u

	return nil
}

func makeRequest(ctx context.Context, token, method string, client *http.Client, body interface{}, unmarshal interface{}) error {
	endpoint, err := url.Parse(fmt.Sprintf("%s/%s/%s", apiURL, fmt.Sprintf("bot%s", token), method))
	if err != nil {
		return fmt.Errorf("build url: %w", err)
	}

	buf, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("json marhal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint.String(), bytes.NewBuffer(buf))
	if err != nil {
		return fmt.Errorf("build req: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("client do: %w", err)
	}
	defer res.Body.Close() // TODO: Add error handler.

	r := &response{}
	err = json.NewDecoder(res.Body).Decode(r)
	if err != nil {
		return fmt.Errorf("json decode: %w", err)
	}

	if !r.Ok {
		return fmt.Errorf("%s", r.Description)
	}

	return json.Unmarshal(r.Result, unmarshal)
}
