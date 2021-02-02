package sqlite

import (
	"database/sql"
	"fmt"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

// NewStore func
func NewStore(dataSourceName string) (*Store, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	return &Store{
		CatStore:         &CatStore{db},
		PostStore:        &PostStore{db},
		CommentStore:     &CommentStore{db},
		UserStore:        &UserStore{db},
		PostLikeStore:    &PostLikeStore{db},
		CommentLikeStore: &CommentLikeStore{db},
	}, nil
}

// Store type
type Store struct {
	*CatStore
	*PostStore
	*CommentStore
	*UserStore
	*PostLikeStore
	*CommentLikeStore
}
