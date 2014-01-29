package main

import (
	"fmt"
	"html/template"
	"net/http"
	// "path/filepath"
)

func thandle(rw http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("./test.html")
	t.Execute(rw, nil)
}

func main() {

	http.HandleFunc("/", thandle)
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./js"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

	http.ListenAndServe("127.0.0.1:9999", func() http.Handler {
		fmt.Println("server Running ")
		return nil
	}())

}
