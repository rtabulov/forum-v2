package handler

import (
	"github.com/rtabulov/forum-v2"
	e "github.com/rtabulov/forum-v2/express"
)

// Prottected func
func (h *Handler) Prottected() e.Middleware {
	return func(req *e.Request, res *e.Response, next e.Next) {
		intfc, ok := req.CustomData["User"]
		user, ok2 := intfc.(*forum.User)
		if !ok || intfc == nil || !ok2 || user == nil {
			res.Redirect("/")
			return
		}

		next()
	}
}

// NotLoggedIn func
func (h *Handler) NotLoggedIn() e.Middleware {
	return func(req *e.Request, res *e.Response, next e.Next) {
		intfc := req.CustomData["User"]
		if user, _ := intfc.(*forum.User); user != nil {
			res.Redirect("/")
			return
		}

		next()
	}
}
