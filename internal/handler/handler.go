package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/ramzsey/password-generator/internal/generator"
)

type Handler struct {
	tmpl *template.Template
}

func New(tmpl *template.Template) *Handler {
	return &Handler{tmpl: tmpl}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	passwords, err := generator.Generate()
	if err != nil {
		http.Error(w, "Failed to generate passwords", http.StatusInternalServerError)
		return
	}

	h.tmpl.Execute(w, passwords)
}

func (h *Handler) DynamicPassword(w http.ResponseWriter, r *http.Request) {
	lengthStr := r.URL.Query().Get("length")
	length, err := strconv.Atoi(lengthStr)
	if err != nil || length < 1 || length > 50 {
		length = 20
	}

	password, err := generator.GenerateSingle(length)
	if err != nil {
		http.Error(w, "Failed to generate password", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<code class="password-text">%s</code>`, password)
}
