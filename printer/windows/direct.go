package windows

// Gracias a BigJk por este código.
// https://github.com/BigJk/snd/blob/master/printing/windows/direct.go
import (
	"fmt"
	"image"
	"math/rand"

	"github.com/alexbrainman/printer"
)

type Direct struct{}

func (dp *Direct) Name() string {
	return "Windows Direct Printing"
}

func (dp *Direct) Description() string {
	return "Directly print to a attached printer. Use the Name of the printer as Endpoint."
}

func (dp *Direct) AvailableEndpoints() (map[string]string, error) {
	names, err := printer.ReadNames()
	if err != nil {
		return nil, err
	}

	available := map[string]string{}
	for i := range names {
		available[names[i]] = names[i]
	}

	return available, nil
}

func (dp *Direct) Print(printerEndpoint string, image image.Image, data []byte) error {
	endpoint := defaultEndpoint()
	if printerEndpoint != "" {
		endpoint = printerEndpoint
	}

	p, err := printer.Open(endpoint)
	if err != nil {
		return err
	}
	defer p.Close()

	if err = p.StartRawDocument(fmt.Sprint(rand.Int())); err != nil {
		return err
	}

	if _, err := p.Write(data); err != nil {
		return err
	}

	if err = p.EndDocument(); err != nil {
		return err
	}

	return nil
}

// defaultEndpoint obtiene la impresora por defecto del sistema Windows.
func defaultEndpoint() string {
	name, err := printer.Default()
	if err != nil {
		return ""
	}
	return name
}
