// Copyright 2015 Google Inc. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/appengine"
)

var clientID string

func init() {

	clientID = os.Getenv("DEVDOT_OAUTH_CLIENT_ID")
	log.Printf("init: %s", clientID)
}

func main() {
	http.HandleFunc("/", handler)
	log.Print("Listening on port 8080")
	appengine.Main()
	// log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello, world! (%s)", clientID)))
}
