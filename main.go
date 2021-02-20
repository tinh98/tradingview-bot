package main

import (
	"fmt"
	"github.com/cmingou/tradingview-bot/internal/tradingview"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
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
		fmt.Printf("Got Text:\n%+v\n", m)
	})

	b.Handle("/chart1d", func(m *tb.Message) {
		tradingview.SearchAndSendStockImage(b, m, tradingview.Time1D)
	})

	b.Handle("/chart1m", func(m *tb.Message) {
		tradingview.SearchAndSendStockImage(b, m, tradingview.Time1M)
	})

	b.Handle("/chart3m", func(m *tb.Message) {
		tradingview.SearchAndSendStockImage(b, m, tradingview.Time3M)
	})

	b.Handle("/chart1y", func(m *tb.Message) {
		tradingview.SearchAndSendStockImage(b, m, tradingview.Time1Y)
	})

	b.Start()
}
