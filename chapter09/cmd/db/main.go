package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "database.sqlite")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	queries := []string{
		`CREATE TABLE IF NOT EXISTS authors (author_id TEXT, author TEXT, PRIMARY KEY (author_id))`,
		`CREATE TABLE IF NOT EXISTS contents (author_id TEXT, title_id TEXT, title TEXT, content TEXT, PRIMARY KEY (author_id, title_id))`,
		`CREATE VIRTUAL TABLE IF NOT EXISTS contents_fts USING fts4(words)`,
	}
	for _, query := range queries {
		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	content := "あばばばば。あばばばば。あばばばば。"
	res, err := db.Exec(`INSERT INTO contents(author_id, title_id, title, content) VALUES(?, ?, ?, ?)`, "000879", "14", "あばばばば", content)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", res)
	// docID, err := res.LastInsertId()
}
