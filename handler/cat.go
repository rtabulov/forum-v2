package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/rtabulov/forum-v2"
	e "github.com/rtabulov/forum-v2/express"
	uuid "github.com/satori/go.uuid"
)

// CatPage func
func (h *Handler) CatPage() e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/home.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		catparam, ok := req.Param("id")
		catID, err := uuid.FromString(catparam)
		cat, err2 := h.Store.Cat(catID)
		if !ok || err != nil || err2 != nil {
			next()
			return
		}
		posts, err := h.Store.CatPostsDTO(cat.ID)
		if err != nil {
			h.ErrorPage(http.StatusInternalServerError, messageInternalError)(req, res, next)
			return
		}

		user, _ := req.CustomData["User"].(*forum.User)
		res.Prepare()
		err = t.Execute(res, responseData{
			Posts: posts,
			User:  user,
			Cat:   cat,
		})
		if err != nil {
			log.Print(fmt.Errorf(`error executing cat page template: %w`, err))
		}
	}
}

// CatsPage func
func (h *Handler) CatsPage() e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/cats.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		cats, _ := h.Store.Cats()

		user, _ := req.CustomData["User"].(*forum.User)
		res.Prepare()
		err := t.Execute(res, responseData{
			User: user,
			Cats: cats,
		})
		if err != nil {
			log.Print(fmt.Errorf(`error executing cats page template: %w`, err))
		}
	}
}
