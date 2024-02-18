package models

import (
	"database/sql"
	"time"
)

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define snippetModel type which wraps a sql.DB connection pool
type SnippetModel struct {
	DB *sql.DB
}

// insert a snippet

func (m *SnippetModel) Insert(title string, content string, created time.Time, expires int) (int, error) {

	stmt := `INSERT INTO snippets (title, content, created, expires)
						VALUES (?, ?, ?, DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

	// USE THE EXEC method to execute the insert statement
	result, err := m.DB.Exec(stmt, title, content, created, expires)
	if err != nil {
		return 0, err
	}

	// Get ID of the inserted record
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// returb a single snippet

func (m *SnippetModel) Get(id int) (*Snippet, error) {

	return nil, nil
}

// return most current snippet
func (m *SnippetModel) Latest() ([]*Snippet, error) {
	return nil, nil
}
