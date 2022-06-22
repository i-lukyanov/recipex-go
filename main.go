package main

import (
	"encoding/json"
	"fmt"
	mux2 "github.com/gorilla/mux"
	"log"
	"net/http"
)

func homeLink(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home of oats=)")
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if (*r).Method == "OPTIONS" {
		return
	}

	json.NewEncoder(w).Encode(Recipes)
}

func public() http.Handler {
	return http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
}

func main() {
	router := mux2.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	router.HandleFunc("/recipes", getAll)
	router.PathPrefix("/public/").Handler(public())
	//router.Handle("/public/", public())

	log.Fatal(http.ListenAndServe(":8080", router))
}
