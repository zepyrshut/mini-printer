package handlers

import (
	"bytes"
	"github.com/zepyrshut/mini-printer/printer"
	"github.com/zepyrshut/mini-printer/printer/windows"
	"html/template"
	"net/http"
)

// Index es el manejeador para la página de inicio.
func Index(w http.ResponseWriter, _ *http.Request) {
	renderTemplate(w, "form")
}

// Print es el manejador que se encarga de obtener el valor del formulario e imprimirlo.
func Print(w http.ResponseWriter, r *http.Request) {
	ui := r.FormValue("data")

	buf := &bytes.Buffer{}
	parsedTemplate, _ := template.ParseFiles("./templates/data.tmpl")

	type H map[string]any
	data := H{
		"data": ui,
	}

	parsedTemplate.Execute(buf, data)

	result := printer.ProcessPrint(buf.String())
	wd := windows.Direct{}
	wd.Print("", *result.Image, result.Data)

	Index(w, r)
}

// renderTemplate es un ayudante que se encarga de renderizar la plantilla.
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/form.tmpl")
	parsedTemplate.Execute(w, nil)
}
