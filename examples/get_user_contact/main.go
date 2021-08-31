package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/Meat-Hook/tg"
	"github.com/Meat-Hook/tg/domain"
)

var TOKEN = "SET_YOUR_TOKEN"

func main() {
	err := run(context.Background(), TOKEN)
	if err != nil {
		log.Fatalf("run: %s\n", err)
	}
}

func run(ctx context.Context, token string) error {
	bot, err := tg.NewBot(ctx, token)
	if err != nil {
		return fmt.Errorf("tg.NewBot: %w", err)
	}

	updates, err := bot.Updates(ctx, tg.UpdateRequest{})
	if err != nil {
		return fmt.Errorf("bot.Updates: %w", err)
	}

	uniqueChatID := make(map[int]bool)
	for _, update := range updates {
		if uniqueChatID[update.Message.Chat.Id] {
			continue
		}

		uniqueChatID[update.Message.Chat.Id] = true
	}

	for id := range uniqueChatID {
		_, err := bot.SendMessage(ctx, tg.SendMessageRequest{
			ChatID: strconv.Itoa(id),
			Text:   "Hi, bro!",
			ReplyMarkup: domain.ReplyKeyboardMarkup{
				Keyboard: [][]domain.KeyboardButton{
					{
						{
							Text:            "Can you send me your phone?",
							RequestContact:  true,
							RequestLocation: false,
							RequestPoll:     nil,
						},
					},
				},
				ResizeKeyboard:        true,
				OneTimeKeyboard:       true,
				InputFieldPlaceholder: "",
				Selective:             true,
			},
		})
		if err != nil {
			return fmt.Errorf("bot.SendMessage: %w", err)
		}

	}

	return nil
}
