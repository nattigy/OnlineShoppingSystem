package main

import (
	"html/template"
	"log"
	"net/http"
)

var templ = template.Must(template.ParseGlob("deliverable/templates/*.html"))

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("deliverable/assets"))
	mux.Handle("/deliverable/assets/", http.StripPrefix("/deliverable/assets/", fs))

	mux.HandleFunc("/", Home)

	log.Fatal(http.ListenAndServe(":3000", mux))
}

func Home(w http.ResponseWriter, r *http.Request) {

}
