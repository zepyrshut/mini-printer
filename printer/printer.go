package printer

import (
	"bytes"
	"image"
)

type PrintData struct {
	Image *image.Image
	Data  []byte
}

func ProcessPrint(result string) PrintData {
	imageHtml, _ := RenderHTML(result, 576)

	bufB := &bytes.Buffer{}
	RasterImage(bufB, imageHtml)
	LineFeed(bufB)
	FullCut(bufB)

	return PrintData{Image: &imageHtml, Data: bufB.Bytes()}
}
