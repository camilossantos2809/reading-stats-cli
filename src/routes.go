package main

import (
	"net/http"
	"text/template"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	data := interface{}(nil)
	tmpl, err := template.ParseFiles("src/templates/index.html")
	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
	}
}
