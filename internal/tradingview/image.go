package tradingview

import (
	"fmt"
	"github.com/cmingou/tradingview-bot/internal/screenshot"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
	"path"
	"runtime"
	"strconv"
	"time"
)

func tradingviewWidget(symbols, desc, timeRange string) string {
	return fmt.Sprintf(`
<!-- TradingView Widget BEGIN -->
<div class="tradingview-widget-container">
  <div id="tradingview_9efce"></div>
  <div class="tradingview-widget-copyright">TinhLee TradingView</div>
  <script type="text/javascript" src="https://s3.tradingview.com/tv.js"></script>
  <script type="text/javascript">
  new TradingView.MediumWidget(
  {
  "symbols": [
    [
      "%s",
      "%s%s"
    ]
  ],
  "chartOnly": false,
  "width": 1000,
  "height": 500,
  "locale": "vi",
  "colorTheme": "dark",
  "gridLineColor": "#2A2E39",
  "trendLineColor": "#1976D2",
  "fontColor": "#787B86",
  "underLineColor": "rgba(55, 166, 239, 0.15)",
  "isTransparent": false,
  "autosize": false,
  "container_id": "tradingview_9efce"
}
  );
  </script>
</div>
<!-- TradingView Widget END -->
`, desc, symbols, timeRange)
}

func tradingviewDetailWidget(symbols, timeRange string) string {
	return fmt.Sprintf(`
<!-- TradingView Widget BEGIN -->
<div class="tradingview-widget-container">
  <div id="tradingview_f9dfa"></div>
  <div class="tradingview-widget-copyright">图表仅供讨论使用，资料由TradingView提供</div>
  <script type="text/javascript" src="https://s3.tradingview.com/tv.js"></script>
  <script type="text/javascript">
  new TradingView.widget(
  {
  "width": 1000,
  "height": 500,
  "symbol": "%v",
  "timezone": "America/New_York",
  "theme": "dark",
  "style": "1",
  "locale": "vi",
  "toolbar_bg": "#f1f3f6",
  "enable_publishing": false,
  "hide_top_toolbar": true,
  "range": "%v",
  "allow_symbol_change": true,
  "save_image": false,
  "studies": [
    "MASimple@tv-basicstudies"
  ],
  "container_id": "tradingview_f9dfa"
}
  );
  </script>
</div>
<!-- TradingView Widget END -->
`, symbols, timeRange)
}

type StockImageOptions struct {
	Symbol      string
	Description string
	Time        string

	// BinaryPath the path to your capture-website binary. REQUIRED
	//
	// Must be absolute path e.g /usr/local/bin/capture-website
	BinaryPath string

	Input             string
	Output            string
	Dir               string
	Html              string
	Format            string
	Width             int64
	Height            int64
	Delay             int
	Overwrite         bool
	Darkmode          bool
	TechnicalAnalysis bool
}

func (s *StockImageOptions) FilePath() string {
	return s.Dir + "/" + s.Output + "." + s.Format
}

func (s *StockImageOptions) FilePathIndex() string {
	return s.Dir + "/index.html"
}

func (s *StockImageOptions) FileName() string {
	return s.Output + "." + s.Format
}

func (s *StockImageOptions) CountdownToDel() {
	go func() {
		// countdown
		timer := time.NewTimer(time.Second * 20)
		<-timer.C
		// delete images
		if err := os.Remove(s.FilePath()); err != nil {
			fmt.Printf("%v\n", err)
		}
	}()
}

func (options *StockImageOptions) GenerateImage() error {

	htmlString := TradingviewDetailWidgetTinhLee(options.Width, options.Height, options.Symbol, options.Time[1:])
	err := os.WriteFile(options.FilePathIndex(), []byte(htmlString), 0666)
	if err != nil {
		return err
	}

	_, fileName, _, _ := runtime.Caller(0)
	baseFilePath := "file:///" + path.Join(fileName, "../../../"+options.FilePathIndex())
	err = screenshot.InitScreenshot().Capture(options.FilePath(), &screenshot.ScreenshotParam{
		Url:     baseFilePath,
		Wait:    5,
		Quality: 100,
		Width:   options.Width,
		Height:  options.Height,
	})
	if err != nil {
		return err
	}
	return nil
}

func SearchAndSendStockImage(b *tb.Bot, m *tb.Message, symbol, timeRange string, delFile, technicalAnalysis bool) {
	var err error
	if symbol == "" {
		// Did not add stock id
		mReply, _ := b.Reply(m, "/chart AAPL")

		// Waiting to delete
		go func() {
			time.Sleep(time.Second * 6)
			err = b.Delete(mReply)
			if err != nil {
				fmt.Printf("%v\n", err)
			}

			err = b.Delete(m)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}()
		return
	}
	tempDir := "temp"
	os.Mkdir(tempDir, 0777)

	imgName := strconv.Itoa(int(m.Unixtime)) + "-" + strconv.Itoa(m.ID) + "-" + symbol + "-" + m.Sender.Username
	const (
		width  = 1000
		height = 500
	)
	//Generate stock image
	s := StockImageOptions{
		Symbol:            symbol,
		Description:       symbol,
		Time:              timeRange,
		Input:             "-",
		Output:            imgName,
		Dir:               tempDir,
		Width:             width,
		Height:            height,
		Delay:             4,
		Overwrite:         true,
		Darkmode:          true,
		TechnicalAnalysis: technicalAnalysis,
		Format:            "png",
	}

	if err = s.GenerateImage(); err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	p := &tb.Photo{
		File: tb.File{
			FileLocal: s.FilePath(),
		},
		Width:  int(s.Width),
		Height: int(s.Height),
	}

	if _, err := b.Reply(m, p); err != nil {
		fmt.Printf("%v\n", err)
	}

	s.CountdownToDel()

	time.Sleep(time.Second)

	// Delete request user command
	if delFile {
		if err = b.Delete(m); err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}

func TradingviewDetailWidgetTinhLee(width, height int64, symbols, timeRange string) string {
	return fmt.Sprintf(`
<!-- TradingView Widget BEGIN -->
<div class="tradingview-widget-container">
  <div id="tradingview_f9dfa"></div>
  <div class="tradingview-widget-copyright">TradingViewTinhLee</div>
  <script type="text/javascript" src="https://s3.tradingview.com/tv.js"></script>
  <script type="text/javascript">
  new TradingView.widget(
  {
  "width": %d,
  "height": %d,
  "symbol": "%v",
  "timezone": "America/New_York",
  "theme": "dark",
  "style": "1",
  "locale": "vi",
  "toolbar_bg": "#f1f3f6",
  "enable_publishing": false,
  "hide_top_toolbar": true,
  "range": "%v",
  "allow_symbol_change": true,
  "save_image": false,
  "studies": [
    "MASimple@tv-basicstudies"
  ],
  "container_id": "tradingview_f9dfa"
}
  );
  </script>
</div>
<!-- TradingView Widget END -->
`, width, height, symbols, timeRange)
}
