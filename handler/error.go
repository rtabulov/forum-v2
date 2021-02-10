package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/rtabulov/forum-v2"
	e "github.com/rtabulov/forum-v2/express"
)

const messageNotFound = "Page not found 🙈"
const messageInternalError = "Something went wrong on the server 😱"
const messageUnauthorized = "Stop right there 🚨! You are not allowed here 👮‍♂️"

// ErrorPage func
func (h *Handler) ErrorPage(status int, message string) e.Middleware {
	t := template.Must(template.ParseFiles("views/header.html", "views/404.html"))
	return func(req *e.Request, res *e.Response, next e.Next) {
		res.Status(status)
		res.Prepare()
		user, _ := req.CustomData["User"].(*forum.User)
		err := t.Execute(res, e.Map{
			"User":    user,
			"Status":  status,
			"Message": message,
		})
		if err != nil {
			log.Print(fmt.Errorf(`error executing error template: %w`, err))
		}
	}
}

// Page404 func
func (h *Handler) Page404() e.Middleware {
	return h.ErrorPage(http.StatusNotFound, messageNotFound)
}
