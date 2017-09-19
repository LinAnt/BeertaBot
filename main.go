package main

import (
	"log"
	"net/http"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	token = "417296719:AAFZUm-0aw6vrNpQ5Slps9XRLhsitvO-Juc"
)

func main() {
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
	go http.ListenAndServeTLS("0.0.0.0:8443", "cert.pem", "key.pem", nil)

	for update := range updates { // update channel
		log.Printf("%+v\n", update)
	}
}
