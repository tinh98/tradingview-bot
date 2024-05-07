package tradingview

import (
	"fmt"
	"path"
	"runtime"
	"testing"
)

func TestName(t *testing.T) {
	//html := tradingviewDetailWidget("BNBUSDT", Time1D[1:])
	//err := os.WriteFile("../screenshot/index.html", []byte(html), 0)
	//if err != nil {
	//	panic(err)
	//}
	//err = screenshot.InitScreenshot().Capture(&screenshot.ScreenshotParam{
	//	Url:      "file:////Users/tinhle/tradingview-bot/internal/screenshot/index.html",
	//	Filename: "tinh",
	//	Wait:     5,
	//	Quality:  100,
	//	Width:    1000,
	//	Height:   500,
	//})
	//fmt.Println(err)
	_, fileName, _, _ := runtime.Caller(0)
	baseFilePath := path.Join("file:///", fileName, "../")
	fmt.Println(baseFilePath)
}
