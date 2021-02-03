package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	e "github.com/rtabulov/express"
	"github.com/rtabulov/forum-v2"
)

// UserPage func
func (h *Handler) UserPage() e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/user.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		username, ok := req.Param("username")
		pageUser, err := h.Store.UserByUsername(username)
		if !ok || err != nil {
			res.Error("user not found", http.StatusNotFound)
			return
		}

		posts, err := h.Store.UserPostsDTO(pageUser.ID)
		if err != nil {
			res.Error("Internal error", http.StatusInternalServerError)
			return
		}
		likedPosts, err := h.Store.LikedPostsDTO(pageUser.ID)
		if err != nil {
			res.Error("Internal error", http.StatusInternalServerError)
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
