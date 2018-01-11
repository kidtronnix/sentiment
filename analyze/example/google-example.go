// https://github.com/GoogleCloudPlatform/golang-samples/blob/master/language/analyze/analyze.go

// Copyright 2016 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// Command analyze performs sentiment, entity, entity sentiment, and syntax analysis
// on a string of text via the Cloud Natural Language API.
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/context"

	// [START imports]
	language "cloud.google.com/go/language/apiv1"
	languagepb "google.golang.org/genproto/googleapis/cloud/language/v1"
	// [END imports]
)

func main() {
	if len(os.Args) < 1 {
		usage("Missing command.")
	}

	// [START init]
	ctx := context.Background()
	client, err := language.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// [END init]

	text := strings.Join(os.Args[1:], " ")
	if text == "" {
		usage("Missing text.")
	}

	// make call to sentiment service
	resp, err := analyzeSentiment(ctx, client, text)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", resp)

}

func usage(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	fmt.Fprintln(os.Stderr, "usage: analyze <text>")
	os.Exit(2)
}

func analyzeSentiment(ctx context.Context, client *language.Client, text string) (*languagepb.AnalyzeSentimentResponse, error) {
	return client.AnalyzeSentiment(ctx, &languagepb.AnalyzeSentimentRequest{
		Document: &languagepb.Document{
			Source: &languagepb.Document_Content{
				Content: text,
			},
			Type: languagepb.Document_PLAIN_TEXT,
		},
	})
}
