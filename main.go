package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Person struct {
	fName string
}

type MyHandler struct {
}

func (this *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[1:]
	log.Println(path)
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404 My Friend - " + http.StatusText(404)))
	}
}

func (p *Person) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("First name : " + p.fName))
}

func main() {

	//p1 := &Person{fName: "Nati"}

	//mux := http.NewServeMux()
	http.Handle("/", new(MyHandler))

	http.ListenAndServe(":8080", nil)
}

func Home(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello universe"))
}
