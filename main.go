package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//go:embed templates
var tplFolder embed.FS

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(tplFolder, "templates/index.html")
	if err != nil {
		log.Fatal(err)
	}
	println("Embedded the index.html file")
	tmpl.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", indexHandler)

	fmt.Println("Server starting on port 8000!")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
