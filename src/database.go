package main

import (
	"database/sql"
	"log"
)

type GetHtmlData struct {
	BookId string
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
		var bookId string
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

type AddBookReadingProgressParams struct {
	BookId   string
	DateRead string
	Progress int
}

func AddBookReadingProgress(db *sql.DB, params AddBookReadingProgressParams) {
	_, err := db.Exec(`
		INSERT INTO book_reading_progress (book_id, date_read, progress) VALUES (?, ?, ?)
			ON CONFLICT (book_id, date_read) DO UPDATE SET progress = excluded.progress
				WHERE book_reading_progress.progress < excluded.progress;
		`, params.BookId, params.DateRead, params.Progress)
	if err != nil {
		log.Fatal(err)
	}
}

// Atualiza a quantidade de páginas lidas no dia, com base na última página lida no dia anterior
func UpdatePagesRead(db *sql.DB) {
	_, err := db.Exec(`
		WITH previous_reading AS (
			SELECT 
				book_id,
				date_read,
				LAG (progress, 1, 0) OVER (PARTITION BY book_id ORDER BY date_read) AS last_progress
			FROM book_reading_progress
		)
		UPDATE book_reading_progress
		SET progress_previous = pr.last_progress,
			pages_read = progress - pr.last_progress
		FROM previous_reading pr
		WHERE book_reading_progress.book_id = pr.book_id
			AND book_reading_progress.date_read = pr.date_read;
		`)
	if err != nil {
		log.Fatal(err)
	}
}
