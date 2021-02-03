package main

import (
	"log"
	"os"

	"github.com/rtabulov/forum-v2/cookiestore"
	e "github.com/rtabulov/forum-v2/express"
	"github.com/rtabulov/forum-v2/handler"
	"github.com/rtabulov/forum-v2/middleware"
	"github.com/rtabulov/forum-v2/sqlite"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	store, err := sqlite.NewStore("forum.db")
	if err != nil {
		log.Fatal(err)
	}

	// if err := migrate.Migrate("forum.db.sql", store.CatStore.DB); err != nil {
	// 	log.Fatal(err)
	// }

	// init
	cs := cookiestore.New()
	h := handler.NewHandler(store, cs)
	app := e.NewApp()

	// middleware
	app.Use(middleware.Auth(cs))

	app.Page404(h.Page404())

	// routes
	app.Get("/", h.Home())

	app.Get("/login", h.NotLoggedIn(), h.LoginPage())
	app.Get("/logout", h.Prottected(), h.Logout())

	app.Post("/login", h.NotLoggedIn(), h.Login())

	app.Get("/signup", h.NotLoggedIn(), h.SignupPage())
	app.Post("/signup", h.NotLoggedIn(), h.Signup())

	app.Get("/post", h.Prottected(), h.CreatePostPage())
	app.Post("/post", h.Prottected(), h.CreatePost())
	app.Get("/post/:id", h.PostPage())

	app.Post("/post/:id/comment", h.Prottected(), h.CreateComment())

	app.Post("/post/:id/like", h.Prottected(), h.LikePost())

	app.Post("/comment/:id/like", h.Prottected(), h.LikeComment())

	app.Get("/user/:username", h.UserPage())

	app.Get("/categories/:id", h.CatPage())

	app.Get("/categories", h.CatsPage())

	app.Listen(port)
}
