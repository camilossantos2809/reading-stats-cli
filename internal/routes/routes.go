package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"reading-stats/internal/database"
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

func HomeHandler(w http.ResponseWriter, r *http.Request) {
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

func BooksListHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		data := database.GetBooks(db)
		var dataWithFormattedDate []database.GetBooksResult
		for _, book := range data {
			fmtDate, err := formatDate(book.Date)
			if err != nil {
				http.Error(w, "Erro ao formatar data", http.StatusInternalServerError)
				return
			}
			dataWithFormattedDate = append(dataWithFormattedDate, database.GetBooksResult{
				Name:     book.Name,
				Isbn:     book.Isbn,
				Date:     fmtDate,
				Progress: book.Progress,
			})
		}
		jsonData, err := json.Marshal(dataWithFormattedDate)
		if err != nil {
			http.Error(w, "Erro ao formatar JSON", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
