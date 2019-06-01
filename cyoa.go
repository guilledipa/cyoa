// Package cyoa contains fuctions to create a cyoa book
// https://github.com/gophercises/cyoa
package cyoa

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var defaultTemplate = `
<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Choose Your Own Adventure</title>
  </head>
  <body>
    <section class="page">
      <h1>{{.Title}}</h1>
      {{range .Paragraphs}}
        <p>{{.}}</p>
      {{end}}
      {{if .Options}}
        <ul>
        {{range .Options}}
          <li><a href="/{{.Chapter}}">{{.Text}}</a></li>
        {{end}}
        </ul>
      {{else}}
        <h3>The End</h3>
      {{end}}
    </section>
  </body>
</html>`

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

// parseJSON reads a file containing a json formatted cyoa book and returns a
// Book map or error.
func (b *BookHandler) parseJSON(jsonStoryFile string) error {
	jsonData, err := ioutil.ReadFile(jsonStoryFile)
	if err != nil {
		return fmt.Errorf("parseJSON read file: %v", err)
	}
	if err := json.Unmarshal(jsonData, b.Book); err != nil {
		return fmt.Errorf("parseJSON could not unmarshall: %v", err)
	}
	return nil
}

// NewBookHandler creates a BookHandler instance.
func NewBookHandler(jsonStoryFile string) (*BookHandler, error) {
	var b BookHandler
	if err := b.parseJSON(jsonStoryFile); err != nil {
		return nil, err
	}
	b.bookTemplate = template.Must(template.New("Default").Parse(defaultTemplate))
	return &b, nil
}
