package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	token        = "THIS_SHOULD_BE_SECRET_TOKEN"
	databasePath = "/tmp/db"
	configPath   = "/tmp"
	port         = -1
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Here be bots and beer!")
}

func main() {
	//Flags
	log.Println("Parsing flags")
	// Try to get port from env, if fails we set envPort to -1 and hope we get port from flags
	envPort, err := strconv.Atoi(os.Getenv("BOT_PORT"))
	if err != nil {
		envPort = -1
	}
	//flag.StringVar(&configFile, "config", "config.yaml", "Config file to read settings from")
	flag.IntVar(&port, "port", envPort, "Port to listen to, default is $BOT_PORT")
	flag.StringVar(&token, "token", os.Getenv("BOT_TOKEN"), "Token provided by BotFather, default is $BOT_TOKEN")
	flag.Parse()

	if port <= 0 {
		// We are in an unrecoverable state, lets panic
		panic("No port specified")
	}
	log.Println("Setting up webpage")
	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+strconv.Itoa(port), nil)
	log.Println("We have webpage")

	log.Println("Creating bot")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
