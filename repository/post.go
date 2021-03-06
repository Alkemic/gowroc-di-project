package repository

import (
	"database/sql"
	"fmt"
	"time"
)

type Post struct {
	Id        int
	Title     string
	Content   string
	CreatedAt time.Time
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{
		db: db,
	}
}

func (s repository) FetchEntries() ([]Post, error) {
	entries := []Post{}
	rows, err := s.db.Query(`SELECT * FROM posts;`)
	if err != nil {
		return nil, fmt.Errorf("error while fetching entries: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		entry := Post{}
		err := rows.Scan(&entry.Id, &entry.Title, &entry.Content, &entry.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error while scaning: %w", err)
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (s repository) GetEntry(id int) (Post, error) {
	entry := Post{}
	row := s.db.QueryRow(`SELECT * FROM posts WHERE id = ?;`, id)
	err := row.Scan(&entry.Id, &entry.Title, &entry.Content, &entry.CreatedAt)
	if err != nil {
		return Post{}, fmt.Errorf("error while scaning: %w", err)
	}
	return entry, nil
}
