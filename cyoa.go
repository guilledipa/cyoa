// Program recreates the "Choose your own adventure" experience via a web
// application where each page will be a portion of the story, and at the end
// of every page the user will be given a series of options to choose from (or
// be told that they have reached the end of that particular story arc).
// https://github.com/gophercises/cyoa
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	rawBook = flag.String("raw_book", "./gopher.json", "Raw book filename.")
	port    = flag.Int("port", 8080, "Port to listen on.")
)

// book is a Choose Your Own Adventure book.
// Each key is the name of a chapter, and each value is a chapter struct.
type book map[string]chapter

// chapter is cyoa chapter.
type chapter struct {
	Title   string       `json:"title"`
	Story   []string     `json:"story"`
	Options []arcOptions `json:"options"`
}

// arcOptions are the choices offered at the end of a chapter.
// Text is the text shown to the readers; Arc is the key of a chapter in the
// Book struct.
type arcOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func parseJSON(jsonStoryFile string) (book, error) {
	jsonData, err := ioutil.ReadFile(jsonStoryFile)
	if err != nil {
		return nil, err
	}
	b := new(book)
	if err := json.Unmarshal(jsonData, b); err != nil {
		return nil, err
	}
	return *b, nil
}

func main() {
	flag.Parse()

	b, err := parseJSON(*rawBook)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}
