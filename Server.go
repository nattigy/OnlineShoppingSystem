package main

import (
	_ "github.com/lib/pq"

	"net/http"
)

func main() {
	mux := http.NewServeMux()

	_ = http.ListenAndServe(":3000", mux)
}
