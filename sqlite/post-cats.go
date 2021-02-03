package sqlite

import (
	"fmt"

	"github.com/rtabulov/forum-v2"
	uuid "github.com/satori/go.uuid"
)

// CatPosts func
func (s *Store) CatPosts(postID uuid.UUID) ([]forum.Post, error) {
	pp := []forum.Post{}

	rows, err := s.CatStore.DB.Query(`SELECT posts.post_id, user_id, title, body, created_at
	FROM post_cats 
	JOIN posts ON posts.post_id = post_cats.post_id  
	WHERE cat_id = ?
	ORDER BY created_at DESC`, postID)
	if err != nil {
		return nil, fmt.Errorf("error getting cat posts: %w", err)
	}

	for rows.Next() {
		p := forum.Post{}

		err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Body, &p.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("error getting cat posts: %w", err)
		}

		pp = append(pp, p)
	}

	return pp, nil
}

// PostCats func
func (s *Store) PostCats(postID uuid.UUID) ([]forum.Cat, error) {
	cc := []forum.Cat{}

	rows, err := s.CatStore.DB.Query(`SELECT cats.cat_id, title, description
	FROM post_cats 
	JOIN cats ON cats.cat_id = post_cats.cat_id  
	WHERE post_id = ?`, postID)
	if err != nil {
		return nil, fmt.Errorf("error getting post cats: %w", err)
	}

	for rows.Next() {
		c := forum.Cat{}
		err := rows.Scan(&c.ID, &c.Title, &c.Description)
		if err != nil {
			return nil, fmt.Errorf("error getting post cats: %w", err)
		}
		cc = append(cc, c)
	}

	return cc, nil
}

// CratePostCats func
func (s *Store) CratePostCats(postID uuid.UUID, catIDs []uuid.UUID) error {
	for _, cat := range catIDs {
		_, err := s.CatStore.DB.Exec(`INSERT INTO post_cats 
		(post_id, cat_id)
		VALUES (?, ?)`, postID, cat)
		if err != nil {
			return fmt.Errorf("error creating post cats: %w", err)
		}
	}

	return nil
}
