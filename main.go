package main // import "github.com/hhtpcd/telegram-send"

import (
	"flag"
	"log"
	"strconv"

	"os"

	"gopkg.in/telegram-bot-api.v4"
)

var (
	messageBody, botToken string
	recvID                int64
	apiKey                = "TELEGRAM_API_TOKEN"
	recipKey              = "TELEGRAM_API_RECIPIENT"
)

func init() {
	flag.StringVar(&messageBody, "m", "", "Message to send")

	flag.Parse()

	if os.Getenv(apiKey) == "" {
		log.Fatalln("Cannot find API credentials")
	}

	if os.Getenv(recipKey) == "" {
		log.Fatalln("Cannot find recipient")
	}

	if messageBody == "" {
		log.Fatalln("Please provide message")
	}
}

func main() {
	botToken := os.Getenv(apiKey)
	recv := os.Getenv(recipKey)

	recvID, err := strconv.ParseInt(recv, 10, 64)
	if err != nil {
		log.Fatalln(err)
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Fatal(err)
	}

	msg := tgbotapi.NewMessage(recvID, messageBody)
	_, err = bot.Send(msg)
	if err != nil {
		log.Fatal(err)
	}

}
