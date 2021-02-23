package handler

import (
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
			// h.ErrorPage(http.StatusBadRequest, "i see what you're trying to do here... don't do that. just don't")(req, res, next)
			res.AddMessage("danger", "i see what you're trying to do here... don't do that. just don't")
			h.PostPage()(req, res, next)
			return
		}

		res.Redirect("/post/" + post)
	}
}
