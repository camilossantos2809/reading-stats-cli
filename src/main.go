package main

import (
	"database/sql"
	"fmt"
	"log"
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
	}
}
