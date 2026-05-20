package main

import (
	"html/template"
	"log"
	"net/http"
)

var tmpl *template.Template

func startServer() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/convert", handleConvert)

	log.Print("Listening on :3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	err := tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// Parse templates once when package initializes
func init() {
	var err error
	tmpl, err = template.ParseFiles("./static/index.html")
	if err != nil {
		log.Fatal("Failed to parse template:", err)
	}
	log.Println("Templates parsed successfully")
}
