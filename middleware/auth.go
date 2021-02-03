package middleware

import (
	"github.com/rtabulov/forum-v2/cookiestore"
	e "github.com/rtabulov/forum-v2/express"
	uuid "github.com/satori/go.uuid"
)

// Auth func
func Auth(cs cookiestore.CookieStore) e.Middleware {

	return func(req *e.Request, res *e.Response, next e.Next) {
		defer next()

		c, err := req.Cookie(cookiestore.CookieName)

		// cookie not set
		if err != nil {
			res.SetCookie(cs.SetNewGuestCookie())
			return
		}

		// cookie invalid
		id, err := uuid.FromString(c.Value)
		data, ok := cs.Get(id)
		if err != nil || !ok {
			res.SetCookie(cs.SetNewGuestCookie())
			return
		}

		req.CustomData["User"] = data.User
	}
}
