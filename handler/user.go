package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/rtabulov/forum-v2"
	e "github.com/rtabulov/forum-v2/express"
)

// UserPage func
func (h *Handler) UserPage() e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/user.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		username, ok := req.Param("username")
		pageUser, err := h.Store.UserByUsername(username)
		if !ok || err != nil {
			next()
			return
		}

		posts, err := h.Store.UserPostsDTO(pageUser.ID)
		if err != nil {
			h.ErrorPage(http.StatusInternalServerError, messageInternalError)(req, res, next)
			return
		}
		likedPosts, err := h.Store.LikedPostsDTO(pageUser.ID)
		if err != nil {
			h.ErrorPage(http.StatusInternalServerError, messageInternalError)(req, res, next)
			return
		}
		user, _ := req.CustomData["User"].(*forum.User)
		err = t.Execute(res, responseData{
			User:       user,
			PageUser:   pageUser,
			Posts:      posts,
			LikedPosts: likedPosts,
		})

		if err != nil {
			log.Print(fmt.Errorf(`error executing login template: %w`, err))
		}
	}
}
