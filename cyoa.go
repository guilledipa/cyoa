// Package cyoa contains fuctions to create a cyoa book
// https://github.com/gophercises/cyoa
package cyoa

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

// BookHandler has functions to render a web cyoa.
type BookHandler struct {
	Book
	bookTemplate *template.Template
}

func (b *BookHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	chapterTitle := strings.TrimLeft(r.URL.Path[1:], "/")
	if chapter, ok := b.Book[chapterTitle]; ok {
		err := b.bookTemplate.Execute(w, chapter)
		if err != nil {
			log.Printf("Template execute: %v", err)
			http.Error(w, "Something went wrong parsing a chapter...", http.StatusInternalServerError)
		}
	}
	http.Error(w, "Chapter not found.", http.StatusNotFound)
}

// ParseJSON reads a file containing a json formatted cyoa book and returns a
// Book map or error.
func (b *BookHandler) ParseJSON(jsonStoryFile string) error {
	jsonData, err := ioutil.ReadFile(jsonStoryFile)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(jsonData, b.Book); err != nil {
		return err
	}
	return nil
}
