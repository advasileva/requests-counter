package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"io"
	"net/http"
	"os"
)

func process(update tgbotapi.Update) (msg string) {
	id := update.Message.Chat.ID

	url := fmt.Sprintf("%s?id=%d", os.Getenv("URL"), id)

	_, err := http.Post(url, "", nil)
	if err != nil {
		panic(err)
	}

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}
		request := string(body)
		msg = fmt.Sprintf("It's your %v request", request)
	} else {
		msg = "Ops..."
	}
	return
}

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := bot.GetUpdatesChan(updateConfig)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		response := process(update)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		msg.ReplyToMessageID = update.Message.MessageID

		if _, err := bot.Send(msg); err != nil {
			panic(err)
		}
	}
}
