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
		doc, err := html.Parse(strings.NewReader(htmlData))
		if err != nil {
			fmt.Println("Error parsing HTML:", err)
			continue
		}

		progresses := ExtractProgress(doc)
		// fmt.Printf("(ID: %d)\n", bookId)
		for _, progress := range progresses {
			fmt.Printf("  {date: %q, progress: %d}\n", progress.Date, progress.Progress)
		}
	}

}
