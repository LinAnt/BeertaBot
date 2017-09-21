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

func main() {
	//Flags
	fmt.Println("Parsing flags")
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

	fmt.Println("Creating bot")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully created bot")
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "/app/cert.pem"))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)

	go http.ListenAndServeTLS("0.0.0.0:"+strconv.Itoa(port), "cert.pem", "key.pem", nil)

	for update := range updates { // update channel
		log.Printf("%+v\n", update)
	}
}
