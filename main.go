package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/ramzsey/password-generator/internal/handler"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal(err)
	}

	h := handler.New(tmpl)

	http.HandleFunc("/", h.Home)
	http.HandleFunc("/generate", h.DynamicPassword)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Printf("Server starting on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
