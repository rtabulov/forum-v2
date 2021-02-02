package sqlite

import (
	"database/sql"
	"fmt"

	"github.com/rtabulov/forum-v2"
	uuid "github.com/satori/go.uuid"
)

// PostLikeStore qwe
type PostLikeStore struct {
	*sql.DB
}

// GetPostLike func
func (s *PostLikeStore) GetPostLike(l *forum.PostLike) (*forum.PostLike, error) {
	like := &forum.PostLike{}
	r := s.DB.QueryRow(`SELECT user_id, post_id, up FROM post_likes
	WHERE post_id = ? AND user_id = ?`, l.PostID, l.UserID)
	if err := r.Scan(&like.UserID, &like.PostID, &like.Up); err != nil {
		return nil, fmt.Errorf("error getting post like: %w", err)
	}
	return like, nil
}

// LikesByPost func
func (s *PostLikeStore) LikesByPost(postID uuid.UUID) ([]forum.PostLike, error) {
	pp := []forum.PostLike{}

	rr, err := s.DB.Query(`SELECT post_id, user_id, up FROM post_likes
	WHERE post_id = ?`, postID)
	if err != nil {
		return nil, fmt.Errorf(`error getting post likes: %w`, err)
	}

	for rr.Next() {
		p := forum.PostLike{}
		err := rr.Scan(&p.PostID, &p.UserID, &p.Up)
		if err != nil {
			return nil, fmt.Errorf(`error getting post likes: %w`, err)
		}
		pp = append(pp, p)
	}

	return pp, nil
}

// CreatePostLike func
func (s *PostLikeStore) CreatePostLike(p *forum.PostLike) error {

	_, err := s.DB.Exec(`INSERT INTO post_likes (post_id, user_id, up) 
	VALUES (?, ?, ?)`, p.PostID, p.UserID, p.Up)
	if err != nil {
		return fmt.Errorf(`error creating post like: %w`, err)
	}

	return nil
}

// DeletePostLike func
func (s *PostLikeStore) DeletePostLike(l *forum.PostLike) error {
	_, err := s.DB.Exec(`DELETE FROM post_likes WHERE post_id = ? AND user_id = ?`, l.PostID, l.UserID)
	if err != nil {
		return fmt.Errorf(`error deleting post like: %w`, err)
	}

	return nil
}
