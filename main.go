package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"google.golang.org/api/googleapi/transport"

	language "cloud.google.com/go/language/apiv1"

	"github.com/smaxwellstewart/sentiment/analyze"
	"github.com/smaxwellstewart/sentiment/api"
	"github.com/urfave/negroni"
)

var (
	addr = flag.String("addr", ":8000", "http socket address")
	key  = flag.String("key", "", "google api key")
)

func main() {
	flag.Parse()

	start(*addr, *key)
	fmt.Printf("Starting service @ %s with auth key %s \n", *addr, *key)
}

// start expects flags to be already parsed or manually set
func start(addr, key string) {

	// ctx := context.Background()

	ctx := context.WithValue(context.Background(), oauth2.HTTPClient, &http.Client{
		Transport: &transport.APIKey{Key: key},
	})
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
