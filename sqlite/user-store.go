package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/rtabulov/forum-v2"
	uuid "github.com/satori/go.uuid"
)

// UserStore qwe
type UserStore struct {
	*sql.DB
}

// User func
func (s *UserStore) User(id uuid.UUID) (*forum.User, error) {
	c := forum.User{}

	r := s.DB.QueryRow(`SELECT user_id, username, email, password, avatar, created_at FROM users WHERE user_id = ?`, id)
	err := r.Scan(&c.ID, &c.Username, &c.Email, &c.Password, &c.Avatar, &c.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf(`error getting user: %w`, err)
	}

	return &c, nil
}

// UserByUsername func
func (s *UserStore) UserByUsername(username string) (*forum.User, error) {
	c := forum.User{}

	r := s.DB.QueryRow(`SELECT user_id, username, email, password, avatar, created_at FROM users WHERE username = ?`, username)
	err := r.Scan(&c.ID, &c.Username, &c.Email, &c.Password, &c.Avatar, &c.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf(`error getting user: %w`, err)
	}

	return &c, nil
}

// CreateUser func
func (s *UserStore) CreateUser(c *forum.User) error {
	if c.ID == (uuid.UUID{}) {
		c.ID = uuid.NewV4()
	}

	_, err := s.DB.Exec(`INSERT INTO users (user_id, username, email, password, avatar) 
	VALUES (?, ?, ?, ?, ?)`, c.ID, c.Username, c.Email, c.Password, c.Avatar)
	if err != nil {
		return fmt.Errorf(`error creating user: %w`, err)
	}

	return nil
}

// UpdateUser func
func (s *UserStore) UpdateUser(c *forum.User) error {
	_, err := s.DB.Exec(`UPDATE users username = ?, email = ?, password = ?, avatar = ?, created_at = ?
	WHERE user_id = ?`,
		c.Username, c.Email, c.Password, c.Avatar, c.CreatedAt, c.ID)
	if err != nil {
		return fmt.Errorf(`error updating user: %w`, err)
	}
	return nil
}

// DeleteUser func
func (s *UserStore) DeleteUser(id uuid.UUID) error {
	_, err := s.DB.Exec(`DELETE FROM users WHERE user_id = ?`, id)
	if err != nil {
		return fmt.Errorf(`error deleting user: %w`, err)
	}

	return nil
}
