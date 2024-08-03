package main

import (
	"database/sql"
	"log"
)

type GetHtmlData struct {
	BookId int
	Html   string
}

func getHtmlData(db *sql.DB) []GetHtmlData {
	rows, err := db.Query("SELECT book_id, html FROM 'reading_progress_html'")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var htmlDataList []GetHtmlData
	for rows.Next() {
		var bookId int
		var htmlData string

		err := rows.Scan(&bookId, &htmlData)
		if err != nil {
			log.Fatal(err)
		}
		htmlDataList = append(htmlDataList, GetHtmlData{BookId: bookId, Html: htmlData})
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return htmlDataList
}
