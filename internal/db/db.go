package db

import "database/sql"

// CreateTable ensures the URLs table exist
func CreateTable(db *sql.DB) error {
	query := `
    CREATE TABLE IF NOT EXISTS urls (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        short_url TEXT NOT NULL,
        original_url TEXT NOT NULL
    );`
	_, err := db.Exec(query)
	return err
}

// StoreURL inserts a new short URL and  the original URL into the database
func StoreURL(db *sql.DB, shortURL string, originalURL string) error {
	query := `INSERT INTO urls (short_url, original_url) VALUES (?, ?);`
	_, err := db.Exec(query, shortURL, originalURL)
	return err
}

// GetOriginalURL fetches the original URL by the short URL
func GetOriginalURL(db *sql.DB, shortURL string) (string, error) {
	var originalURL string
	query := `SELECT original_url FROM urls WHERE short_url = ? LIMIT 1;`
	err := db.QueryRow(query, shortURL).Scan(&originalURL)
	if err != nil {
		return "", err
	}
	return originalURL, nil
}
