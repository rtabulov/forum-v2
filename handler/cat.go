package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	e "github.com/rtabulov/express"
	"github.com/rtabulov/forum-v2"
	uuid "github.com/satori/go.uuid"
)

// CatPage func
func (h *Handler) CatPage() e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/home.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		catparam, ok := req.Param("id")
		catID, err := uuid.FromString(catparam)
		if !ok || err != nil {
			res.Redirect("/")
			return
		}
		posts, err := h.Store.CatPostsDTO(catID)
		if err != nil {
			res.Status(http.StatusInternalServerError).JSON(e.Map{
				"error": err.Error(),
			})
			return
		}
		cat, _ := h.Store.Cat(catID)

		user, _ := req.CustomData["User"].(*forum.User)
		err = t.Execute(res, responseData{
			Posts: posts,
			User:  user,
			Cat:   cat,
		})
		if err != nil {
			log.Print(fmt.Errorf(`error executing home template: %w`, err))
		}
	}
}

// CatsPage func
func (h *Handler) CatsPage() e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/cats.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		cats, _ := h.Store.Cats()

		user, _ := req.CustomData["User"].(*forum.User)
		err := t.Execute(res, responseData{
			User: user,
			Cats: cats,
		})
		if err != nil {
			log.Print(fmt.Errorf(`error executing home template: %w`, err))
		}
	}
}
