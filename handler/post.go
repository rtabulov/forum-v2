package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/rtabulov/forum-v2"
	e "github.com/rtabulov/forum-v2/express"
	uuid "github.com/satori/go.uuid"
)

// PostPage func
func (h *Handler) PostPage() e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/post.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		id, ok := req.Param("id")
		uid, err := uuid.FromString(id)

		if !ok || err != nil {
			next()
			return
		}

		post, err := h.Store.PostDTO(uid)
		if err != nil {
			next()
			return
		}

		user, _ := req.CustomData["User"].(*forum.User)

		res.Prepare()
		err = t.Execute(res, responseData{
			Post: post,
			User: user,
		})

		if err != nil {
			log.Print(fmt.Errorf(`error executing post template: %w`, err))
		}
	}
}

// CreatePostPage func
func (h *Handler) CreatePostPage() e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/create-post.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		cats, err := h.Store.Cats()
		if err != nil {
			h.ErrorPage(http.StatusInternalServerError, messageInternalError)(req, res, next)
			return
		}

		user, _ := req.CustomData["User"].(*forum.User)

		res.Prepare()
		err = t.Execute(res, responseData{
			User:     user,
			Cats:     cats,
			Messages: res.GetMessages(),
		})

		if err != nil {
			log.Print(fmt.Errorf(`error executing create-post template: %w`, err))
		}
	}
}

// CreatePost func
func (h *Handler) CreatePost() e.Middleware {
	return func(req *e.Request, res *e.Response, next e.Next) {
		user, _ := req.CustomData["User"].(*forum.User)
		if user == nil {
			h.ErrorPage(http.StatusUnauthorized, messageUnauthorized)(req, res, next)
			return
		}
		title := strings.TrimSpace(req.FormValue("title"))
		body := strings.TrimSpace(req.FormValue("body"))

		p := &forum.Post{
			Title:  title,
			Body:   body,
			UserID: user.ID,
		}

		if err := h.Store.CreatePost(p); err != nil {
			// res.Error("Bad request", http.StatusBadRequest)
			res.Status(http.StatusBadRequest)
			res.AddMessage("danger", "Please enter valid data")
			h.CreatePostPage()(req, res, next)
			log.Println(err)
			return
		}

		cats := req.Form["cat-id"]
		cids := []uuid.UUID{}
		for _, c := range cats {
			cid, err := uuid.FromString(c)
			if err != nil {
				continue
			}
			cids = append(cids, cid)
		}

		if len(cids) == 0 {
			res.Status(http.StatusBadRequest)
			res.AddMessage("danger", "Select at least one category")
			h.CreatePostPage()(req, res, next)
			h.Store.DeletePost(p.ID)
			return
		}

		h.Store.CratePostCats(p.ID, cids)
		res.Redirect("/")
	}
}
