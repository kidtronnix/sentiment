package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	language "cloud.google.com/go/language/apiv1"

	"github.com/smaxwellstewart/sentiment/analyze"
	"github.com/smaxwellstewart/sentiment/api"
	"github.com/urfave/negroni"
)

var (
	addr = flag.String("addr", ":8000", "http socket address")
)

func main() {
	flag.Parse()

	start(":8000")
}

// start expects flags to be already parsed or manually set
func start(addr string) {

	ctx := context.Background()
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	handler := api.Handler{
		Anlyzr: &analyze.GoogleAnalyzer{
			Client: client,
			Ctx:    ctx,
		},
	}

	// add middleware to our router
	n := negroni.Classic()
	n.UseHandler(handler)

	http.Handle("/api", n)
	log.Fatal(http.ListenAndServe(addr, nil))
}
