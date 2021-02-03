package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/rtabulov/forum-v2"
	uuid "github.com/satori/go.uuid"
)

// CatStore qwe
type CatStore struct {
	*sql.DB
}

// Cat func
func (s *CatStore) Cat(id uuid.UUID) (*forum.Cat, error) {
	c := forum.Cat{}

	r := s.DB.QueryRow(`SELECT cat_id, title, description FROM cats WHERE cat_id = ?`, id)
	err := r.Scan(&c.ID, &c.Title, &c.Description)
	if err != nil {
		return nil, fmt.Errorf(`error getting cat: %w`, err)
	}

	return &c, nil
}

// Cats func
func (s *CatStore) Cats() ([]forum.Cat, error) {
	cc := []forum.Cat{}

	rr, err := s.DB.Query(`SELECT cat_id, title, description FROM cats`)
	if err != nil {
		return nil, fmt.Errorf(`error getting cats: %w`, err)
	}

	for rr.Next() {
		c := forum.Cat{}
		err := rr.Scan(&c.ID, &c.Title, &c.Description)
		if err != nil {
			return nil, fmt.Errorf(`error getting cats: %w`, err)
		}
		cc = append(cc, c)
	}

	return cc, nil
}

// CreateCat func
func (s *CatStore) CreateCat(c *forum.Cat) error {
	if c.ID == (uuid.UUID{}) {
		c.ID = uuid.Must(uuid.NewV4())
	}

	_, err := s.DB.Exec(`INSERT INTO cats (cat_id, title, description) 
	VALUES (?, ?, ?)`, c.ID, c.Title, c.Description)
	if err != nil {
		return fmt.Errorf(`error creating cat: %w`, err)
	}

	return nil
}

// UpdateCat func
func (s *CatStore) UpdateCat(c *forum.Cat) error {
	_, err := s.DB.Exec(`UPDATE cats title = ?, description = ?
	WHERE cat_id = ?`,
		c.Title, c.Description, c.ID)
	if err != nil {
		return fmt.Errorf(`error updating cat: %w`, err)
	}
	return nil
}

// DeleteCat func
func (s *CatStore) DeleteCat(id uuid.UUID) error {
	_, err := s.DB.Exec(`DELETE FROM cats WHERE cat_id = ?`, id)
	if err != nil {
		return fmt.Errorf(`error deleting cat: %w`, err)
	}

	return nil
}
