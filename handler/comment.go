package handler

import (
	"log"
	"net/http"

	e "github.com/rtabulov/express"
	"github.com/rtabulov/forum-v2"
	uuid "github.com/satori/go.uuid"
)

// CreateComment func
func (h *Handler) CreateComment() e.Middleware {
	return func(req *e.Request, res *e.Response, next e.Next) {
		user, _ := req.CustomData["User"].(*forum.User)
		if user == nil {
			res.Error("Unauthorized", http.StatusUnauthorized)
			return
		}

		body := req.FormValue("comment")
		post, ok := req.Param("id")
		postID, err := uuid.FromString(post)
		if !ok || err != nil {
			res.Error("Bad request", http.StatusBadRequest)
			return
		}

		c := &forum.Comment{
			Body:   body,
			PostID: postID,
			UserID: user.ID,
		}

		if err := h.Store.CreateComment(c); err != nil {
			res.Error("Bad request", http.StatusBadRequest)
			log.Println(err)
			return
		}

		res.Redirect("/post/" + post)
	}
}
