package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reading-stats/internal/database"
	"reading-stats/internal/routes"
	"reading-stats/internal/webscraping"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/net/html"
)

func main() {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := http.NewServeMux()
	mux.HandleFunc("/", routes.HomeHandler)
	mux.HandleFunc("GET /books", routes.BooksListHandler(db))
	mux.HandleFunc("POST /parseSkoobHtml", func(w http.ResponseWriter, r *http.Request) {
		parseSkoobHtml(db)
		json.NewEncoder(w).Encode("{response: 'ok'}")
	})
	port := ":8080"
	server := &http.Server{
		Addr:    port,
		Handler: mux,
	}

	log.Println("Servidor iniciado em http://localhost:" + port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}

func parseSkoobHtml(db *sql.DB) {
	htmlDataList := database.GetHtmlData(db)

	for _, htmlData := range htmlDataList {
		doc, err := html.Parse(strings.NewReader(htmlData.Html))
		if err != nil {
			fmt.Println("Error parsing HTML:", err)
			continue
		}
		progresses := webscraping.ExtractProgress(doc)
		fmt.Printf("(ID: %s)\n", htmlData.BookId)
		for _, progress := range progresses {
			database.AddBookReadingProgress(db, database.AddBookReadingProgressParams{
				BookId: htmlData.BookId, DateRead: progress.Date, Progress: progress.Progress,
			})
		}
		database.UpdatePagesRead(db)
	}
}
