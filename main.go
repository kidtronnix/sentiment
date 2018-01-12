package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/api/option"

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
	fmt.Printf("Starting service @ %s with auth key %s \n", *addr, *key)
	start(*addr, *key)

}

// start expects flags to be already parsed or manually set
func start(addr, key string) {

	// if you have google default application creds setup this will add auth creds to ctx
	ctx := context.Background()

	// let client use API key (little unsure if this is actually used)
	// TODO! add client logging so can debug reqs easily.
	clientApiKeyOpt := option.WithAPIKey(key)
	client, err := language.NewClient(ctx, clientApiKeyOpt)
	if err != nil {
		log.Fatal(err)
	}

	// construct http handler
	handler := api.Handler{
		Anlyzr: &analyze.GoogleAnalyzer{
			Client: client,
			Ctx:    ctx,
		},
	}

	// add useful middleware to handler: logging, panic recovery, etc...
	n := negroni.Classic()
	n.UseHandler(handler)

	// serve handler on /api route
	http.Handle("/api", n)
	log.Fatal(http.ListenAndServe(addr, nil))
}
