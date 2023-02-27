package db

import "database/sql"

// Store provides all functions to execute db queries and transactions.
type DBStore interface {
	Querier
}

// SQLStore provides all functions to execute SQL queries and transactions.
type SQLStore struct {
	*Queries
	db *sql.DB
}

// Initializes a new Store struct to execute queries and transactions.
func NewStore(db *sql.DB) DBStore {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
