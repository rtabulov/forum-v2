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

// PostPage func
func (h *Handler) PostPage() e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/post.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		id, ok := req.Param("id")
		uid, err := uuid.FromString(id)

		if !ok || err != nil {
			res.Redirect("/")
			return
		}

		post, err := h.Store.PostDTO(uid)
		if err != nil {
			res.Status(http.StatusInternalServerError).JSON(e.Map{
				"error": err.Error(),
			})
			return
		}

		user, _ := req.CustomData["User"].(*forum.User)

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
			res.Error("internal error", http.StatusInternalServerError)
			return
		}
		user, _ := req.CustomData["User"].(*forum.User)
		err = t.Execute(res, responseData{
			User: user,
			Cats: cats,
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
			res.Error("Unauthorized", http.StatusUnauthorized)
			return
		}
		title := req.FormValue("title")
		body := req.FormValue("body")

		p := &forum.Post{
			Title:  title,
			Body:   body,
			UserID: user.ID,
		}

		if err := h.Store.CreatePost(p); err != nil {
			res.Error("Bad request", http.StatusBadRequest)
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
			h.Store.DeletePost(p.ID)
		} else {
			h.Store.CratePostCats(p.ID, cids)
		}

		res.Redirect("/")
	}
}
