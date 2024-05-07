package screenshot

import (
	"context"
	"github.com/chromedp/chromedp"
	"os"
	"time"
)

type ScreenshotHandler struct {
}

func InitScreenshot() ScreenshotHandler {
	return ScreenshotHandler{}
}

func fullScreenshot(screenshotParam *ScreenshotParam, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.EmulateViewport(screenshotParam.Width, screenshotParam.Height),
		chromedp.Navigate(screenshotParam.Url),
		chromedp.Sleep(screenshotParam.Wait * time.Second),
		chromedp.FullScreenshot(res, screenshotParam.Quality),
	}
}

func writeFile(tempPath string, buf []byte) error {
	if err := os.WriteFile(tempPath, buf, 0o666); err != nil {
		return err
	}
	return nil
}

func CreateTempDir() (string, error) {
	tempDir, err := os.MkdirTemp("", "dir")
	if err != nil {
		return tempDir, err
	}
	return tempDir, nil
}

func RemoveTempDir(tempDir string) error {
	err := os.RemoveAll(tempDir)
	if err != nil {
		return err
	}
	return nil
}

func (h ScreenshotHandler) Capture(filePath string, screenshotParam *ScreenshotParam) (err error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	var buf []byte

	width := &screenshotParam.Width
	height := &screenshotParam.Height

	if *width == 0 {
		*width = 1490
	}

	if *height == 0 {
		*height = 1080
	}

	if err = chromedp.Run(ctx, fullScreenshot(screenshotParam, &buf)); err != nil {
		return err
	}

	if err = writeFile(filePath, buf); err != nil {
		return err
	}
	return nil
}
