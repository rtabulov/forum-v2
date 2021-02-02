package handler

import (
	"fmt"

	e "github.com/rtabulov/express"
	"github.com/rtabulov/forum-v2"
)

// Prottected func
func (h *Handler) Prottected() e.Middleware {
	return func(req *e.Request, res *e.Response, next e.Next) {
		if user, _ := req.CustomData["User"]; user == nil {
			res.Redirect(req.Referer())
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
			fmt.Println("Already logged in, User: ", user.Username)
			res.Redirect(req.Referer())
			return
		}

		next()
	}
}
