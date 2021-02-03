package express

import (
	"encoding/json"
	"net/http"
	"time"
)

// Response q
type Response struct {
	http.ResponseWriter
	req       *http.Request
	statusSet bool
	App       *App
	status    int
}

func newResponse(r *http.Request, w http.ResponseWriter, app *App) *Response {
	return &Response{
		req:            r,
		ResponseWriter: w,
		statusSet:      false,
		App:            app,
	}
}

type responseFunc func(interface{}) error

func (res *Response) respond(fn responseFunc, data interface{}) error {
	if res.statusSet {
		res.WriteHeader(res.status)
	} else {
		res.WriteHeader(http.StatusOK)
	}

	return fn(data)
}

// JSON write json repsonse
func (res *Response) JSON(data interface{}) error {
	res.Header().Add("Content-Type", "application/json")
	fn := json.NewEncoder(res).Encode

	return res.respond(fn, data)
}

// Send write string rensponse
func (res *Response) Send(data string) error {
	res.Header().Add("Content-Type", "text/plain")

	fn := func(v interface{}) error {
		bytes := v.([]byte)
		_, err := res.Write(bytes)
		return err
	}

	return res.respond(fn, data)
}

// Status func
func (res *Response) Status(code int) *Response {
	res.status = code
	res.statusSet = true
	return res
}

// Error func
func (res *Response) Error(err string, code int) {
	w := res

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	res.Status(code).JSON(err)
}

// SetCookie func
func (res *Response) SetCookie(c *http.Cookie) {
	http.SetCookie(res.ResponseWriter, c)
}

// ClearCookie func
func (res *Response) ClearCookie(name string) {
	c := &http.Cookie{
		Name:    name,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),

		HttpOnly: true,
	}

	http.SetCookie(res.ResponseWriter, c)

}

// Redirect func
func (res *Response) Redirect(url string) {
	status := res.status
	if !res.statusSet {
		status = http.StatusFound
	}
	http.Redirect(res.ResponseWriter, res.req, url, status)
}
