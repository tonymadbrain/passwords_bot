package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sethvargo/go-password/password"
)

var helpMessage = "Hello! Touch or send 1 for simple password, 2 for middle password and 3 for strong."
var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
)

func newStart(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	log.Printf("new start")
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpMessage)
	msg.ReplyMarkup = numericKeyboard
	bot.Send(msg)
}

func processUpdate(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

	switch update.Message.Text {
	case "1":
		res, err := password.Generate(8, 0, 0, false, false)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("generated 1")
		msg.Text = res
	case "2":
		res, err := password.Generate(12, 4, 0, false, false)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("generated 2")
		msg.Text = res
	case "3":
		res, err := password.Generate(16, 4, 4, false, false)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("generated 3")
		msg.Text = res
	case "4":
		res, err := password.Generate(32, 0, 0, false, false)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("generated 4")
		msg.Text = res
	case "5":
		res, err := password.Generate(32, 8, 0, false, false)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("generated 5")
		msg.Text = res
	case "6":
		res, err := password.Generate(32, 8, 8, false, false)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("generated 6")
		msg.Text = res
	case "open":
		msg.Text = "keyboard open"
		msg.ReplyMarkup = numericKeyboard
	case "close":
		msg.Text = "keyboard close"
		msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
	default:
		msg.Text = helpMessage
	}

	bot.Send(msg)
}

func main() {
	token := os.Getenv("TG_BOT_TOKEN")
	if token == "" {
		log.Print("no TG_BOT_TOKEN set")
		os.Exit(1)
	}

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	debug := false
	if os.Getenv("TG_BOT_DEBUG") == "true" {
		debug = true
	}
	bot.Debug = debug

	log.Printf("authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text == "/start" {
			go newStart(bot, update)
			continue
		}

		go processUpdate(bot, update)
	}
}
