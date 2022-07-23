package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/rshramko/redirector/html"

	"net/http"
)

var port int
var baseRedirectUrl string
var title string
var redirectIn int
var extraInfo string

func init() {
	flag.IntVar(&port, "port", 8080, "listen on port")
	flag.StringVar(&baseRedirectUrl, "url", "http://localhost", "base redirect url")
	flag.StringVar(&title, "title", "Service", "name of the service that has been moved")
	flag.IntVar(&redirectIn, "wait", 15, "wait before redirecting")
	flag.StringVar(&extraInfo, "extra", "", "extra information to show")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	url := baseRedirectUrl + r.URL.Path
	if r.URL.RawQuery != "" {
		url += "?" + r.URL.RawQuery
	}

	p := html.RedirectPageParams{
		Title:        title,
		RedirectUrl:  url,
		RedirectIn:   redirectIn,
		ExtraMessage: extraInfo,
	}

	html.Redirect(w, p)
}

func main() {
	flag.Parse()
	fmt.Printf("Starting redirector on port %d...\n", port)
	http.HandleFunc("/", redirect)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
