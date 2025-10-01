package main

import (
	"html/template"
	"net/http"

	"github.com/sebgott/event-creator/internal/handlers"
)

var templates = template.Must(template.ParseGlob("internal/templates/*.html"))

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "base.html", nil)
		templates.ExecuteTemplate(w, "index.html", nil)
	})

	http.HandleFunc("/generate/topic", handlers.GenerateTopicHandler)
	http.HandleFunc("/generate/rolebinding", handlers.GenerateRoleBindingHandler)
	http.HandleFunc("/generate/all", handlers.GenerateAllHandler)
	http.HandleFunc("/form/add-config", handlers.AddConfigHandler)

	println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
