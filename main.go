package main

import (
	"log"

	"github.com/Andrem19/telegramGPT/helpers"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	config, err := helpers.LoadConfig(".")
	helpers.StartWithDb(config)
	bot, err := tgbotapi.NewBotAPI(config.BOT_API_TOKEN)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			go func() {
				answer, err := helpers.Switcher(update.Message.Text, update.Message.Chat.ID)
				if err != nil {
					helpers.AddToLog(err.Error())
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
				msg.ReplyToMessageID = update.Message.MessageID
				bot.Send(msg)
			}()
		}
	}
}