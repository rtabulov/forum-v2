package forum

import (

	// load environment variables
	"fmt"
	"time"

	// load environment varaiables
	_ "github.com/rtabulov/envar/autoload"
	uuid "github.com/satori/go.uuid"
)

// Cat type
type Cat struct {
	ID          uuid.UUID `db:"cat_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
}

// Post type
type Post struct {
	ID        uuid.UUID `db:"post_id"`
	CatID     uuid.UUID `db:"cat_id"`
	UserID    uuid.UUID `db:"user_id"`
	Title     string    `db:"title"`
	Body      string    `db:"body"`
	CreatedAt int64     `db:"created_at"`
}

// Comment type
type Comment struct {
	ID        uuid.UUID `db:"comment_id"`
	UserID    uuid.UUID `db:"user_id"`
	PostID    uuid.UUID `db:"post_id"`
	Body      string    `db:"body"`
	CreatedAt int64     `db:"created_at"`
}

// User type
type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	Password  string
	Avatar    string
	CreatedAt int64
}

// PostLike type
type PostLike struct {
	PostID uuid.UUID
	UserID uuid.UUID
	Up     bool
}

// CommentLike type
type CommentLike struct {
	CommentID uuid.UUID
	UserID    uuid.UUID
	Up        bool
}

// CommentLikes type
type CommentLikes []CommentLike

// PostDTO type
type PostDTO struct {
	*Post
	Cat      *Cat
	User     *User
	Comments []CommentDTO
	Likes    PostLikes
}

// CommentDTO type
type CommentDTO struct {
	*Comment
	User  *User
	Likes CommentLikes
}

// PostLikes type
type PostLikes []PostLike

// FormattedTime func
func (p PostDTO) FormattedTime() string {
	return formatTime(p.CreatedAt)
}

// FormattedTimePassed func
func (p PostDTO) FormattedTimePassed() string {
	return formatPassedTime(p.CreatedAt)
}

// FormattedTime func
func (p Comment) FormattedTime() string {
	return formatTime(p.CreatedAt)
}

// FormattedTimePassed func
func (p Comment) FormattedTimePassed() string {
	return formatPassedTime(p.CreatedAt)
}

func formatTime(t int64) string {
	return time.Unix(t, 0).Local().Format("2 Jan, 2006")

}

func formatPassedTime(t int64) string {
	duration := time.Unix(time.Now().Unix()-t, 0).Unix()
	if d := duration / 60 / 60 / 24; d > 0 {
		return fmt.Sprintf("%dd", d)
	}
	if h := duration / 60 / 60; h > 0 {
		return fmt.Sprintf("%dh", h)
	}
	if m := duration / 60; m > 0 {
		return fmt.Sprintf("%dm", m)
	}

	return fmt.Sprintf("%ds", duration)
}

// Votes func
func (ll PostLikes) Votes() int {
	n := 0
	for _, l := range ll {
		if l.Up {
			n++
		} else {
			n--
		}
	}
	return n
}

// Upvotes func
func (ll PostLikes) Upvotes() int {
	n := 0
	for _, l := range ll {
		if l.Up {
			n++
		}
	}
	return n
}

// Downvotes func
func (ll PostLikes) Downvotes() int {
	n := 0
	for _, l := range ll {
		if !l.Up {
			n++
		}
	}
	return n
}

// Votes func
func (ll CommentLikes) Votes() int {
	n := 0
	for _, l := range ll {
		if l.Up {
			n++
		} else {
			n--
		}
	}
	return n
}

// Upvotes func
func (ll CommentLikes) Upvotes() int {
	n := 0
	for _, l := range ll {
		if l.Up {
			n++
		}
	}
	return n
}

// Downvotes func
func (ll CommentLikes) Downvotes() int {
	n := 0
	for _, l := range ll {
		if !l.Up {
			n++
		}
	}
	return n
}
