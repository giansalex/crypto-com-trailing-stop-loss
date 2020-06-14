package stoploss

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// Notify notify stoploss
type Notify struct {
	tlgToken  string
	channelID int64
}

// NewNotify create Notify instance
func NewNotify(telegramToken string, channelID int64) *Notify {
	return &Notify{telegramToken, channelID}
}

// Send send message
func (notify *Notify) Send(message string) {
	fmt.Println(message)

	if notify.tlgToken == "" {
		return
	}

	bot, err := tgbotapi.NewBotAPI(notify.tlgToken)
	if err != nil {
		fmt.Println("Cannot connect to Telegram:", err.Error())

		return
	}

	msg := tgbotapi.NewMessage(notify.channelID, message)

	bot.Send(msg)
}
