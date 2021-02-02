package handler

import (
	"net/http"

	e "github.com/rtabulov/express"
	"github.com/rtabulov/forum-v2"
	uuid "github.com/satori/go.uuid"
)

// LikePost func
func (h *Handler) LikePost() e.Middleware {
	return func(req *e.Request, res *e.Response, next e.Next) {
		user, _ := req.CustomData["User"].(*forum.User)
		if user == nil {
			res.Error("Unauthorized", http.StatusUnauthorized)
			return
		}

		post, ok := req.Param("id")
		postID, err := uuid.FromString(post)
		if !ok || err != nil {
			res.Error("Bad request", http.StatusBadRequest)
			return
		}

		ups := req.FormValue("up")
		up := false
		if ups == "true" {
			up = true
		} else if ups == "false" {
			up = false
		} else {
			res.Error("Bad request", http.StatusBadRequest)
			return
		}

		l := &forum.PostLike{
			PostID: postID,
			UserID: user.ID,
			Up:     up,
		}

		like, err := h.Store.GetPostLike(l)
		if err != nil {
			h.Store.CreatePostLike(l)
		} else {
			h.Store.DeletePostLike(l)
			if like.Up != l.Up {
				h.Store.CreatePostLike(l)
			}
		}

		res.Redirect("/post/" + post)
	}
}

// LikeComment func
func (h *Handler) LikeComment() e.Middleware {
	return func(req *e.Request, res *e.Response, next e.Next) {
		user, _ := req.CustomData["User"].(*forum.User)
		if user == nil {
			res.Error("Unauthorized", http.StatusUnauthorized)
			return
		}

		c, ok := req.Param("id")
		cID, err := uuid.FromString(c)
		if !ok || err != nil {
			res.Error("Bad request", http.StatusBadRequest)
			return
		}

		ups := req.FormValue("up")
		up := false
		if ups == "true" {
			up = true
		} else if ups == "false" {
			up = false
		} else {
			res.Error("Bad request", http.StatusBadRequest)
			return
		}

		l := &forum.CommentLike{
			CommentID: cID,
			UserID:    user.ID,
			Up:        up,
		}

		like, err := h.Store.GetCommentLike(l)
		if err != nil {
			err = h.Store.CreateCommentLike(l)
		} else {
			err = h.Store.DeleteCommentLike(l)
			if like.Up != l.Up {
				err = h.Store.CreateCommentLike(l)
			}
		}

		comm, _ := h.Store.Comment(cID)
		res.Redirect("/post/" + comm.PostID.String())
	}
}
