package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("GET /books", booksListHandler(db))
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
		log.Fatal(err)
	}
}

func parseSkoobHtml(db *sql.DB) {
	htmlDataList := getHtmlData(db)

	for _, htmlData := range htmlDataList {
		doc, err := html.Parse(strings.NewReader(htmlData.Html))
		if err != nil {
			fmt.Println("Error parsing HTML:", err)
			continue
		}
		progresses := ExtractProgress(doc)
		fmt.Printf("(ID: %s)\n", htmlData.BookId)
		for _, progress := range progresses {
			AddBookReadingProgress(db, AddBookReadingProgressParams{
				BookId: htmlData.BookId, DateRead: progress.Date, Progress: progress.Progress,
			})
		}
		UpdatePagesRead(db)
	}
}
