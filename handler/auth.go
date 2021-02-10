package handler

import (
	"net/http"
	"regexp"
	"strings"

	"github.com/rtabulov/forum-v2"
	"github.com/rtabulov/forum-v2/cookiestore"
	e "github.com/rtabulov/forum-v2/express"
	"golang.org/x/crypto/bcrypt"
)

// Signup func
func (h *Handler) Signup() e.Middleware {
	return func(req *e.Request, res *e.Response, next e.Next) {
		username := strings.TrimSpace(req.FormValue("username"))
		password := req.FormValue("password")
		email := req.FormValue("email")

		// username check
		match, _ := regexp.MatchString(`^[a-zA-Z\d]{3,}$`, username)
		if !match {
			res.Status(http.StatusBadRequest).AddMessage("danger", "Username invalid")
			h.SignupPage()(req, res, next)
			return
		}

		// password check
		if len(password) < 3 {
			res.Status(http.StatusBadRequest).AddMessage("danger", "Password invalid")
			h.SignupPage()(req, res, next)
			return
		}

		// email check
		match, _ = regexp.MatchString(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`, email)
		if !match {
			res.Status(http.StatusBadRequest).AddMessage("danger", "Email invalid")
			h.SignupPage()(req, res, next)
			return
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
		if err != nil {
			h.ErrorPage(http.StatusInternalServerError, messageInternalError)(req, res, next)
			return
		}

		user := &forum.User{Username: username, Email: email, Password: string(hash)}

		if err := h.Store.CreateUser(user); err != nil {
			res.Status(http.StatusConflict).AddMessage("danger", "User already exists")
			h.SignupPage()(req, res, next)
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
			res.Status(http.StatusUnauthorized).AddMessage("danger", "User does not exist")
			h.LoginPage()(req, res, next)
			return
		}

		// passwords do not match
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			res.Status(http.StatusUnauthorized).AddMessage("danger", "Password is incorrect")
			h.LoginPage()(req, res, next)
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
