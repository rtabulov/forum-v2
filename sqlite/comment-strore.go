package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/rtabulov/forum-v2"
	uuid "github.com/satori/go.uuid"
)

// CommentStore qwe
type CommentStore struct {
	*sql.DB
}

// Comment func
func (s *CommentStore) Comment(id uuid.UUID) (*forum.Comment, error) {
	p := forum.Comment{}

	r := s.DB.QueryRow(`SELECT comment_id, user_id, post_id, body, created_at
	FROM comments WHERE comment_id = ?`, id)
	err := r.Scan(&p.ID, &p.UserID, &p.PostID, &p.Body, &p.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf(`error getting comment: %w`, err)
	}

	return &p, nil
}

// CommentsByPost func
func (s *CommentStore) CommentsByPost(postID uuid.UUID) ([]forum.Comment, error) {
	pp := []forum.Comment{}

	rr, err := s.DB.Query(`SELECT comment_id, user_id, post_id, body, created_at FROM comments
	WHERE post_id = ?`, postID)
	if err != nil {
		return nil, fmt.Errorf(`error getting comments: %w`, err)
	}

	for rr.Next() {
		p := forum.Comment{}
		err := rr.Scan(&p.ID, &p.UserID, &p.PostID, &p.Body, &p.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf(`error getting comments: %w`, err)
		}

		pp = append(pp, p)
	}

	return pp, nil
}

// CreateComment func
func (s *CommentStore) CreateComment(p *forum.Comment) error {
	if p.ID == (uuid.UUID{}) {
		p.ID = uuid.Must(uuid.NewV4())
	}

	_, err := s.DB.Exec(`INSERT INTO comments (comment_id, user_id, post_id, body) 
	VALUES (?, ?, ?, ?)`, p.ID, p.UserID, p.PostID, p.Body)
	if err != nil {
		return fmt.Errorf(`error creating comment: %w`, err)
	}

	return nil
}

// UpdateComment func
func (s *CommentStore) UpdateComment(p *forum.Comment) error {
	_, err := s.DB.Exec(`UPDATE user_id = ?, post_id = ?, body = ?
	WHERE comment_id = ?`,
		p.UserID, p.PostID, p.Body, p.ID)
	if err != nil {
		return fmt.Errorf(`error updating comment: %w`, err)
	}
	return nil
}

// DeleteComment func
func (s *CommentStore) DeleteComment(id uuid.UUID) error {
	_, err := s.DB.Exec(`DELETE FROM comments WHERE comment_id = ?`, id)
	if err != nil {
		return fmt.Errorf(`error deleting comment: %w`, err)
	}

	return nil
}
