package main

import (
	"fmt"
	"github.com/cmingou/tradingview-bot/internal/tradingview"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"regexp"
	"strings"
	"time"
)

var (
	b *tb.Bot
)

func main() {
	var err error
	b, err = tb.NewBot(tb.Settings{
		// Token for bot
		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tb.OnText, func(m *tb.Message) {
		if strings.HasPrefix(m.Text, "$") && strings.Count(m.Text, "$") == 1 {
			symbol := m.Text[1:]

			match, err := regexp.MatchString("^[A-Za-z]+$", symbol)
			if err != nil {
				fmt.Printf("%v\n", err)
			}

			if match {
				tradingview.SearchAndSendStockImage(b, m, symbol, tradingview.Time1D, false)
			}
		}
	})

	b.Handle("/chart1d", func(m *tb.Message) {
		tradingview.SearchAndSendStockImage(b, m, m.Payload, tradingview.Time1D, true)
	})

	b.Handle("/chart1m", func(m *tb.Message) {
		tradingview.SearchAndSendStockImage(b, m, m.Payload, tradingview.Time1M, true)
	})

	b.Handle("/chart3m", func(m *tb.Message) {
		tradingview.SearchAndSendStockImage(b, m, m.Payload, tradingview.Time3M, true)
	})

	b.Handle("/chart1y", func(m *tb.Message) {
		tradingview.SearchAndSendStockImage(b, m, m.Payload, tradingview.Time1Y, true)
	})

	b.Start()
}
