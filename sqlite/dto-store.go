package sqlite

import (
	"github.com/rtabulov/forum-v2"
	uuid "github.com/satori/go.uuid"
)

// PostsDTO func
func (s *Store) PostsDTO() ([]forum.PostDTO, error) {
	pp, err := s.Posts()
	if err != nil {
		return nil, err
	}
	posts := make([]forum.PostDTO, len(pp))

	for i, p := range pp {
		cat, err := s.Cat(p.CatID)
		if err != nil {
			return nil, err
		}
		user, err := s.User(p.UserID)
		if err != nil {
			return nil, err
		}
		likes, err := s.LikesByPost(p.ID)
		if err != nil {
			return nil, err
		}
		comments, err := s.CommentsDTO(p.ID)
		if err != nil {
			return nil, err
		}

		cp := p
		posts[i] = forum.PostDTO{
			Post:     &cp,
			Cat:      cat,
			User:     user,
			Likes:    likes,
			Comments: comments}
	}

	return posts, nil
}

// UserPostsDTO func
func (s *Store) UserPostsDTO(userID uuid.UUID) ([]forum.PostDTO, error) {
	pp, err := s.PostsByUser(userID)
	if err != nil {
		return nil, err
	}
	posts := make([]forum.PostDTO, len(pp))

	for i, p := range pp {
		cat, err := s.Cat(p.CatID)
		if err != nil {
			return nil, err
		}
		user, err := s.User(p.UserID)
		if err != nil {
			return nil, err
		}
		likes, err := s.LikesByPost(p.ID)
		if err != nil {
			return nil, err
		}
		comments, err := s.CommentsDTO(p.ID)
		if err != nil {
			return nil, err
		}

		cp := p
		posts[i] = forum.PostDTO{
			Post:     &cp,
			Cat:      cat,
			User:     user,
			Likes:    likes,
			Comments: comments}
	}

	return posts, nil
}

// LikedPostsDTO func
func (s *Store) LikedPostsDTO(userID uuid.UUID) ([]forum.PostDTO, error) {
	pp, err := s.Posts()
	if err != nil {
		return nil, err
	}
	posts := []forum.PostDTO{}

	for _, p := range pp {
		likes, err := s.LikesByPost(p.ID)
		if err != nil {
			return nil, err
		}

		for _, l := range likes {
			if l.UserID == userID {
				cat, err := s.Cat(p.CatID)
				if err != nil {
					return nil, err
				}
				user, err := s.User(p.UserID)
				if err != nil {
					return nil, err
				}

				comments, err := s.CommentsDTO(p.ID)
				if err != nil {
					return nil, err
				}

				cp := p
				posts = append(posts, forum.PostDTO{
					Post:     &cp,
					Cat:      cat,
					User:     user,
					Likes:    likes,
					Comments: comments,
				})
				break
			}
		}

	}

	return posts, nil
}

// CatPostsDTO func
func (s *Store) CatPostsDTO(catID uuid.UUID) ([]forum.PostDTO, error) {
	pp, err := s.PostsByCat(catID)
	if err != nil {
		return nil, err
	}
	posts := make([]forum.PostDTO, len(pp))

	for i, p := range pp {
		cat, err := s.Cat(p.CatID)
		if err != nil {
			return nil, err
		}
		user, err := s.User(p.UserID)
		if err != nil {
			return nil, err
		}
		likes, err := s.LikesByPost(p.ID)
		if err != nil {
			return nil, err
		}
		comments, err := s.CommentsDTO(p.ID)
		if err != nil {
			return nil, err
		}

		cp := p
		posts[i] = forum.PostDTO{
			Post:     &cp,
			Cat:      cat,
			User:     user,
			Likes:    likes,
			Comments: comments}
	}

	return posts, nil
}

// PostDTO func
func (s *Store) PostDTO(id uuid.UUID) (*forum.PostDTO, error) {

	p, err := s.Post(id)
	if err != nil {
		return nil, err
	}
	cat, err := s.Cat(p.CatID)
	if err != nil {
		return nil, err
	}
	user, err := s.User(p.UserID)
	if err != nil {
		return nil, err
	}
	likes, err := s.LikesByPost(p.ID)
	if err != nil {
		return nil, err
	}
	comments, err := s.CommentsDTO(p.ID)
	if err != nil {
		return nil, err
	}

	post := &forum.PostDTO{
		Post:     p,
		Cat:      cat,
		User:     user,
		Likes:    likes,
		Comments: comments,
	}

	return post, nil
}

// CommentsDTO func
func (s *Store) CommentsDTO(postID uuid.UUID) ([]forum.CommentDTO, error) {

	cc, err := s.CommentsByPost(postID)
	if err != nil {
		return nil, err
	}
	comments := make([]forum.CommentDTO, len(cc))

	for i, c := range cc {

		user, err := s.User(c.UserID)
		if err != nil {
			return nil, err
		}
		likes, err := s.LikesByComment(c.ID)
		if err != nil {
			return nil, err
		}

		cp := c
		comments[i] = forum.CommentDTO{
			Comment: &cp,
			User:    user,
			Likes:   likes,
		}
	}

	return comments, nil
}
