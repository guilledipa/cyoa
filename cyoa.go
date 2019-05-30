// Package cyoa contains fuctions to create a cyoa book
// https://github.com/gophercises/cyoa
package cyoa

import (
	"encoding/json"
	"io/ioutil"
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
