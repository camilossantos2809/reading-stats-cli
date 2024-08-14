package main

import (
	"database/sql"
	"net/http"
	"text/template"
	"time"
)

func formatDate(dateStr string) (string, error) {
	// layout da string de entrada
	inputLayout := "2006-01-02"
	// layout da string de saída
	outputLayout := "02/01/2006"

	t, err := time.Parse(inputLayout, dateStr)
	if err != nil {
		return "", err
	}
	return t.Format(outputLayout), nil
}

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

func booksListHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("src/templates/books.html")
		if err != nil {
			http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
			return
		}
		data := GetBooks(db)
		var dataWithFormattedDate []GetBooksResult
		for _, book := range data {
			fmtDate, err := formatDate(book.Date)
			if err != nil {
				http.Error(w, "Erro ao formatar data", http.StatusInternalServerError)
				return
			}
			dataWithFormattedDate = append(dataWithFormattedDate, GetBooksResult{
				Name:     book.Name,
				Isbn:     book.Isbn,
				Date:     fmtDate,
				Progress: book.Progress,
			})
		}
		err = tmpl.Execute(w, dataWithFormattedDate)
		if err != nil {
			http.Error(w, "Erro ao renderizar template", http.StatusInternalServerError)
		}
	}
}
