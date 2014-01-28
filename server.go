package main

import (
	// "fmt"
	"html/template"
	"net/http"
	// "path/filepath"
)

func thandle(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("./test2.html")
	t.Execute(rw, nil)
}

func main() {

	http.HandleFunc("/", thandle)

	http.ListenAndServe("127.0.0.1:9999", nil)

}
