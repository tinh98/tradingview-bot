package tradingview

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"regexp"
	"strings"
	"time"
)

var (
	b *tb.Bot
)

func InitBot(token string) {
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
				SearchAndSendStockImage(b, m, symbol, Time1D, false, true)
			}
		}
	})

	//chart1d - 查询股票图表(1天)
	//chart1m - 查询股票图表(1个月)
	//chart3m - 查询股票图表(3个月)
	//chart1y - 查询股票图表(1年)
	//chart - 查询股票图表(all)

	b.Handle("/chart1d", func(m *tb.Message) {
		SearchAndSendStockImage(b, m, m.Payload, Time1D, false, true)
	})

	b.Handle("/chart1m", func(m *tb.Message) {
		SearchAndSendStockImage(b, m, m.Payload, Time1M, false, true)
	})

	b.Handle("/chart3m", func(m *tb.Message) {
		SearchAndSendStockImage(b, m, m.Payload, Time3M, false, true)
	})

	b.Handle("/chart1y", func(m *tb.Message) {
		SearchAndSendStockImage(b, m, m.Payload, Time1Y, false, true)
	})

	b.Handle("/chart", func(m *tb.Message) {
		SearchAndSendStockImage(b, m, m.Payload, TimeAll, false, true)
	})

	fmt.Printf("bot started!!\n")
	b.Start()
}
