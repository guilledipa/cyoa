// Program recreates the "Choose your own adventure" experience via a web
// application where each page will be a portion of the story, and at the end
// of every page the user will be given a series of options to choose from (or
// be told that they have reached the end of that particular story arc).
// https://github.com/gophercises/cyoa
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/guilledipa/cyoa"
)

var (
	rawBook = flag.String("raw_book", "./gopher.json", "Raw book filename.")
	port    = flag.Int("port", 8080, "Port to listen on.")
)

func main() {
	flag.Parse()

	b := new(cyoa.BookHandler)
	if err := b.ParseJSON(*rawBook); err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.Handle("/intro", b)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))

	fmt.Println(b)
}
