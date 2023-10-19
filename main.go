package main

import (
	"log"
)

func main() {
	// Telegram bot tokeniniz
	botToken := "BURAYA_BOT_TOKENINIZI_GIRIN"

	// Telegram botunu oluşturun
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	// Bot bilgilerini görüntüle
	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Gönderilecek mesajı oluşturun
	message := tgbotapi.NewMessage(TARGET_CHAT_ID, "Merhaba, bu bir test mesajıdır!")

	// Mesajı gönderin
	_, err = bot.Send(message)
	if err != nil {
		log.Panic(err)
	}
}
