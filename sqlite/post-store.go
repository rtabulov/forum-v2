package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/rtabulov/forum-v2"
	uuid "github.com/satori/go.uuid"
)

// PostStore qwe
type PostStore struct {
	*sql.DB
}

// Post func
func (s *PostStore) Post(id uuid.UUID) (*forum.Post, error) {
	p := forum.Post{}

	r := s.DB.QueryRow(`SELECT post_id, user_id, title, body, created_at
	FROM posts WHERE post_id = ?`, id)
	err := r.Scan(&p.ID, &p.UserID, &p.Title, &p.Body, &p.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf(`error getting post: %w`, err)
	}

	return &p, nil
}

// Posts func
func (s *PostStore) Posts() ([]forum.Post, error) {
	pp := []forum.Post{}

	rr, err := s.DB.Query(`SELECT post_id, user_id, title, body, created_at 
	FROM posts 
	ORDER BY created_at DESC`)
	if err != nil {
		return nil, fmt.Errorf(`error getting posts: %w`, err)
	}

	for rr.Next() {
		p := forum.Post{}
		err := rr.Scan(&p.ID, &p.UserID, &p.Title, &p.Body, &p.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf(`error getting posts: %w`, err)
		}

		pp = append(pp, p)
	}

	return pp, nil
}

// PostsByCat func
// func (s *PostStore) PostsByCat(catIDs uuid.UUID) ([]forum.Post, error) {
// 	pp := []forum.Post{}

// 	rr, err := s.DB.Query(`SELECT post_id, cat_id, user_id, title, body, created_at FROM posts
// 	WHERE cat_id = ?
// 	ORDER BY created_at DESC`, catID)
// 	if err != nil {
// 		return nil, fmt.Errorf(`error getting posts: %w`, err)
// 	}

// 	for rr.Next() {
// 		p := forum.Post{}
// 		err := rr.Scan(&p.ID, &p.CatID, &p.UserID, &p.Title, &p.Body, &p.CreatedAt)
// 		if err != nil {
// 			return nil, fmt.Errorf(`error getting posts: %w`, err)
// 		}
// 		pp = append(pp, p)
// 	}

// 	return pp, nil
// }

// PostsByUser func
func (s *PostStore) PostsByUser(userID uuid.UUID) ([]forum.Post, error) {
	pp := []forum.Post{}

	rr, err := s.DB.Query(`SELECT post_id, user_id, title, body, created_at FROM posts
	WHERE user_id = ?
	ORDER BY created_at DESC`, userID)
	if err != nil {
		return nil, fmt.Errorf(`error getting posts: %w`, err)
	}

	for rr.Next() {
		p := forum.Post{}
		err := rr.Scan(&p.ID, &p.UserID, &p.Title, &p.Body, &p.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf(`error getting posts: %w`, err)
		}
		pp = append(pp, p)
	}

	return pp, nil
}

// CreatePost func
func (s *PostStore) CreatePost(p *forum.Post) error {
	if p.ID == (uuid.UUID{}) {
		p.ID = uuid.NewV4()
	}

	_, err := s.DB.Exec(`INSERT INTO posts (post_id, user_id, title, body) 
	VALUES (?, ?, ?, ?)`, p.ID, p.UserID, p.Title, p.Body)
	if err != nil {
		return fmt.Errorf(`error creating post: %w`, err)
	}

	return nil
}

// UpdatePost func
func (s *PostStore) UpdatePost(p *forum.Post) error {
	_, err := s.DB.Exec(`UPDATE posts post_id = ?,
	user_id = ?, title = ?, body = ?, created_at = ?
	WHERE post_id = ?`,
		p.UserID, p.Title, p.Body, p.CreatedAt, p.ID)
	if err != nil {
		return fmt.Errorf(`error updating post: %w`, err)
	}
	return nil
}

// DeletePost func
func (s *PostStore) DeletePost(id uuid.UUID) error {
	_, err := s.DB.Exec(`DELETE FROM posts WHERE post_id = ?`, id)
	if err != nil {
		return fmt.Errorf(`error deleting post: %w`, err)
	}

	return nil
}
