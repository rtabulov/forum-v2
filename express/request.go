package express

import (
	"net/http"
)

// Request q
type Request struct {
	*http.Request
	params     map[string]string
	res        http.ResponseWriter
	App        *App
	CustomData map[string]interface{}
}

func newRequest(r *http.Request, w http.ResponseWriter, app *App) *Request {
	return &Request{
		Request:    r,
		App:        app,
		res:        w,
		CustomData: make(map[string]interface{}),
	}
}

// Params func
func (r *Request) Params() map[string]string {
	return r.params
}

// Param func
func (r *Request) Param(p string) (string, bool) {
	v, ok := r.params[p]
	return v, ok
}

// request.params
// request.body
// requst.cookies
// request.method
// request.url
// request.path
