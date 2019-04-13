package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Someone visited the website.")
	w.Header().Set("Content-Type", "text/html")

	if r.URL.Path == "/" {
		fmt.Fprint(w, "<h1>Homepage</h1>")
	} else {
		fmt.Fprint(w, "<h1>Other pages</h1>")
	}
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
