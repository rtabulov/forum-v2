package handler

import (
	"log"
	"net/http"
	"strings"

	"github.com/rtabulov/forum-v2"
	e "github.com/rtabulov/forum-v2/express"
	uuid "github.com/satori/go.uuid"
)

// CreateComment func
func (h *Handler) CreateComment() e.Middleware {
	return func(req *e.Request, res *e.Response, next e.Next) {
		body := strings.TrimSpace(req.FormValue("comment"))
		post, ok := req.Param("id")
		postID, err := uuid.FromString(post)
		if !ok || err != nil {
			h.PostPage()(req, res, next)
			return
		}

		user, _ := req.CustomData["User"].(*forum.User)
		if user == nil {
			h.ErrorPage(http.StatusUnauthorized, messageUnauthorized)(req, res, next)
			return
		}

		c := &forum.Comment{
			Body:   body,
			PostID: postID,
			UserID: user.ID,
		}

		if err := h.Store.CreateComment(c); err != nil {
			h.ErrorPage(http.StatusInternalServerError, messageInternalError)(req, res, next)
			log.Println(err)
			return
		}

		res.Redirect("/post/" + post)
	}
}
