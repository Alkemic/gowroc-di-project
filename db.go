package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	if db, err = sql.Open("mysql", "root:root@/blog?parseTime=true"); err != nil {
		log.Fatalln("error connecting to database", err)
	}
}

type post struct {
	Id        int
	Title     string
	Content   string
	CreatedAt time.Time
}

func fetchEntries() ([]*post, error) {
	entries := []*post{}
	rows, err := db.Query(`SELECT * FROM posts;`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		entry := &post{}
		err := rows.Scan(&entry.Id, &entry.Title, &entry.Content, &entry.CreatedAt)
		if err != nil {
			return nil, err
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func getEntry(id int) (*post, error) {
	entry := &post{}
	row := db.QueryRow(`SELECT * FROM posts WHERE id = ?;`, id)
	err := row.Scan(&entry.Id, &entry.Title, &entry.Content, &entry.CreatedAt)
	if err != nil {
		return nil, err
	}
	return entry, nil
}
