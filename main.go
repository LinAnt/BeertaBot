package main

import (
	"flag"
	"log"
	"net/http"
	"strconv"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	token        = "THIS_SHOULD_BE_SECRET_TOKEN"
	databasePath = "/tmp/db"
	configPath   = "/tmp"
	port         = 8443
)

func main() {
	//Flags
	var configFile string
	flag.StringVar(&configFile, "config", "config.yaml", "Config file to read settings from")
	flag.IntVar(&port, "port", 8080, "Port to listen to (default: 8443")
	flag.Parse()

	// Read settings from config file
	//var conf configuration.Config

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	_, err = bot.SetWebhook(tgbotapi.NewWebhookWithCert("https://www.google.com:8443/"+bot.Token, "cert.pem"))
	if err != nil {
		log.Fatal(err)
	}

	updates := bot.ListenForWebhook("/" + bot.Token)

	go http.ListenAndServeTLS("0.0.0.0:"+strconv.Itoa(port), "cert.pem", "key.pem", nil)

	for update := range updates { // update channel
		log.Printf("%+v\n", update)
	}
}
