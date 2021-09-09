package main

import (
	"log"
	"strconv"
	"time"

	tb "gopkg.in/tucnak/telebot.v2"
)

var (
	botID string
)

func main() {
	b, err := tb.NewBot(tb.Settings{
		Token:  botID,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})

	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		log.Print("Request from " + m.Sender.Username + " URL : " + m.Text)
		b.Send(m.Sender, "Scrapping : "+m.Text+"\n")

		result, err := scrap(m.Text)

		if err != nil {
			b.Send(m.Sender, "Scrapping Error : "+err.Error())
		}

		for i, text := range result {
			content := "Part " + strconv.Itoa(i+1) + " of " + strconv.Itoa(len(result)) + "\n"
			content = content + text
			b.Send(m.Sender, content)
		}
	})

	b.Start()
}
