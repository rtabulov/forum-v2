package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/rtabulov/forum-v2"
	uuid "github.com/satori/go.uuid"
)

// CommentLikeStore qwe
type CommentLikeStore struct {
	*sql.DB
}

// GetCommentLike func
func (s *PostLikeStore) GetCommentLike(l *forum.CommentLike) (*forum.CommentLike, error) {
	like := &forum.CommentLike{}
	r := s.DB.QueryRow(`SELECT user_id, comment_id, up FROM comment_likes
	WHERE comment_id = ? AND user_id = ?`, l.CommentID, l.UserID)
	if err := r.Scan(&like.UserID, &like.CommentID, &like.Up); err != nil {
		return nil, fmt.Errorf("error getting comment like: %w", err)
	}
	return like, nil
}

// LikesByComment func
func (s *CommentLikeStore) LikesByComment(commentID uuid.UUID) ([]forum.CommentLike, error) {
	pp := []forum.CommentLike{}

	rr, err := s.DB.Query(`SELECT comment_id, user_id, up FROM comment_likes
	WHERE comment_id = ?`, commentID)
	if err != nil {
		return nil, fmt.Errorf(`error getting comment likes: %w`, err)
	}

	for rr.Next() {
		p := forum.CommentLike{}
		err := rr.Scan(&p.CommentID, &p.UserID, &p.Up)
		if err != nil {
			return nil, fmt.Errorf(`error getting comment likes: %w`, err)
		}
		pp = append(pp, p)
	}

	return pp, nil
}

// CreateCommentLike func
func (s *CommentLikeStore) CreateCommentLike(p *forum.CommentLike) error {

	_, err := s.DB.Exec(`INSERT INTO comment_likes (comment_id, user_id, up) 
	VALUES (?, ?, ?)`, p.CommentID, p.UserID, p.Up)
	if err != nil {
		return fmt.Errorf(`error creating comment like: %w`, err)
	}

	return nil
}

// DeleteCommentLike func
func (s *CommentLikeStore) DeleteCommentLike(l *forum.CommentLike) error {
	_, err := s.DB.Exec(`DELETE FROM comment_likes WHERE comment_id = ? AND user_id = ?`, l.CommentID, l.UserID)
	if err != nil {
		return fmt.Errorf(`error deleting comment like: %w`, err)
	}

	return nil
}
