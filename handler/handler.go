package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/rtabulov/forum-v2"
	"github.com/rtabulov/forum-v2/cookiestore"
	e "github.com/rtabulov/forum-v2/express"
	"github.com/rtabulov/forum-v2/sqlite"
)

type message struct {
	Text string
	Type string
}

type messages []message

type responseData struct {
	Messages   []e.Message
	User       *forum.User
	Posts      []forum.PostDTO
	Post       *forum.PostDTO
	PageUser   *forum.User
	Cats       []forum.Cat
	Cat        *forum.Cat
	LikedPosts []forum.PostDTO
}

// NewHandler func
func NewHandler(store *sqlite.Store, cs cookiestore.CookieStore) *Handler {
	return &Handler{Store: store, CS: cs}
}

// Handler type
type Handler struct {
	Store *sqlite.Store
	CS    cookiestore.CookieStore
}

// Home func
func (h *Handler) Home() e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/home.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		posts, err := h.Store.PostsDTO()
		if err != nil {
			h.ErrorPage(http.StatusInternalServerError, messageInternalError)
			return
		}

		user, _ := req.CustomData["User"].(*forum.User)

		res.Prepare()
		err = t.Execute(res, responseData{
			Posts:    posts,
			User:     user,
			Messages: res.GetMessages(),
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
		user, _ := req.CustomData["User"].(*forum.User)

		res.Prepare()
		err := t.Execute(res, responseData{
			User:     user,
			Messages: res.GetMessages(),
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
		user, _ := req.CustomData["User"].(*forum.User)

		res.Prepare()
		err := t.Execute(res, responseData{
			User:     user,
			Messages: res.GetMessages(),
		})

		if err != nil {
			log.Print(fmt.Errorf(`error executing signup template: %w`, err))
		}
	}
}
