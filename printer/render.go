package printer

// Gracias a BigJk por este código.
// https://github.com/BigJk/snd/blob/master/rendering/rendering.go
import (
	"bytes"
	"errors"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
	"image"
	"net/url"
)

var browser *rod.Browser

func init() {
	browser = rod.New().MustConnect()
}

func screenshotPage(page *rod.Page, width int) (image.Image, error) {
	imageData, err := page.MustSetViewport(width, 10000, 1.0, false).MustElement("body").Screenshot(proto.PageCaptureScreenshotFormatPng, 100)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewBuffer(imageData))
	if err != nil {
		return nil, err
	}

	if img.Bounds().Max.Y >= 9500 {
		return nil, errors.New("too large")
	}

	return img, nil
}

func RenderHTML(html string, width int) (image.Image, error) {
	page := browser.MustPage("data:text/html," + url.PathEscape(html))
	page.MustWaitLoad().MustWaitIdle()
	defer page.Close()

	return screenshotPage(page, width)
}
