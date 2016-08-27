package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"

	"github.com/gorilla/mux"
)

var sansItUp = `
<style>
* {
		font-family: "Comic Sans MS", "Comic Sans", cursive !important;
}
</style>
<audio autoplay>
		<source src="http://siobud.com/files/business/mp3.mp3" type="audio/mpeg"/>
</audio>
`

func rewriteMarkup(doc *goquery.Document) {
	doc.Find("body").Each(func(i int, body *goquery.Selection) {
		body.PrependHtml(sansItUp)
	})
}

func comicDansHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	var (
		err error
		doc *goquery.Document
	)

	if len(url) == 0 {
		doc, err = goquery.NewDocument("http://siobud.com/files/business")
	} else {
		doc, err = goquery.NewDocument(r.URL.Query().Get("url"))
	}

	if err == nil {
		rewriteMarkup(doc)
		html, err := doc.Html()
		if err == nil {
			fmt.Fprintf(w, html)
		}
	}
}

type notFoundHandler struct {
}

func (h *notFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	doc, err := goquery.NewDocument("http://siobud.com/files/business")
	if err == nil {
		rewriteMarkup(doc)
		html, err := doc.Html()
		if err == nil {
			fmt.Fprintf(w, html)
		}
	}
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", comicDansHandler)
	myRouter.NotFoundHandler = &notFoundHandler{}

	log.Fatal(http.ListenAndServe(":4444", myRouter))
}
