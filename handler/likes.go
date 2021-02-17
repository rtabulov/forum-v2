package handler

import (
	"net/http"

	"github.com/rtabulov/forum-v2"
	e "github.com/rtabulov/forum-v2/express"
	uuid "github.com/satori/go.uuid"
)

// LikePost func
func (h *Handler) LikePost() e.Middleware {
	return func(req *e.Request, res *e.Response, next e.Next) {
		user, _ := req.CustomData["User"].(*forum.User)
		if user == nil {
			h.ErrorPage(http.StatusUnauthorized, messageUnauthorized)(req, res, next)
			return
		}

		post, ok := req.Param("id")
		postID, err := uuid.FromString(post)
		if !ok || err != nil {
			next()
			return
		}

		ups := req.FormValue("up")
		up := false
		if ups == "true" {
			up = true
		} else if ups == "false" {
			up = false
		} else {
			res.Status(http.StatusBadRequest)
			h.PostPage()
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
			h.ErrorPage(http.StatusUnauthorized, messageUnauthorized)(req, res, next)
			return
		}

		c, ok := req.Param("id")
		cID, err := uuid.FromString(c)
		if !ok || err != nil {
			res.Status(http.StatusBadRequest)
			h.Home()
			return
		}

		comm, err := h.Store.Comment(cID)
		if err != nil {
			next()
			return
		}

		ups := req.FormValue("up")
		up := false
		if ups == "true" {
			up = true
		} else if ups == "false" {
			up = false
		} else {
			res.Status(http.StatusBadRequest)
			req.SetParam("id", comm.PostID.String())
			h.PostPage()
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

		res.Redirect("/post/" + comm.PostID.String())
	}
}
