package express

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rtabulov/forum-v2/express/helpers"
)

// App q
type App struct {
	tree          routeTree
	useMiddleware []Middleware
	page404       Middleware
}

// Map type
type Map map[string]interface{}

// map[method][]Middleware
type allowedMethods map[string][]Middleware

// map[path]allowedMethods
type routeTree map[string]allowedMethods

// NewApp func
func NewApp() *App {
	return &App{
		useMiddleware: make([]Middleware, 0),
		tree:          make(routeTree),
		page404: func(req *Request, res *Response, next Next) {
			res.Error("Not found", http.StatusNotFound)
		},
	}
}

// Use func
func (app *App) Use(mws ...Middleware) {
	app.useMiddleware = append(app.useMiddleware, mws...)
}

// Page404 func
func (app *App) Page404(mw Middleware) {
	app.page404 = mw
}

func (app *App) createEndpoint(path, method string, mw []Middleware) {
	methods, ok := app.tree[path]
	if !ok {
		app.tree[path] = make(allowedMethods)
	}

	methods = app.tree[path]

	if mws, ok := methods[method]; ok {
		methods[method] = append(mws, mw...)
	} else {
		methods[method] = mw
	}
}

// Get func
func (app *App) Get(path string, mw ...Middleware) {
	app.createEndpoint(path, http.MethodGet, mw)
}

// Post func
func (app *App) Post(path string, mw ...Middleware) {
	app.createEndpoint(path, http.MethodPost, mw)
}

// Put func
func (app *App) Put(path string, mw ...Middleware) {
	app.createEndpoint(path, http.MethodPut, mw)
}

// Patch func
func (app *App) Patch(path string, mw ...Middleware) {
	app.createEndpoint(path, http.MethodPatch, mw)
}

// Delete func
func (app *App) Delete(path string, mw ...Middleware) {
	app.createEndpoint(path, http.MethodDelete, mw)
}

func (app *App) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	request, response := newRequestResponse(app, req, w)
	queue := app.useMiddleware

	found := false
	for pattern, methods := range app.tree {
		if match, params := helpers.MatchAndParams(pattern, request.URL.Path); match {
			found = true
			request.params = helpers.MergeStringMaps(request.params, params)

			if mws, ok := methods[request.Method]; ok {
				queue = append(queue, mws...)
			} else {
				response.Status(http.StatusMethodNotAllowed)
				queue = append(queue, app.page404)
			}
		}
	}

	if !found {
		response.Status(http.StatusNotFound)
		queue = append(queue, app.page404)
	}

	// execute middleware
	for _, mw := range queue {
		cntnue := false
		next := func() {
			cntnue = true
		}

		mw(request, response, next)

		if !cntnue {
			break
		}
	}
}

// Listen func
func (app *App) Listen(port string) {
	fmt.Printf("Express app starting on port %s\n", port)
	err := http.ListenAndServe(":"+port, app)
	log.Fatal(err)
}

func newRequestResponse(app *App, r *http.Request, w http.ResponseWriter) (*Request, *Response) {
	return newRequest(r, w, app), newResponse(r, w, app)
}
