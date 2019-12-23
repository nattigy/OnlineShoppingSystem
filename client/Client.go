package main

import (
	"github.com/nattigy/parentschoolcommunicationsystem/client/request"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("client/assets"))
	mux.Handle("/client/assets/", http.StripPrefix("/client/assets/", fs))

	request.NewStudentRequest(mux)

	_ = http.ListenAndServe(":3002", mux)
}
