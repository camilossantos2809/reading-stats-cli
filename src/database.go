package main

import (
	"database/sql"
	"log"
)

func getHtmlData(db *sql.DB) []string {
	rows, err := db.Query("SELECT book_id, html FROM 'reading_progress_html'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var htmlDataList []string
	for rows.Next() {
		var bookId int
		var htmlData string

		err := rows.Scan(&bookId, &htmlData)
		if err != nil {
			log.Fatal(err)
		}
		htmlDataList = append(htmlDataList, htmlData)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return htmlDataList
}
