package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	e "github.com/rtabulov/express"
	"github.com/rtabulov/forum-v2"
	"github.com/rtabulov/forum-v2/cookiestore"
	"github.com/rtabulov/forum-v2/sqlite"
)

type message struct {
	Text string
	Type string
}

type messages []message

type responseData struct {
	Messages messages
	User     *forum.User
	Posts    []forum.PostDTO
	Post     *forum.PostDTO
	PageUser *forum.User
	Cats     []forum.Cat
}

// NewHandler func
func NewHandler(store *sqlite.Store, cs cookiestore.CookieStore) *Handler {
	t := template.Must(template.ParseFiles("views/header.html"))
	return &Handler{Store: store, t: t, CS: cs}
}

// Handler type
type Handler struct {
	Store *sqlite.Store
	CS    cookiestore.CookieStore
	t     *template.Template
}

// Home func
func (h *Handler) Home() e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/home.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		posts, err := h.Store.PostsDTO()
		if err != nil {
			res.Status(http.StatusInternalServerError).JSON(e.Map{
				"error": err.Error(),
			})
			return
		}

		user, _ := req.CustomData["User"].(*forum.User)
		err = t.Execute(res, responseData{
			Posts: posts,
			User:  user,
		})
		if err != nil {
			log.Print(fmt.Errorf(`error executing home template: %w`, err))
		}
	}
}

// LoginPage func
func (h *Handler) LoginPage() e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/login.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		msgs := messages{}
		if err := req.FormValue("error"); err != "" {
			msgs = append(msgs, message{err, "danger"})
		}

		user, _ := req.CustomData["User"].(*forum.User)
		err := t.Execute(res, responseData{
			User:     user,
			Messages: msgs,
		})

		if err != nil {
			log.Print(fmt.Errorf(`error executing login template: %w`, err))
		}
	}
}

// SignupPage func
func (h *Handler) SignupPage() e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/signup.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		msgs := messages{}
		if err := req.FormValue("error"); err != "" {
			msgs = append(msgs, message{err, "danger"})
		}

		user, _ := req.CustomData["User"].(*forum.User)
		err := t.Execute(res, responseData{
			User:     user,
			Messages: msgs,
		})

		if err != nil {
			log.Print(fmt.Errorf(`error executing signup template: %w`, err))
		}
	}
}
