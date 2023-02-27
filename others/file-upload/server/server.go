package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/upload", http.HandlerFunc(upload))
	http.ListenAndServe(":9090", nil)
}

func upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)
	log.Println(handler.Header)
}
