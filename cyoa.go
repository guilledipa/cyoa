// Package cyoa contains fuctions to create a cyoa book
// https://github.com/gophercises/cyoa
package cyoa

import (
	"encoding/json"
	"io/ioutil"
)

// Book is a Choose Your Own Adventure book.
// Each key is the name of a chapter, and each value is a chapter struct.
type Book map[string]Chapter

// Chapter is cyoa chapter.
type Chapter struct {
	Title   string       `json:"title"`
	Story   []string     `json:"story"`
	Options []ArcOptions `json:"options"`
}

// ArcOptions are the choices offered at the end of a chapter.
// Text is the text shown to the readers; Arc is the key of a chapter in the
// Book struct.
type ArcOptions struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

// ParseJSON reads a file containing a json formatted cyoa book and returns a
// Book map or error.
func ParseJSON(jsonStoryFile string) (Book, error) {
	jsonData, err := ioutil.ReadFile(jsonStoryFile)
	if err != nil {
		return nil, err
	}
	b := new(Book)
	if err := json.Unmarshal(jsonData, b); err != nil {
		return nil, err
	}
	return *b, nil
}
