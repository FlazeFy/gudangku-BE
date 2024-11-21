package messager

import (
	"fmt"
	"gudangku/configs"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendTelegramMessage(chatID int64, message string) error {
	botToken := configs.GetConfigTele().TELE_TOKEN
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return fmt.Errorf("failed to create Telegram bot: %w", err)
	}

	msg := tgbotapi.NewMessage(chatID, message)

	_, err = bot.Send(msg)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	log.Printf("Message sent to chat ID: %d", chatID)
	return nil
}
