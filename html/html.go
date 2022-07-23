package html

import (
	"embed"
	"io"
	"log"
	"text/template"
)

//go:embed templates/*
var files embed.FS
var page = parse("templates/redirect.html")

type RedirectPageParams struct {
	Title        string
	RedirectUrl  string
	ExtraMessage string
	RedirectIn   int
}

func Redirect(w io.Writer, p RedirectPageParams) error {
	return page.Execute(w, p)
}

func parse(file string) *template.Template {
	t, err := template.ParseFS(files, file)
	if err != nil {
		log.Fatal(err)
	}

	return t
}
