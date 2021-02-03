package handler

import (
	"net/http"
	"net/url"
	"regexp"

	"github.com/rtabulov/forum-v2"
	"github.com/rtabulov/forum-v2/cookiestore"
	e "github.com/rtabulov/forum-v2/express"
	"golang.org/x/crypto/bcrypt"
)

// Signup func
func (h *Handler) Signup() e.Middleware {
	return func(req *e.Request, res *e.Response, next e.Next) {
		username := req.FormValue("username")
		password := req.FormValue("password")
		email := req.FormValue("email")

		// username check
		match, _ := regexp.MatchString(`^[a-zA-Z\d]{3,}$`, username)
		if !match {
			res.Redirect("/signup?error=Username%20invalid")
			return
		}

		// password check
		if len(password) < 3 {
			res.Redirect("/signup?error=Password%20invalid")
			return
		}

		// email check
		match, _ = regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, email)
		if !match {
			res.Redirect("/signup?error=Email%20invalid")
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
		if err != nil {
			res.Error("internal error", http.StatusInternalServerError)
			return
		}

		user := &forum.User{Username: username, Email: email, Password: string(hash)}

		if err := h.Store.CreateUser(user); err != nil {
			res.Redirect("/signup?error=" + url.QueryEscape(err.Error()))
			return
		}

		c := h.CS.SetNewCookie(user)
		res.SetCookie(c)

		res.Redirect("/")
	}
}

// Login func
func (h *Handler) Login() e.Middleware {
	return func(req *e.Request, res *e.Response, next e.Next) {
		username := req.FormValue("username")
		password := req.FormValue("password")

		user, err := h.Store.UserByUsername(username)

		// user does not exist
		if err != nil {
			res.Redirect("/login?error=User%20does%20not%20exist")
			return
		}

		// passwords do not match
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {

			res.Redirect("/login?error=Password%20incorrect")
			return
		}

		// success
		c := h.CS.SetNewCookie(user)
		res.SetCookie(c)

		res.Redirect("/")
	}
}

// Logout func
func (h *Handler) Logout() e.Middleware {
	return func(req *e.Request, res *e.Response, next e.Next) {
		res.ClearCookie(cookiestore.CookieName)
		res.Redirect("/")
	}
}
