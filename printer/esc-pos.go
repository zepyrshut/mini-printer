package printer

// Gracias a BigJk por este código.
// https://github.com/BigJk/snd/blob/master/thermalprinter/epson/epson.go
import (
	"image"
	"io"
	"math"
)

func FullCut(buf io.Writer) {
	_, _ = buf.Write([]byte{0x1B, 0x6D, '\n'})
}

func LineFeed(buf io.Writer) {
	_, _ = buf.Write([]byte{0x1B, 0x64, '\005'})
}

func RasterImage(buf io.Writer, image image.Image) {
	bb := image.Bounds()

	width := bb.Max.X
	if width%8 != 0 {
		width += 8
	}

	xL := uint8(width % 2048 / 8)
	xH := uint8(width / 2048)
	yL := uint8(bb.Max.Y % 256)
	yH := uint8(bb.Max.Y / 256)

	_, _ = buf.Write([]byte{0x1d, 0x76, 0x30, 48, xL, xH, yL, yH})

	var cb byte
	var bindex byte
	for y := 0; y < bb.Max.Y; y++ {
		for x := 0; x < int(math.Ceil(float64(bb.Max.X)/8.0))*8; x++ {
			if x < bb.Max.X {
				r, g, b, a := image.At(x, y).RGBA()
				r, g, b, a = r>>8, g>>8, b>>8, a>>8

				if a > 255/2 {
					grayscale := 0.2126*float64(r) + 0.7152*float64(g) + 0.0722*float64(b)

					if grayscale < 185 {
						mask := byte(1) << byte(7-bindex)
						cb |= mask
					}
				}
			}

			bindex++
			if bindex == 8 {
				_, _ = buf.Write([]byte{cb})
				bindex = 0
				cb = 0
			}
		}
	}
}
