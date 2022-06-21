package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".json"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &Page{
		Title: title,
		Body:  body,
	}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/"):]
	p, _ := LoadPage(title)

	fmt.Fprintf(w, "%s", p.Body)
}

func public() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", viewHandler)
	mux.Handle("/public/", public())
	log.Fatal(http.ListenAndServe(":8080", mux))
}
